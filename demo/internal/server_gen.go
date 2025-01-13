package internal

import "github.com/gin-gonic/gin"

type HTTPServer struct {
	e *gin.Engine
}

func NewHttpServer() *HTTPServer {
	s := &HTTPServer{}
	e := gin.Default()
	e.UseH2C = true
	s.e = e
	s.register()
	return s
}

func (s *HTTPServer) Run(addr string) error {
	return s.e.Run(addr)
}

func (s *HTTPServer) register() {
	Register(s.e, NewHelloService())
}
