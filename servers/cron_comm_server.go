package servers

import (
	"github.com/9299381/wego/servers/commons"
	"github.com/9299381/wego/servers/cronjobs"
	"github.com/9299381/wego/servers/transports"
	"github.com/go-kit/kit/endpoint"
)

type CronCommServer struct {
	*cronjobs.Server
}

func (it *CronCommServer) Desc() {
	//1）星号(*)
	//表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
	//2）斜线(/)
	//表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15
	//3）逗号(,)
	//用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
	//4）连字号(-)
	//表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
	//5）问号(?)
	//只用于日(Day of month)和星期(Day of week)，\表示不指定值，可以用于代替 *
	//每隔5秒执行一次：*/5 * * * * ?
	//每隔1分钟执行一次：0 */1 * * * ?
	//每天23点执行一次：0 0 23 * * ?
	//每天凌晨1点执行一次：0 0 1 * * ?
	//每月1号凌晨1点执行一次：0 0 1 1 * ?
	//在26分、29分、33分执行一次：0 26,29,33 * * * ?
	//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
}

func NewCronCommServer() *CronCommServer {
	ss := &CronCommServer{
		Server: cronjobs.NewServer(),
	}
	return ss
}

func (it *CronCommServer) Route(spec string, endpoint endpoint.Endpoint) {

	handler := &commons.CommHandler{
		Handler: transports.NewCronJob(endpoint),
	}
	it.Register(spec, handler)
}

func (it *CronCommServer) Load() {
	//通用加载 todo
}

func (it *CronCommServer) Start() error {

	return it.Serve()
}
func (it *CronCommServer) Close() {
	it.Server.Close()
}
