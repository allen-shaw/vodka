type {{.Name}} interface {
{{range .Methods}}
    {{.Name}}(ctx *gin.Context, req *api.{{.Request}}) (*api.{{.Reply}}, error)
{{end}}
}

func New{{.Name}}() {{.Name}} {
    return &{{.PrivateName}}{}
}

type {{.PrivateName}} struct {
    // TODO: user implement
}

{{$name=.PrivateName}}
{{range .Methods}}
func (s *{{$name}}) {{.Name}}(ctx *gin.Context, req *api.{{.Request}}) (*api.{{.Reply}}, error) {
    // TODO: user implement
	panic("not implement")
}
{{end}}