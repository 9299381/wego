package contracts

type IRouter interface {
	Boot()
	Load()
	Register()
	Start() error
	Close()
}
