{{range .Services}}
func Register{{.Name}}(router gin.IRouter, s *{{.Name}}) {
    group := router.group("{{.Group}}")
    {{range .Methods}}
    group.{{.Method}}("{{.Path}}", s.{{.Name}})
    {{end}}
}
{{end}}