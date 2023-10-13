package domain

import "encoding/json"

// sdk接口参数的数据结构定义

// 接口调用的数据结构定义
type CallSdkReq struct {
	Schema     string `json:"schema"`
	Collection string `json:"collection"`
}

// 接口返回的结果数据结构定义
type CallSdkResp[T any] struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
	Result     T      `json:"result"`
}

type CreateCollectionResp struct {
	CollectionName string `json:"collection_name"`
	ConnectStr     string `json:"connect_str"`
}

type SaveHookReq struct {
	Req  CallSdkReq
	Data []any `json:"data"`
}

type DeleteHookReq struct {
	Req  CallSdkReq
	Data []string `json:"data"`
}

func (req *CallSdkReq) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

func (req *SaveHookReq) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

func (req *DeleteHookReq) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

type SdkServiceReq interface {
	ToJson() []byte
}
