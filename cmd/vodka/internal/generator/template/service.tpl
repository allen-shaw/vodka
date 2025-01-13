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
{{- $service := .Service}}
type {{$service.Name}} struct {
    server {{$service.Prefix}}.{{$service.Name}}
}

func New{{$service.Name}}() *{{$service.Name}} {
    return &{{$service.Name}}{server: {{$service.Prefix}}.New{{$service.Name}}}
}

func (s *{{$service.Name}}) Error(c *gin.Context, code int, msg string) {
	c.AbortWithError(code, errors.New(msg))
}

func (s *{{$service.Name}}) ErrorWithResp(c *gin.Context, code int, resp any) {
	c.AbortWithStatusJSON(code, resp)
}

func (s *{{$service.Name}}) Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

{{$serviceName:=$service.Name}}
{{- range $service.Methods}} 
func (s *{{$serviceName}}) {{.Name}}(c *gin.Context) {
    var req api.{{.Request}}
    {{ if eq .Method "Post"}}
    err := c.ShouldBindJSON(&req)
    {{- else if eq .Method "Get"}}
    err := c.ShouldBindQuery(&req)
    {{- end}}
	if err != nil {
		s.Error(c, http.StatusBadRequest, "invalid request")
		return
	}
    
    resp, err := s.server.{{.Name}}(c, &req)
	if err != nil {
		s.ErrorWithResp(c, http.StatusInternalServerError, resp)
	}

	s.Success(c, resp)
}
{{end}}