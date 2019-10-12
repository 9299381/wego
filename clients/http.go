package clients

import (
	"encoding/json"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/tools/convert"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func NewHttpPostCall(host, service string, params map[string]interface{}) (ret contracts.Response) {

	path := "http://" + host + "/" + strings.Replace(service, ".", "/", -1)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.PostForm(path, convert.FormEncode(params))
	if err != nil {
		ret = contracts.ResponseFailed(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	response := &contracts.Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		ret = contracts.ResponseFailed(err)
	} else {
		m := response.Data.(map[string]interface{})
		m["call_method"] = "http"
		response.Data = m
		ret = *response
	}
	return
}
