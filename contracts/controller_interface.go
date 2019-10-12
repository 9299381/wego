package contracts

type IController interface {
	Valid(ctx Context) error
	Handle(ctx Context) (interface{}, error)
}
