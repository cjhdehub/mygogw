package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"gogw-server/common"
	"gogw-server/logger"
	"gogw-server/schema"
)

type Server struct {
	ServerAddr    string
	Clients *sync.Map

	TimeoutSecond time.Duration
}

func NewServer(serverAddr string, timeoutSecond int) *Server {
	return &Server{
		ServerAddr:    serverAddr,
		Clients:       &sync.Map{},
		TimeoutSecond: time.Second * time.Duration(timeoutSecond),
	}
}

//client register
func (s *Server) registerHandler(w http.ResponseWriter, req *http.Request) {
	defer func(){
		req.Body.Close()
	}()

	msgPack, err := schema.ReadMsg(req.Body)
	if err != nil {
		logger.Error(err)
		return
	}

	msg, ok := msgPack.Msg.(*schema.RegisterRequest)
	if ! ok {
		return
	}
	logger.Info(msg.VP)

	clientId := common.UUID("clientid")
	//todo portam
	client := NewClient(
		clientId,
		req.RemoteAddr,
		8800,
		"reverse",
		"tcp",
		req.RemoteAddr + ":22",
		"webservice",
		true,
		"http1.1",
	)

	s.Clients.Store(clientId, client)
	defer func(){
		if err != nil {
			s.Clients.Delete(clientId)
		}
	}()

	if err = client.Start(); err != nil {
		return
	}

	msgPack = & schema.MsgPack {
		MsgType: schema.MSG_TYPE_REGISTER_RESPONSE,
		Msg: & schema.RegisterResponse {
			ClientId: clientId,
			Status: schema.STATUS_SUCCESS,
		},
	}

	err = schema.WriteMsg(w, msgPack)
}

//msg to client 
func (s *Server) msgHandler(w http.ResponseWriter, req *http.Request) {
	defer func(){
		req.Body.Close()
	}()

	if its, ok := req.URL.Query()["clientid"]; ok && len(its[0]) > 0 {
		clientId := its[0]
		if ci, ok := s.Clients.Load(clientId); ok {
			client, _ := ci.(*Client)
			client.HttpHandler(w, req)
		}
	}
}

func (s *Server) cleanerLoop() {
	for {
		t := time.Now()
		shouldDelete := []string{}
		s.Clients.Range(func (k, v interface{}) bool {
			client, _ := v.(*Client)
			if t.Sub(client.LastHeartbeatTime).Milliseconds() > s.TimeoutSecond.Milliseconds() {
				shouldDelete = append(shouldDelete, client.ClientId)
				client.Stop()
			}
			return true
		})

		for _, clientId := range shouldDelete {
			s.Clients.Delete(clientId)
		}

		time.Sleep(time.Second * 10)
	}
}

func (s *Server) Start() {
	logger.Info(fmt.Sprintf("\nserver start\nAddr: %v\n", s.ServerAddr))

	//start client cleaner
	go s.cleanerLoop()

	http.HandleFunc("/register", s.registerHandler)
	http.HandleFunc("/msg", s.msgHandler)
	http.HandleFunc("/heartbeat", s.heartbeatHandler)
	http.ListenAndServe(s.ServerAddr, nil)
}
