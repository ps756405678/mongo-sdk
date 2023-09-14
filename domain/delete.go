package domain

import "encoding/json"

type DeleteWrapper struct {
	CallSdkReq
	Data []string `json:"data"`
}

func (req *DeleteWrapper) ToJson() []byte {
	bdata, _ := json.Marshal(req)
	return bdata
}
