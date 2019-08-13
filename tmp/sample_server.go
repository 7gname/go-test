package tmp

import "time"

type Server struct {
	Addr    string
	timeout time.Duration
}

func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	srv := &Server{
		Addr: addr,
	}
	for _, option := range options {
		option(srv)
	}
	return srv, nil
}

func Timeout(d time.Duration) func(*Server) {
	return func(srv *Server) {
		srv.timeout = d
	}
}
