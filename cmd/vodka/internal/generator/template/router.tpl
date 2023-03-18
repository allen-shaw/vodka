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
{{- range $server.Services}}
func Register{{.Name}}(router gin.IRouter, s *{{.Name}}) {
{{- if .Group }}
    group := router.group("{{.Group}}")
    {{ range .Methods}}
    group.{{.Method}}("{{.Path}}", s.{{.Name}})
    {{- end}} 
{{- else}}
    {{- range .Methods}}
    router.{{.Method}}("{{.Path}}", s.{{.Name}})
    {{- end}}
{{- end}}
}
{{end}}