package di_test

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/test/di_test/demo"
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	str := "aaaaaaa"
	people := &demo.Man{}
	wego.DI().MapTo(people, (*demo.IPeople)(nil))
	ret, _ := wego.DI().Invoke(func(p demo.IPeople) string {
		return p.Write(str)
	})
	t.Log(ret)
}

func TestTwo(t *testing.T) {
	str := "aaaaaaa"
	people := &demo.Man{}
	wego.DI().Map(people)
	ret, _ := wego.DI().Invoke(func(p *demo.Man) string {
		return p.Write(str)
	})
	t.Log(ret)
}

func TestThree(t *testing.T) {
	people := &demo.Man{}
	wego.DI().MapTo(people, (*demo.IPeople)(nil))
	_, _ = wego.DI().Invoke(MyFastInvoker(nil))
}

type MyFastInvoker func(people demo.IPeople)

func (invoker MyFastInvoker) Invoke(args []interface{}) ([]reflect.Value, error) {
	if people, ok := args[0].(demo.IPeople); ok {
		people.Work()
	}
	return nil, nil
}
