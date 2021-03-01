package code

import "github.com/silenceper/wechat/v2/util"

//UploadCodeParam 上传代码参数
type UploadCodeParams struct {
	TemplateId  int64  `json:"template_id"`
	ExtJson     string `json:"ext_json"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
}

//SubmitAuditParams 上传代码参数
type SubmitAuditParams struct {
	ItemList      *SubmitAuditItemList    `json:"item_list"`      //审核项列表（选填，至多填写 5 项）
	PreviewInfo   *SubmitAuditPreviewInfo `json:"preview_info"`   //预览信息（小程序页面截图和操作录屏）
	VersionDesc   string                  `json:"version_desc"`   //小程序版本说明和功能解释
	FeedbackInfo  string                  `json:"feedback_info"`  //反馈内容，至多 200 字
	FeedbackStuff string                  `json:"feedback_stuff"` //用 | 分割的 media_id 列表，至多 5 张图片, 可以通过新增临时素材接口上传而得到
	UgcDeclare    *SubmitAuditUgcDeclare  `json:"ugc_declare"`    //用户生成内容场景（UGC）信息安全声明
}

//SubmitAuditItemList 审核项列表（选填，至多填写 5 项）
type SubmitAuditItemList struct {
	Address     string `json:"address"`      //	否	小程序的页面，可通过获取小程序的页面列表接口获得
	Tag         string `json:"tag"`          //	否	小程序的标签，用空格分隔，标签至多 10 个，标签长度至多 20
	FirstClass  string `json:"first_class"`  //	否	一级类目名称
	SecondClass string `json:"second_class"` //	否	二级类目名称
	ThirdClass  string `json:"third_class"`  //	否	三级类目名称
	FirstId     string `json:"first_id"`     //	否	一级类目的 ID
	SecondId    string `json:"second_id"`    //	否	二级类目的 ID
	ThirdId     string `json:"third_id"`     //	否	三级类目的 ID
	Title       string `json:"title"`        //	否	小程序页面的标题,标题长度至多 32
}

//SubmitAuditPreviewInfo 预览信息（小程序页面截图和操作录屏）
type SubmitAuditPreviewInfo struct {
	VideoIdList []string `json:"video_id_list"` //否	录屏mediaid列表，可以通过提审素材上传接口获得
	PicIdList   []string `json:"pic_id_list"`   //否	截屏mediaid列表，可以通过提审素材上传接口获得
}

//SubmitAuditUgcDeclare 用户生成内容场景（UGC）信息安全声明
type SubmitAuditUgcDeclare struct {
	Scene          []int64 `json:"scene"`            //否	UGC场景 0,不涉及用户生成内容, 1.用户资料,2.图片,3.视频,4.文本,5其他, 可多选,当scene填0时无需填写下列字段
	OtherSceneDesc string  `json:"other_scene_desc"` //	否	当scene选其他时的说明,不超时256字
	Method         []int64 `json:"method"`           //	否	内容安全机制 1.使用平台建议的内容安全API,2.使用其他的内容审核产品,3.通过人工审核把关,4.未做内容审核把关
	HasAuditTeam   int64   `json:"has_audit_team"`   //	否	是否有审核团队, 0.无,1.有,默认0
	AuditDesc      string  `json:"audit_desc"`       //	否	说明当前对UGC内容的审核机制,不超过256字
}

//GetAuditstatusResponse 获取审核状态返回结果
type GetAuditstatusResponse struct {
	util.CommonError
	Status     int64  `json:"status"`     //0	审核成功  1	审核被拒绝 2	审核中 3	已撤回 4	审核延后
	Reason     string `json:"reason"`     //当 status = 1 时，返回的拒绝原因; status = 4 时，返回的延后原因
	Screenshot string `json:"screenshot"` //当 status = 1 时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

//GetLastAuditstatusResponse 获取最后一次审核状态返回结果
type GetLastAuditstatusResponse struct {
	util.CommonError
	Auditid    int64  `json:"auditid"`
	Status     int64  `json:"status"`     //0	审核成功  1	审核被拒绝 2	审核中 3	已撤回 4	审核延后
	Reason     string `json:"reason"`     //当 status = 1 时，返回的拒绝原因; status = 4 时，返回的延后原因
	Screenshot string `json:"ScreenShot"` //当 status = 1 时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}
