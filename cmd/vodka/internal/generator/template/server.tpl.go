package template

import "google.golang.org/protobuf/compiler/protogen"

// go:embed server.tpl
var serverTpl string

const (
	serverFileName = "sever_gen.go"
)

type Server struct {
	Package  string
	FileName string
	Imports  []protogen.GoImportPath
	Services []*Service
}

func NewServer(services []*Service) *Server {
	return &Server{
		Package:  PkgInternal,
		FileName: serverFileName,
		Imports:  []protogen.GoImportPath{ginPkg},
		Services: services,
	}
}

func (s *Server) Gen(gen *protogen.Plugin) {
	// g := gen.NewGeneratedFile(s.FileName, s.Imports)

	// comment := GetComment()
	// g.p()
}
