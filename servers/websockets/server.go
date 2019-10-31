package websockets

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/constants"
	"github.com/9299381/wego/contracts"
	"github.com/9299381/wego/servers/commons"
	"github.com/gorilla/websocket"
	"net/http"
)

type Server struct {
	handlers map[string]*commons.CommHandler
	ctx      context.Context
	Logger   contracts.ILogger
}

func NewServer() *Server {
	s := &Server{
		ctx:      context.Background(),
		handlers: make(map[string]*commons.CommHandler),
	}
	return s
}

func (s *Server) Register(name string, handler *commons.CommHandler) {
	s.handlers[name] = handler

}

func (s *Server) Serve() error {
	config := configs.LoadWebSocketConfig()
	address := config.WebSocketHost + ":" + config.WebSocketPort
	s.Logger.Info("WebSocket Server Start ", address)
	http.HandleFunc(config.Path, s.wsHandler)
	return http.ListenAndServe(address, nil)
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {

	var socket = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := socket.Upgrade(w, r, nil)
	if err != nil {
		s.Logger.Debug("upgrade err:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			s.Logger.Debug("read message err:", err)
			break
		}
		payload := &contracts.Payload{}
		err = json.Unmarshal(message, payload)
		if err != nil {
			e := s.faild(constants.ErrJson)
			err = c.WriteMessage(mt, e)
			break
		}
		handler, isExist := s.handlers[payload.Route]
		if isExist == false {
			e := s.faild(constants.ErrRoute)
			err = c.WriteMessage(mt, e)
			break
		}
		ctx := context.Background()
		//请求服务
		response, errResp := handler.Handle(ctx, payload.Params)

		if errResp != nil {
			e := s.faild(errResp.Error())
			err = c.WriteMessage(mt, e)
			break
		}
		resp, _ := json.Marshal(response)
		err = c.WriteMessage(mt, resp)
		if err != nil {
			s.Logger.Debug("write err:", err)
			break
		}
	}
}

func (s *Server) faild(message string) []byte {
	response := contracts.ResponseFailed(errors.New(message))
	ret, _ := json.Marshal(response)
	return ret
}

func (s *Server) Close() {

}
