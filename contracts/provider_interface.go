package contracts

type IProvider interface {
	Register() //在这里注册路由
	Boot()     //加载配置文件等
}
