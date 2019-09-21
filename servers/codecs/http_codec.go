package codecs

import (
	"context"
	"encoding/json"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"net/http"
	"strings"
)

// HTTP请求数据解码函数
func HttpFormDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	_ = r.ParseForm()
	vars := r.Form
	requestId, ok := vars["request_id"]
	if ok == false {
		requestId = make([]string, 1)
		requestId[0] = wego.ID()
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
	if autoken := r.Header.Get("autoken"); autoken != "" {
		data["autoken"] = autoken
	}

	return contracts.Request{
		Id:   requestId[0],
		Data: data,
	}, nil
}

func HttpJsonDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//把body json直接转换为request
	var request contracts.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func HttpMuxDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//mux.Vars(r)
	//这种方式 适合路由上为  /exam/test/{type}/{value} 模式的解析,
	return nil, nil
}

// HTTP返回数据编码函数
func HttpEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)

}
