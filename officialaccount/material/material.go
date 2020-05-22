package material

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	addNewsURL     = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	addMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	delMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	getMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/get_material"
)

//Material 素材管理
type Material struct {
	*context.Context
}

//NewMaterial init
func NewMaterial(context *context.Context) *Material {
	material := new(Material)
	material.Context = context
	return material
}

//Article 永久图文素材
type Article struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	ThumbURL         string `json:"thumb_url"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     int    `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
	URL              string `json:"url"`
	DownURL          string `json:"down_url"`
}

// GetNews 获取/下载永久素材
func (material *Material) GetNews(id string) ([]*Article, error) {
	accessToken, err := material.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", getMaterialURL, accessToken)

	var req struct {
		MediaID string `json:"media_id"`
	}
	req.MediaID = id
	responseBytes, err := util.PostJSON(uri, req)

	var res struct {
		NewsItem []*Article `json:"news_item"`
	}
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return nil, err
	}

	return res.NewsItem, nil
}

//reqArticles 永久性图文素材请求信息
type reqArticles struct {
	Articles []*Article `json:"articles"`
}

//resArticles 永久性图文素材返回结果
type resArticles struct {
	util.CommonError

	MediaID string `json:"media_id"`
}

//AddNews 新增永久图文素材
func (material *Material) AddNews(articles []*Article) (mediaID string, err error) {
	req := &reqArticles{articles}

	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", addNewsURL, accessToken)
	responseBytes, err := util.PostJSON(uri, req)
	var res resArticles
	err = json.Unmarshal(responseBytes, &res)
	if err != nil {
		return
	}
	mediaID = res.MediaID
	return
}

//resAddMaterial 永久性素材上传返回的结果
type resAddMaterial struct {
	util.CommonError

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

//AddMaterial 上传永久性素材（处理视频需要单独上传）
func (material *Material) AddMaterial(mediaType MediaType, filename string) (mediaID string, url string, err error) {
	if mediaType == MediaTypeVideo {
		err = errors.New("永久视频素材上传使用 AddVideo 方法")
	}
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=%s", addMaterialURL, accessToken, mediaType)
	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

type reqVideo struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

//AddVideo 永久视频素材文件上传
func (material *Material) AddVideo(filename, title, introduction string) (mediaID string, url string, err error) {
	var accessToken string
	accessToken, err = material.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=video", addMaterialURL, accessToken)

	videoDesc := &reqVideo{
		Title:        title,
		Introduction: introduction,
	}
	var fieldValue []byte
	fieldValue, err = json.Marshal(videoDesc)
	if err != nil {
		return
	}

	fields := []util.MultipartFormField{
		{
			IsFile:    true,
			Fieldname: "media",
			Filename:  filename,
		},
		{
			IsFile:    false,
			Fieldname: "description",
			Value:     fieldValue,
		},
	}

	var response []byte
	response, err = util.PostMultipartForm(fields, uri)
	if err != nil {
		return
	}

	var resMaterial resAddMaterial
	err = json.Unmarshal(response, &resMaterial)
	if err != nil {
		return
	}
	if resMaterial.ErrCode != 0 {
		err = fmt.Errorf("AddMaterial error : errcode=%v , errmsg=%v", resMaterial.ErrCode, resMaterial.ErrMsg)
		return
	}
	mediaID = resMaterial.MediaID
	url = resMaterial.URL
	return
}

type reqDeleteMaterial struct {
	MediaID string `json:"media_id"`
}

//DeleteMaterial 删除永久素材
func (material *Material) DeleteMaterial(mediaID string) error {
	accessToken, err := material.GetAccessToken()
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s?access_token=%s", delMaterialURL, accessToken)
	response, err := util.PostJSON(uri, reqDeleteMaterial{mediaID})
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(response, "DeleteMaterial")
}
