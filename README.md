# 功能
采用go-kit模式封装
专注编写业务逻辑
# 流程  步骤
~~~~
    1 编写controller负责, 请求验证,响应,swagger格式,service中间件调用
    2.handler 由 filter(endpoint)和controller组合形成 
    3.provider中注册handler
    4.server的路由中注册各种handler
    5.支持以下server
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
    8 内部event server
    9 gateway server
    10 mqtt订阅server
~~~~
举例
~~~~    
    //经过jwt认证后的用户id,和name
	fmt.Println(ctx.Get("request.claim.Id"))
	fmt.Println(ctx.Get("request.claim.Name"))
    //cache使用
    v, _ := cache.Get("aaaaa")
	v := make(map[string]interface{})
	v["aaa"] = "bbb"
	v["ccc"] = "ddd"
	_ = cache.Set("aaaaa", v, 60)
	//日志使用
	ctx.Log.Info("one....")
	ctx.Log.Infof(format,arg...)
	//请求参数
	dto := ctx.Request().(*dto.Request)
	//redis使用
    client := clients.Redis() //从pool中获取一个链接
    defer client.Close()      //延时释放链接,本方法执行完毕时释放
    _, _ = client.Do("SET", "go_key", "value")
    //mysql使用
    user := model.CommUser{Id: id}
    has, _ := clients.DB().Get(&user)
    //event使用
	params := make(map[string]interface{})
	payload := &contracts.Payload{
		Route:  "two", ->接收处理的handler
		Params: params,
	}
	events.Fire(payload)
	//redis queue使用 默认db->1
    msg := make(map[string]interface{})
    msg["aaa"] = "bbb"
    err := queues.Fire(
        "demo1",     ->发送的redis 队列
        "queue_test",  ->侦听队列的server需要处理的路由handler
        msg,
    )
    //远程服务调用,// 为现有php模式而封装
    params:=make(map[string]interface{})
    params["test_rpc_post"] = "test_rpc_post"
    resp := clients.
        Micro("consul_demo").    //服务的名称
        Service("demo.post").    //服务的注册的handler
        Params(params).
        Run()
    该方法会从consul中获取注册的服务,并随机选择一个进行请求,支持grpc和http post
    http post 对应的远端路由为 http:/host+port/demo/post
~~~~

# 采用组件:
~~~~
	github.com/cespare/xxhash
	github.com/coocood/freecache  内存cache	
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
    github.com/hashicorp/consul v1.6.1 //  consul http grpc 服务注册,发现
    github.com/hashicorp/consul/api v1.2.0
~~~~

# 样例

main
~~~~
    //如果参数配置了registy,则自动进行consul的服务注册 grpc http 都可
	//例如go run main.go -name=test_service  -registy=127.0.0.1:8500 -server=grpc
	wego.Provider(&providers.ConsulRegistyProvider{})
    
    //这里注册自己的handler
	wego.Provider(&provider.ExamProvider{})
	
	//下面的server,根据启动args参数决定
	wego.Router("grpc",&router.GrpcRouter{})
	wego.Router("http",&router.HttpRouter{})
	wego.Router("queue",&router.QueueRouter{})
	wego.Router("command",&router.CommandRouter{})
	wego.Router("websocket",&router.WebSocketRouter{})
	wego.Router("timer",&router.TimerRouter{})
	wego.Router("cron",&router.CronRouter{})
	
	//内置加载事件服务,无需路由,直接调用  handler
    wego.Router("event", servers.NewEventCommServer())
	
	wego.Start()
~~~~
ExamProvider
~~~~
	wego.Handler("one", filters.Limit(&controller.OneController{}))
	wego.Handler("two", filters.New(&controller.TwoController{}))

~~~~

OneController 请求,响应,swagger,调用service中间件chain
~~~~

type OneController struct {
}

//swagger:route GET .....
func (it *OneController) Handle(ctx contracts.Context) (interface{}, error) {
    
	chain := services.Chain(
		&service.OneService{},
		&service.TwoService{},
	)
	_ = chain.Handle(ctx)

	ret := &FirstResp{
		Id:       idwork.ID(),
		UserName: ctx.Get("k.a").(string),
	}
	return ret, nil
}

// swagger:parameters .....
type FirstRequest struct {
	Param1 string `json:"param_1"`
	Param2 int    `json:"param_2"`
}

// swagger:response .....
type FirstResp struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
}
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
	client := clients.Redis() //从pool中获取一个链接
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
	return nil
}
~~~~
SqlService
~~~~
func (it *SqlService)Handle(ctx contracts.Context) error  {
	repo := &repository2.UserRepo{Context: ctx}
	user := repo.FetchId("1189164474851006208")
	ctx.Set("user",user)
	return nil
}
~~~~
struct validations 由 beego的validations 修改而来
github.com/astaxie/beego/validation
~~~~
	req := ctx.Get("request")
	st := &dto.TestDto{}
	err := convert.Map2Struct(req, st)
	if err != nil {
		return err
	}
	err = validations.Valid(st)

type TestDto struct {
	Name string `json:"name" valid:"Required;MinSize(1);MaxSize(5)"`
	Age  int    `json:"age" valid:"Required"`
	//Name   string `json:"name" valid:"Required;Match(/^wego.*/)"` // Name 不能为空并且以 wego 开头
	////有问题
	//Age    string    `json:"age" valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
	//Email  string `json:"email" valid:"Email; MaxSize(100)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	//Mobile string `json:"mobile" valid:"Mobile"` // Mobile 必须为正确的手机号
	//IP     string `json:"ip" valid:"IP"` // IP 必须为一个正确的 IPv4 地址
	Desc string `json:"desc" valid:"Required;Custom(CheckDesc)"` //自定义处理
}

//自定义方法
func (it *TestDto) CheckDesc(v *validations.Validation) {
	if strings.Index(it.Desc, "desc") != -1 {
		_ = v.SetError("Desc", "名称里不能含有 desc")
	}
}

//全部通过后最后执行
func (it *TestDto) Finish(v *validations.Validation) {
	if strings.Index(it.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		_ = v.SetError("Name", "名称里不能含有 admin")
	}
}


~~~~
snowflake id生成器
~~~~
    idwork.ID()
~~~~

command
~~~~
    main -cmd="注册的路由" -args="json参数"
~~~~
request过程
~~~~
    request
    server
        decode
        route(handler)
            endpoint.....
                controller
                    service ......(并行,串行运行)
        encode
    response
~~~~

url
~~~
    https://goswagger.io
~~~


