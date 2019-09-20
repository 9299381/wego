package contracts

type IReader interface {
	Read(filePath string) interface{}
}
