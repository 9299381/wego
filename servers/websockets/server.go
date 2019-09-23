package websockets

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	handlers map[string]*commons.CommHandler
	ctx      context.Context
	Logger   *logrus.Logger
}

func NewServer() *Server {
	s := &Server{
		ctx:      context.Background(),
		handlers: make(map[string]*commons.CommHandler),
	}
	return s
}

func (it *Server) Register(name string, handler *commons.CommHandler) {
	it.handlers[name] = handler

}

func (it *Server) Serve() error {
	config := (&configs.WebSocketConfig{}).Load().(*configs.WebSocketConfig)
	wego.App.Logger.Info("WebSocket Server Start ", config.WebSocketPort)
	http.HandleFunc(config.Path, it.wsHandler)
	return http.ListenAndServe(config.WebSocketPort, nil)
}

func (it *Server) wsHandler(w http.ResponseWriter, r *http.Request) {

	var socket = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := socket.Upgrade(w, r, nil)
	if err != nil {
		it.Logger.Debug("upgrade err:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			it.Logger.Debug("read message err:", err)
			break
		}
		payload := &contracts.Payload{}
		err = json.Unmarshal(message, payload)
		if err != nil {
			e := it.faild(constants.ErrJson)
			err = c.WriteMessage(mt, e)
			break
		}
		handler, isExist := it.handlers[payload.Route]
		if isExist == false {
			e := it.faild(constants.ErrRoute)
			err = c.WriteMessage(mt, e)
			break
		}
		ctx := context.Background()
		//请求服务
		response, errResp := handler.Handle(ctx, payload.Params)

		if errResp != nil {
			e := it.faild(errResp.Error())
			err = c.WriteMessage(mt, e)
			break
		}
		resp, _ := json.Marshal(response)
		err = c.WriteMessage(mt, resp)
		if err != nil {
			it.Logger.Debug("write err:", err)
			break
		}
	}
}

func (it *Server) faild(message string) []byte {
	response := contracts.ResponseFaile(errors.New(message))
	ret, _ := json.Marshal(response)
	return ret
}
