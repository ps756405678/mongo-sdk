package domain

type DeleteWrapper struct {
	CallSdkReq
	Data []string `json:"data"`
}
