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
type {{$service.Name}} interface {
{{- range $service.Methods}}
    {{.Name}}(ctx *gin.Context, req *api.{{.Request}}) (*api.{{.Reply}}, error)
{{- end}}
}

func New{{$service.Name}}() {{$service.Name}} {
    return &{{$service.PrivateName}}{}
}

type {{$service.PrivateName}} struct {
    // TODO: user implement
}
{{$name := $service.PrivateName}}
{{- range $service.Methods}}
func (s *{{$name}}) {{.Name}}(ctx *gin.Context, req *api.{{.Request}}) (*api.{{.Reply}}, error) {
    // TODO: user implement
	panic("not implement")
}
{{end}}