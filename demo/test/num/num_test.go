package num

import (
	"github.com/9299381/wego/clients"
	"github.com/9299381/wego/demo/src/model"
	"github.com/9299381/wego/repos"
	"github.com/9299381/wego/tools/idwork"
	"testing"
)

func TestNum(t *testing.T) {

}

func TestDemoTestModel(t *testing.T) {
	err := clients.DB().Sync2(new(model.CommDemoModel))
	if err != nil {
		t.Error(err)
	}
}
func TestInsert(t *testing.T) {
	demo := &model.CommDemoModel{
		Id:      idwork.ID(),
		NumInt1: 9876543210,
		NumInt2: 0,
		NumBig:  1111111111111111111,
	}
	repos.Insert(demo)
}
func TestUpdate(t *testing.T) {
	demo := &model.CommDemoModel{Id: "1298580054575415296"}
	_ = repos.First(demo)
	demo.NumInt1 = 0
	demo.NumInt2 = 123
	_, _ = repos.DB().
		Id(demo.Id).
		Cols("num_int1", "num_int2").
		Update(demo)

}
