package contracts

type IController interface {
	GetRules() interface{}
	Handle(ctx Context) (interface{}, error)
}

type Controller struct {
}

func (it *Controller) Handle(ctx Context) (interface{}, error) {
	return nil, nil
}
func (it *Controller) GetRules() interface{} {
	return nil
}
