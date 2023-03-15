type {{.Name}} struct {
    server {{.Group}}.{{.Name}}
}

func New{{.Name}}() *{{.Name}} {
    return &{{.Name}}{server: {{.Group}}.New{{.Name}}}
}

func (s *{{.Name}}) Error(c *gin.Context, code int, msg string) {
	c.AbortWithError(code, errors.New(msg))
}

func (s *{{.Name}}) ErrorWithResp(c *gin.Context, code int, resp any) {
	c.AbortWithStatusJSON(code, resp)
}

func (s *{{.Name}}) Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

{{$serviceName:=.Name}}
{{- range .Methods}} 
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