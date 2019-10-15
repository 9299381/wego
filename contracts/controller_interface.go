package contracts

type IController interface {
	Valid(ctx Context) error
	Handle(ctx Context) (interface{}, error)
}

type Controller struct {
}

func (it *Controller) Handle(ctx Context) (interface{}, error) {
	return nil, nil
}
func (it *Controller) Valid(ctx Context) {

}
