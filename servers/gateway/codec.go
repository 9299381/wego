package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego"
	"github.com/9299381/wego/contracts"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

func decodeRequest(_ context.Context, r *http.Request) (req *contracts.GateWayRequest, err error) {
	data := make(map[string]interface{})
	var vars url.Values
	//r.Body 读取一次就消失了,因此重写了http.request中的parsePostForm
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		vars, err = parsePostForm(r)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	} else {
		_ = r.ParseForm()
		vars = r.Form
	}
	for k, v := range vars {
		data[k] = v[0]
	}
	if _, ok := vars["request_id"]; ok == false {
		data["request_id"] = wego.ID()
	}
	if strings.Index(r.RemoteAddr, "::") > 0 {
		data["client_ip"] = "127.0.0.1"
	} else {
		data["client_ip"] = r.RemoteAddr
	}
	if autoken := r.Header.Get("autoken"); autoken != "" {
		data["autoken"] = autoken
	}
	req = parseUrl(r)
	req.Data = data
	req.Id = data["request_id"].(string)

	return
}

//从http.request 中拷贝修改,目的是body读出,写入
func parsePostForm(r *http.Request) (vs url.Values, err error) {
	if r.Body == nil {
		err = errors.New("missing form body")
		return
	}
	ct := r.Header.Get("Content-Type")
	// RFC 7231, section 3.1.1.5 - empty type
	//   MAY be treated as application/octet-stream
	if ct == "" {
		ct = "application/octet-stream"
	}
	ct, _, err = mime.ParseMediaType(ct)
	switch {
	case ct == "application/x-www-form-urlencoded":
		//备份body
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var reader io.Reader = r.Body
		maxFormSize := int64(1<<63 - 1)
		if _, ok := r.Body.(*maxBytesReader); !ok {
			maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
			reader = io.LimitReader(r.Body, maxFormSize+1)
		}
		b, e := ioutil.ReadAll(reader)
		if e != nil {
			if err == nil {
				err = e
			}
			break
		}
		if int64(len(b)) > maxFormSize {
			err = errors.New("http: POST too large")
			return
		}
		vs, e = url.ParseQuery(string(b))
		if err == nil {
			err = e
		}
		//回写body
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	case ct == "multipart/form-data":
		// handled by ParseMultipartForm (which is calling us, or should be)
		// TODO(bradfitz): there are too many possible
		// orders to call too many functions here.
		// Clean this up and write more tests.
		// request_test.go contains the start of this,
		// in TestParseMultipartFormOrder and others.
	}

	return
}

func parseUrl(r *http.Request) *contracts.GateWayRequest {
	var service, route, dest string
	pathArray := strings.Split(r.URL.Path, "/")
	if len(pathArray) <= 2 {
		//这是本地的
		//如果是health,则返回 SERVING
		service = pathArray[1]
		route = ""
		dest = r.URL.Path
	} else {
		service = pathArray[1]
		route = ""
		for _, v := range pathArray[2:] {
			route += v + "."
		}
		route = route[:len(route)-1]
		dest = "/" + strings.Join(pathArray[2:], "/")
	}
	return &contracts.GateWayRequest{
		Dest:    dest,
		Method:  r.Method,
		Service: service,
		Route:   route,
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)
}
