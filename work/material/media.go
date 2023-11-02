package material

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// uploadImgURL 上传图片
	uploadImgURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
	// uploadTempFile 上传临时素材
	uploadTempFile = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	// uploadAttachment 上传附件资源
	uploadAttachment = "https://qyapi.weixin.qq.com/cgi-bin/media/upload_attachment?access_token=%s&media_type=%s&attachment_type=%d"
)

// UploadImgResponse 上传图片响应
type UploadImgResponse struct {
	util.CommonError
	URL string `json:"url"`
}

// UploadTempFileResponse 上传临时素材响应
type UploadTempFileResponse struct {
	util.CommonError
	MediaID  string `json:"media_id"`
	CreateAt string `json:"created_at"`
	Type     string `json:"type"`
}

// UploadAttachmentResponse 上传资源附件响应
type UploadAttachmentResponse struct {
	util.CommonError
	MediaID  string `json:"media_id"`
	CreateAt int64  `json:"created_at"`
	Type     string `json:"type"`
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
	if response, err = util.PostFile("media", filename, fmt.Sprintf(uploadImgURL, accessToken)); err != nil {
		return nil, err
	}
	result := &UploadImgResponse{}
	err = util.DecodeWithError(response, result, "UploadImg")
	return result, err
}

// UploadTempFile 上传临时素材
// @see https://developer.work.weixin.qq.com/document/path/90253
// @mediaType 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
func (r *Client) UploadTempFile(filename string, mediaType string) (*UploadTempFileResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostFile("media", filename, fmt.Sprintf(uploadTempFile, accessToken, mediaType)); err != nil {
		return nil, err
	}
	result := &UploadTempFileResponse{}
	err = util.DecodeWithError(response, result, "UploadTempFile")
	return result, err
}

// UploadAttachment 上传附件资源
// @see https://developer.work.weixin.qq.com/document/path/95098
// @mediaType 媒体文件类型，分别有图片（image）、视频（video）、普通文件（file）
// @attachment_type 附件类型，不同的附件类型用于不同的场景。1：朋友圈；2:商品图册
func (r *Client) UploadAttachment(filename string, mediaType string, attachmentType int) (*UploadAttachmentResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostFile("media", filename, fmt.Sprintf(uploadAttachment, accessToken, mediaType, attachmentType)); err != nil {
		return nil, err
	}
	result := &UploadAttachmentResponse{}
	err = util.DecodeWithError(response, result, "UploadAttachment")
	return result, err
}
