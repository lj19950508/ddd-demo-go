// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"
)
      

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler,opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}
	
	

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}


//chan是error类型，（从隧道流出到变量里）
// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}
// 从变量传入数据到channel （返回一个可传入的error）
// func (s *Server) Siri() chan<- error {
// 	return s.notify
// }

// Shutdown -.
func (s *Server) Shutdown(rootCtx context.Context) error {
	//有没有办法等所有工作进程跑完再关闭。	
	ctx, cancel := context.WithTimeout(rootCtx, s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
