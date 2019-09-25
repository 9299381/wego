package args

import "flag"

var Name string
var Mode string

var Registy string

var Server string
var Config string

var Cmd string
var Args string

func init() {

	flag.StringVar(&Name, "name", "app", "服务名称")
	flag.StringVar(&Mode, "mode", "dev", "开发模式")
	flag.StringVar(&Registy, "registy", "", "consul服务注册中心")

	flag.StringVar(&Server, "server", "http,event", "需要启动的服务器")

	flag.StringVar(&Config, "config", ".env", "环境配置")

	flag.StringVar(&Cmd, "cmd", "cmd", "cli命令")
	flag.StringVar(&Args, "args", "{}", "json参数")

	if !flag.Parsed() {
		flag.Parse()
	}
}
