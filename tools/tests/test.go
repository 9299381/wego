package tests

import (
	"context"
	"errors"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/filters"
	"github.com/9299381/wego/tools/idwork"
)

type TestStruct struct {
	controller contracts.IController
	request    map[string]interface{}
}

func NewTest() *TestStruct {
	return &TestStruct{
		request: make(map[string]interface{}),
	}
}
func (it *TestStruct) Controller(controller contracts.IController) *TestStruct {
	it.controller = controller
	return it
}
func (it *TestStruct) Request(m map[string]interface{}) *TestStruct {
	if m != nil {
		it.request = m
	}
	return it
}
func (it *TestStruct) Run() (contracts.Response, error) {
	e := filters.New(it.controller)
	request := contracts.Request{
		Id:   idwork.ID(),
		Data: it.request,
	}
	response, err := e(context.Background(), request)
	resp := response.(contracts.Response)
	if err != nil {
		return resp, err
	}
	if resp.Code != "0000" {
		return resp, errors.New(resp.Message)
	}
	return resp, nil
}
