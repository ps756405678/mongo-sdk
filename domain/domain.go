package domain

// sdk接口参数的数据结构定义

// 接口调用的数据结构定义
type CallSdkReq struct {
	Method     string `json:"method"`
	Schema     string `json:"schema"`
	Collection string `json:"collection"`
}

// 接口返回的结果数据结构定义
type CallSdkResp[T any] struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
	Result     T      `json:"result"`
}

type SdkServiceReq interface {
	ToJson() []byte
}
