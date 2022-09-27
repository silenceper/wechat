package material

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// UploadImgURL 上传图片
	UploadImgURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
)

// UploadImgResponse 上传图片响应
type UploadImgResponse struct {
	util.CommonError
	URL string `json:"url"`
}

// UploadImg 上传图片
// @see https://developer.work.weixin.qq.com/document/path/90256
func (r *Client) UploadImg(filename string) (*UploadImgResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostFile("media", filename, fmt.Sprintf(UploadImgURL, accessToken)); err != nil {
		return nil, err
	}
	result := &UploadImgResponse{}
	if err = util.DecodeWithError(response, result, "UploadImg"); err != nil {
		return nil, err
	}
	return result, nil
}
