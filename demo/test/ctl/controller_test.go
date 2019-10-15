package ctl

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/src/controller"
	"github.com/9299381/wego/demo/src/provider"
	"github.com/9299381/wego/tools"
	"testing"
)

func TestOneController(t *testing.T) {
	wego.Provider(&provider.DemoProvider{})
	resp, err := tools.Test().
		Controller(&controller.OneController{}).
		Request(nil).
		Run()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Data)
	}

}
func TestTwoController(t *testing.T) {
	wego.Provider(&provider.DemoProvider{})
	resp, err := tools.Test().
		Controller(&controller.TwoController{}).
		Request(nil).
		Run()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Data)
	}

}
func TestParallelController(t *testing.T) {
	wego.Provider(&provider.DemoProvider{})
	resp, err := tools.Test().
		Controller(&controller.ParallelController{}).
		Request(nil).
		Run()

	if err != nil {
		t.Error(err)
	} else {
		t.Log(resp.Data)
	}

}
