package codecs

import (
	"context"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/idwork"
	"net/http"
	"strings"
)

func WebDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	_ = r.ParseForm()
	vars := r.Form
	requestId, ok := vars["request_id"]
	if ok == false {
		requestId = make([]string, 1)
		requestId[0] = idwork.ID()
	}
	data := make(map[string]interface{})
	for k, v := range vars {
		data[k] = v[0]
	}
	if strings.Index(r.RemoteAddr, "::") > 0 {
		data["client_ip"] = "127.0.0.1"
	} else {
		data["client_ip"] = r.RemoteAddr
	}
	// web 方式的验证 需要与html配合,采用 注意js端
	if authToken := r.Header.Get("authToken"); authToken != "" {
		data["authToken"] = authToken
	}
	return contracts.Request{
		Id:   requestId[0],
		Data: data,
	}, nil
}

// HTTP返回数据编码函数
func WebEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	//这里ResponseWriter 对 response 进行渲染操作

	return nil
}
