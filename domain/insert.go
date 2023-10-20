package domain

import "encoding/json"

type InsertOneReq[T any] struct {
	CallSdkReq
	InsertOption string `json:"insert_option"` // abort / replace
	Data         T      `json:"data"`
}

type InsertManyReq[T any] struct {
	CallSdkReq
	InsertOption string `json:"insert_option"` // abort / replace / error
	Data         []T    `json:"data"`
}

func (req *InsertOneReq[T]) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

func (req *InsertManyReq[T]) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}

func (req *InsertManyReq[T]) AbortWhenDuplicateKey() *InsertManyReq[T] {
	req.InsertOption = "abort"
	return req
}

func (req *InsertManyReq[T]) ReplaceWhenDuplicateKey() *InsertManyReq[T] {
	req.InsertOption = "replace"
	return req
}

func (req *InsertManyReq[T]) ErrorWhenDuplicateKey() *InsertManyReq[T] {
	req.InsertOption = "error"
	return req
}
