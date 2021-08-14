package kf

// BaseModel 基础数据
type BaseModel struct {
	ErrCode int    `json:"errcode"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg"`  // 返回码提示语
}
