// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/lj19950508/ddd-demo-go/config"
	"go.uber.org/fx"
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
func New(lc fx.Lifecycle,cfg *config.Config,handler http.Handler) *Server {
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

	// for _, opt := range opts {
	// 	opt(s)
	// }

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s.start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.Shutdown()
			return nil
		},
	})


	return s
}

func (s *Server) start() {
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
func (s *Server) Shutdown() error {
	//有没有办法等所有工作进程跑完再关闭。	
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
