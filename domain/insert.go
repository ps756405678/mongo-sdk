package domain

import "encoding/json"

type InsertOneReq[T any] struct {
	Req  CallSdkReq
	Data T `json:"data"`
}

type InsertManyReq[T any] struct {
	Req  CallSdkReq
	Data []T `json:"data"`
}

func (req *InsertOneReq[T]) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

func (req *InsertManyReq[T]) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}
