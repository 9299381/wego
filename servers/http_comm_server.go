package servers

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/servers/transports"
	"github.com/9299381/wego/tools/errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

type HttpCommServer struct {
	*mux.Router
}

func NewHttpCommServer() *HttpCommServer {
	return &HttpCommServer{
		Router: mux.NewRouter(),
	}
}

func (it *HttpCommServer) Route(method string, path string, endpoint endpoint.Endpoint) {
	it.Methods(method).
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (it *HttpCommServer) Post(path string, endpoint endpoint.Endpoint) {
	it.Methods("POST").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

//
func (it *HttpCommServer) Get(path string, endpoint endpoint.Endpoint) {
	it.Methods("GET").
		Path(path).
		Handler(transports.NewHTTP(endpoint))
}

func (it *HttpCommServer) Load() {
	//注册通用路由
}

func (it *HttpCommServer) Start() error {
	config := (&configs.HttpConfig{}).Load()
	port := config.Get("HttpPort")
	if port != "" {
		wego.App.Logger.Info("Http Server Start ", port)
		handler := it.Router
		return http.ListenAndServe(port, handler)
	}
	return errors.New("9999", "No HttpPort Define")
}
