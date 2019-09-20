# 功能
采用go-kit模式封装
专注编写业务逻辑
# 流程:
~~~~
    1.编写service中间件
    2.组合中间件形成service chain
    3.handler 由 filter(endpoint)和service chain组合形成 
    4.provider中注册handler
    5.server的路由中注册各种handler
    6.支持以下server
~~~~
内置支持Server
~~~~
    1 http
    2 grpc
    3 websocket
    4 cronjob
    5 timer
    6 command
    7 redis queue
~~~~
# 采用组件:
~~~~
	github.com/cespare/xxhash
	github.com/coocood/freecache  内存cache
	        //set
	    	v := make(map[string]interface{})
        	v["aaa"] = "bbb"
        	v["ccc"] = "ddd"
        	wego.Cache("aaaaa",v,60)
        	//get
            v:=wego.Cache("aaaaa")
            d := make(map[string]interface{})
            err:=json.Unmarshal(v,&d)
        	
	github.com/fastly/go-utils 
	github.com/go-kit/kit  //核心组件
	github.com/go-logfmt/logfmt   格式化输出
	
	github.com/go-sql-driver/mysql   mysql
	github.com/go-xorm/xorm          mysql orm
	
	github.com/golang/protobuf      grpc pb协议
	github.com/gomodule/redigo      redis
	github.com/gorilla/mux          http router
	github.com/gorilla/websocket    websocket   
	github.com/jehiah/go-strftime 
	github.com/jonboulle/clockwork  
	github.com/juju/ratelimit       限流endpoint
	github.com/lestrrat/go-envload  
	github.com/lestrrat/go-file-rotatelogs   日志分割
	github.com/lestrrat/go-strftime  
	github.com/rifflock/lfshook             日志hook
	github.com/robfig/cron/v3             cronjob
	github.com/sirupsen/logrus       日志组件
	github.com/tebeka/strftime  
	google.golang.org/grpc           grpc服务
	gopkg.in/check.v1  
~~~~

# 样例

main
~~~~
	wego.Provider(&providers.BootStrap{})
	wego.Provider(&provider.ExamProvider{})
	wego.Router("grpc",&router.GrpcRouter{})
	wego.Router("http",&router.HttpRouter{})
	wego.Router("queue",&router.QueueRouter{})
	wego.Router("command",&router.CommandRouter{})
	wego.Router("websocket",&router.WebSocketRouter{})
	wego.Router("timer",&router.TimerRouter{})
	wego.Router("cron",&router.CronRouter{})
	wego.Start()
~~~~
ExamProvider
~~~~
	wego.Handler("one",filters.Limit(
			services.Chain(
				&service.OneService{},
				&service.TwoService{},
			)),
		)
	wego.Handler("auth", filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.JwtEndpoint{},
		&filters.LimitEndpoint{},
		&filters.CommEndpoint{
			Service: services.Chain(
				&service.AuthService{},
			),},))

	wego.Handler("two",filters.New(services.Chain(&service.TwoService{})))

~~~~
HttpRouter
~~~~
	it.Get("/exam/one", wego.Handler("one"))
	it.Get("/exam/two", wego.Handler("two"))
	it.Post("/exam/auth", wego.Handler("auth"))
~~~~
CronRouter
~~~~
	it.Route("*/5 * * * * *", wego.Handler("one"))
	it.Route("*/2 * * * * *", wego.Handler("two"))
~~~~
GrpcRouter
~~~~
	it.Route("Two", wego.Handler("two"))
~~~~


RedisService
~~~~
func (it *RedisService)Handle(ctx contracts.Context) error  {
	client := wego.Redis() //从pool中获取一个链接
	defer client.Close()   //延时释放链接,本方法执行完毕时释放
	_, _ = client.Do("SET", "go_key", "value")
	res,_ :=redis.String(client.Do("GET","go_key"))
	exists, _ := redis.Bool(client.Do("EXISTS", "foo"))
	if exists {
		ctx.Log.Info("foo 存在")
	}else{
		_, _ = client.Do("SET", "foo", "value")
		ctx.Log.Info("foo 不存在")

	}
	ctx.Log.Info("redis-go_key 的值:", res)
	return it.next.Handle(ctx)
}
~~~~
SqlService
~~~~
func (it *SqlService)Handle(ctx contracts.Context) error  {
	repo := &repository2.UserRepo{Context: ctx}
	user := repo.FetchId("1189164474851006208")
	ctx.Response("user",user)
	ctx.Response("request",ctx.GetValue("request"))
	return it.next.Handle(ctx)
}
~~~~

snowflake id生成器
~~~~
    wego.ID()
~~~~
