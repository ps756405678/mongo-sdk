package entry

// sdk接口定义

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ps756405678/mongo-sdk/consts"
	"github.com/ps756405678/mongo-sdk/domain"
	"github.com/ps756405678/mongo-sdk/method"
)

const (
	applicationId = "Application-Id"
	modelId       = "Model-Id"
	instanceId    = "Instance-Id"
)

func CreateCollection(httpReq *http.Request, req domain.CallSdkReq) (result domain.CreateCollectionResp, err error) {
	resp, err := callSdkService[domain.CreateCollectionResp](httpReq, &req, method.CreateCollection)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func FindOne[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result T, err error) {
	resp, err := callSdkService[T](httpReq, &query, method.FindOne)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func FindMany[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result []T, err error) {
	resp, err := callSdkService[[]T](httpReq, &query, method.Find)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func InsertOne[T any](httpReq *http.Request, req domain.InsertOneReq[T]) (result string, err error) {
	resp, err := callSdkService[string](httpReq, &req, method.InsertOne)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func InsertMany[T any](httpReq *http.Request, req domain.InsertManyReq[T]) (result []string, err error) {
	resp, err := callSdkService[[]string](httpReq, &req, method.InsertMany)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func UpdateById[T any](httpReq *http.Request, req domain.UpdateOneReq[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &req, method.UpdateById)
	if err != nil {
		return
	}
	result = resp.Result
	return
}

func UpdateMany[T any](httpReq *http.Request, req domain.UpdateWrapper[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &req, method.UpdateMany)
	if err != nil {
		return
	}
	result = resp.Result
	return
}

func DeleteOne[T any](httpReq *http.Request, query domain.QueryWrapper[T]) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &query, method.DeleteOne)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func DeleteMany[T any](httpReq *http.Request, query domain.DeleteWrapper) (result int, err error) {
	resp, err := callSdkService[int](httpReq, &query, method.DeleteMany)
	if err != nil {
		return
	}
	result = resp.Result

	return
}

func SaveHook(httpReq *http.Request, req domain.SaveHookReq) (err error) {
	_, err = callSdkService[any](httpReq, &req, method.SaveHook)

	return
}

func DeleteHook(httpReq *http.Request, req domain.DeleteHookReq) (err error) {
	_, err = callSdkService[any](httpReq, &req, method.SaveHook)

	return
}

// 调用SDK sevice
func callSdkService[T any](httpReq *http.Request, req domain.SdkServiceReq, m string) (resp domain.CallSdkResp[T], err error) {
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
	request.Header.Add(consts.MethodHeaderKey, m)

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	var buff = make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	// 反序列化结果
	err = json.Unmarshal(buff, &resp)
	if err != nil {
		return
	}

	if resp.ErrCode != consts.Success {
		err = errors.New(resp.ErrMessage)
	}
	return
}
