package contracts

type Iconfig interface {
	Load() Iconfig
	Get(key string) string
}
