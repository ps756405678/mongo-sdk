package entry

// sdk接口定义

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/ps756405678/mongo-sdk/domain"
)

const (
	applicationId = "Application-Id"
	modelId       = "Model-Id"
	instanceId    = "Instance-Id"
)

func FindOne[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result T, err error) {
	resp, err := callSdkService[T](httpReq, &query)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func FindMany[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result []T, err error) {
	resp, err := callSdkService[[]T](httpReq, &query)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func InsertOne[T any](httpReq *http.Request, req domain.InsertOneReq[T]) (result string, err error) {
	resp, err := callSdkService[string](httpReq, &req)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func InsertMany[T any](httpReq *http.Request, req domain.InsertManyReq[T]) (result []string, err error) {
	resp, err := callSdkService[[]string](httpReq, &req)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func UpdateById[T any](httpReq *http.Request, req domain.UpdateOneReq[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &req)
	if err != nil {
		return
	}
	result = resp.Result
	return
}

func UpdateMany[T any](httpReq *http.Request, req domain.UpdateWrapper[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &req)
	if err != nil {
		return
	}
	result = resp.Result
	return
}

func DeleteOne[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &query)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func DeleteMany[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &query)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

// 调用SDK sevice
func callSdkService[T any](httpReq *http.Request, req domain.SdkServiceReq) (resp domain.CallSdkResp[T], err error) {
	// 序列化参数
	bData := req.ToJson()

	// 调用sdk service
	// TODO:此处调用的链接为临时解决方案，后续会给出serverless的sdk，使用serverless的sdk来获取此链接
	request, err := http.NewRequest("POST", "http://mongo-sdk-v1.mongo-sdk-klskoz.svc.cluster.local", bytes.NewReader(bData))
	if err != nil {
		return
	}

	// 将model db上下文参数传递给sdk service
	request.Header.Add(applicationId, httpReq.Header.Get(applicationId))
	request.Header.Add(modelId, httpReq.Header.Get(modelId))
	request.Header.Add(instanceId, httpReq.Header.Get(instanceId))

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	var buff = make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	// 反序列化结果
	err = json.Unmarshal(buff, &resp)
	return
}
