package student

import (
	context2 "context"
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// quickCheckStudentIdentityURL 快速获取学生身份
	quickCheckStudentIdentityURL = "https://api.weixin.qq.com/intp/quickcheckstudentidentity?access_token=%s"
)

// Student 学生身份核验
type Student struct {
	*context.Context
}

// NewStudent init
func NewStudent(ctx *context.Context) *Student {
	return &Student{ctx}
}

// BindStatus 学生身份绑定状态
type BindStatus int

const (
	// BindStatusUnbound 未绑定
	BindStatusUnbound BindStatus = 1
	// BindStatusAuditing 审核中
	BindStatusAuditing BindStatus = 2
	// BindStatusBound 已绑定
	BindStatusBound BindStatus = 3
)

// QuickCheckStudentIdentityRequest 快速获取学生身份请求参数
type QuickCheckStudentIdentityRequest struct {
	OpenID             string `json:"openid"`               // 用户在业务方下的openid，需为已开通接口权限小程序对应的用户openid
	WxStudentCheckCode string `json:"wx_studentcheck_code"` // 用户授权查询code，由授权插件返回，code有效期两小时
}

// QuickCheckStudentIdentityResponse 快速获取学生身份返回参数
type QuickCheckStudentIdentityResponse struct {
	util.CommonError
	BindStatus BindStatus `json:"bind_status"` // 绑定状态：1-未绑定，2-审核中，3-已绑定
	IsStudent  bool       `json:"is_student"`  // 用户学生身份绑定状态说明，true-是学生，false-不是学生
}

// QuickCheckStudentIdentity 快速获取学生身份
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/student/api_quickcheckstudentidentity.html
func (student *Student) QuickCheckStudentIdentity(in *QuickCheckStudentIdentityRequest) (out QuickCheckStudentIdentityResponse, err error) {
	return student.QuickCheckStudentIdentityContext(context2.Background(), in)
}

// QuickCheckStudentIdentityContext 快速获取学生身份
func (student *Student) QuickCheckStudentIdentityContext(ctx context2.Context, in *QuickCheckStudentIdentityRequest) (out QuickCheckStudentIdentityResponse, err error) {
	var accessToken string
	if accessToken, err = student.GetAccessTokenContext(ctx); err != nil {
		return
	}

	uri := fmt.Sprintf(quickCheckStudentIdentityURL, accessToken)
	var response []byte
	if response, err = util.PostJSONContext(ctx, uri, in); err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "student.QuickCheckStudentIdentity")
	return
}
