package args

import "flag"

var Name string
var Mode string
var Server string
var Config string
var Registy string

var Cmd string
var Args string

func init() {

	flag.StringVar(&Name, "name", "demo", "服务名称")
	flag.StringVar(&Mode, "mode", "dev", "开发模式")
	flag.StringVar(&Server, "server", "http", "需要启动的服务器")

	flag.StringVar(&Config, "config", ".env", "环境配置")
	flag.StringVar(&Registy, "registy", "", "服务注册中心")

	flag.StringVar(&Cmd, "cmd", "cmd", "cli命令")
	flag.StringVar(&Args, "args", "{}", "json参数")

	flag.Parse()
}
