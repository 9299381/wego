package contracts

type IService interface {
	Handle(ctx Context) error
}
