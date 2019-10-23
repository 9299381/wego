package constants

const Err string = "9999::错误"
const ErrStop string = "9990::暂停服务"
const ErrMQTTConnect string = "9991::MQTT链接失败"
const ErrCacheInit string = "9992::Cache初始化错误"
const ErrLoadEnv string = "9993::ENV环境加载失败"
const ErrSign string = "9994::数据签名错误"

const ErrNotExist string = "9000::数据不存在"
const ErrIsExist string = "9001::数据已存在"
const ErrSaveFailed string = "9002::数据保存失败"

const ErrConvert string = "9010::类型转换错误"
const ErrJson string = "9011::JSON报文解析错误"

const ErrNoToken string = "9020::缺少authToken"
const ErrTokenFmt string = "9021::Token格式错误"
const ErrTokenExp string = "9022::Token过期,请重新登录"
const ErrTokenSign string = "9023::签名错误,请重新登录"

const ErrRoute string = "9030::路由解析错误"
