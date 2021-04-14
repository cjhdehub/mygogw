package server

import (
	"net/http"
	"time"
)

func (server *Server) heartbeatHandler(w http.ResponseWriter, req *http.Request) {
	if cs, ok := req.URL.Query()["clientid"]; ok && len(cs[0]) > 0 {
		clientId := cs[0]
		if value, ok := server.Clients.Load(clientId); ok {
			client, _ := value.(*Client)
			client.LastHeartbeatTime = time.Now()
		}
	}
}
