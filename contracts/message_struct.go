package contracts

import "strings"

type Request struct {
	Id   string `json:"request_id"`
	Data map[string]interface{}
}

type Response struct {
	Ret     int         `json:"ret"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func MakeResponse(data interface{}, err error) Response {
	if err != nil {
		return ResponseFailed(err)
	} else {
		return ResponseSucess(data)
	}
}
func ResponseSucess(data interface{}) Response {
	return Response{
		Code:    "0000",
		Data:    data,
		Ret:     200,
		Message: "请求成功",
	}
}
func ResponseFailed(err error) Response {
	errMap := strings.Split(err.Error(), "::")
	if len(errMap) == 2 {
		return Response{
			Code:    errMap[0],
			Data:    make(map[string]interface{}),
			Ret:     200,
			Message: errMap[1],
		}
	} else {
		return Response{
			Code:    "9999",
			Data:    make(map[string]interface{}),
			Ret:     200,
			Message: err.Error(),
		}
	}
}

type Payload struct {
	Route  string                 `json:"route"`
	Params map[string]interface{} `json:"params"`
}

type GateWayRequest struct {
	Dest    string
	Method  string
	Id      string
	Service string
	Route   string
	Data    map[string]interface{}
}
