{{.Comment}}

package {{.Pkg}}

import (
{{- range .Imports }}   
{{- if .Alias}} 
    {{.Alias}} {{.Path}}
{{- else}}
    {{.Path}}
{{- end}}
{{- end}}
)

{{- $server := .Server}}
type HTTPServer struct {
    e *gin.Engine
}

func NewHttpServer() *HTTPServer {
	s := new(HTTPServer)
	e := gin.Default()
	s.e = e
	s.register()
	return s
}

func (s *HTTPServer) Run(addr string) error {
	return s.e.Run(addr)
}

func (s *HTTPServer) register() {
{{- range $server.Services}}
	Register{{.Name}}(s.e, New{{.Name}}())
{{- end}}
}
