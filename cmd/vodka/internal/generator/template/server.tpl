type HTTPServer struct {
    e *gin.Engine
}

func NewHttpServer() *HTTPServer {
	s := &HTTPServer{}
	e := gin.Default()
	s.e = e
	s.register()
	return s
}

func (s *HTTPServer) Run(addr string) error {
	return s.e.Run(addr)
}

func (s *HTTPServer) register() {
{{range .Services}}
	Register(s.e, New{{.Name}}())
{{end}}
}
