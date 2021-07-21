# finance
企业微信会话存档SDK（基于企业微信C版官方SDK封装），暂时只支持在`linux`环境下使用当前SDK。

### 官方文档地址
https://open.work.weixin.qq.com/api/doc/90000/90135/91774

### Example

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/NICEXAI/WeWorkFinanceSDK"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	corpID := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	corpSecret := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	rsaPrivateKey := `XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX`

	//初始化客户端
	client, err := WeWorkFinanceSDK.NewClient(corpID, corpSecret, rsaPrivateKey)
	if err != nil {
		fmt.Printf("SDK 初始化失败：%v \n", err)
		return
	}

	//同步消息
	chatDataList, err := client.GetChatData(0, 100, "", "", 3)
	if err != nil {
		fmt.Printf("消息同步失败：%v \n", err)
		return
	}

	for _, chatData := range chatDataList {
		//消息解密
		chatInfo, err := client.DecryptData(chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if err != nil {
			fmt.Printf("消息解密失败：%v \n", err)
			return
		}

		if chatInfo.Type == "image" {
			image := chatInfo.GetImageMessage()
			sdkfileid := image.Image.SdkFileId

			isFinish := false
			buffer := bytes.Buffer{}
			for !isFinish {
				//获取媒体数据
				mediaData, err := client.GetMediaData("", sdkfileid, "", "", 5)
				if err != nil {
					fmt.Printf("媒体数据拉取失败：%v \n", err)
					return
				}
				buffer.Write(mediaData.Data)
				if mediaData.IsFinish {
					isFinish = mediaData.IsFinish
				}
			}
			filePath, _ := os.Getwd()
			filePath = path.Join(filePath, "test.png")
			err := ioutil.WriteFile(filePath, buffer.Bytes(), 0666)
			if err != nil {
				fmt.Printf("文件存储失败：%v \n", err)
				return
			}
			break
		}
	}
}



```
