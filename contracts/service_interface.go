package contracts

type IService interface {
	Next(srv IService) IService
	Handle(ctx Context) error
}
