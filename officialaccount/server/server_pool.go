package server

import (
	"sync"

	"github.com/silenceper/wechat/v2/officialaccount/context"
)

var serverPool *sync.Pool

// 初始化复用池
func init() {
	serverPool = &sync.Pool{
		New: func() interface{} {
			return &OfficialServer{}
		},
	}
}

// PoolServer 带有复用池的server
type PoolServer struct {
	*OfficialServer
}

// NewServerByPool init
func NewServerByPool(context *context.Context) Server {
	srv := serverPool.Get().(*OfficialServer)
	srv.Context = context
	return &PoolServer{
		OfficialServer: srv,
	}
}

// Close 回收这个 server
func (s *PoolServer) Close() {
	s.Context = nil
	s.Writer = nil
	s.Request = nil
	s.skipValidate = false
	s.openID = ""
	s.messageHandler = nil
	s.RequestRawXMLMsg = nil
	s.RequestMsg = nil
	s.ResponseRawXMLMsg = nil
	s.ResponseMsg = nil
	s.isSafeMode = false
	s.random = nil
	s.nonce = ""
	s.timestamp = 0
	serverPool.Put(s)
}
