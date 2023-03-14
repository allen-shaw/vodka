package template

import (
	"fmt"
	"strings"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

// go:embed server.tpl
var serverTpl string

const (
	serverFileName = "sever_gen.go"
	routerFileName = "router_gen.go"
)

const (
	contextPkg = protogen.GoImportPath("context")
	ginPkg     = protogen.GoImportPath("github.com/gin-gonic/gin")
)

const (
	PkgInternal = "internal"
)

const (
	commentTpl = `// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-gin %v
// 	protoc        %v
// source: %v`
)

func GetComment(protocVersion string, source ...string) string {
	src := strings.Join(source, ",")
	return fmt.Sprintf(commentTpl, version.Version, protocVersion, src)
}

type Method struct {
	Name    string // SayHello
	Num     int    // 一个 rpc 方法可以对应多个 http 请求
	Request string // SayHelloReq
	Reply   string // SayHelloResp
	// http_rule
	Path         string // 路由 /hello
	Method       string // HTTP Post/Get
	Body         string
	ResponseBody string
}

type Service struct {
	Name        string // GreeterService
	PrivateName string // greeterService

	FullName string // helloworld.Greeter
	FilePath string // api/helloworld/helloworld.proto

	// HTTP properties
	Group     string // greeter
	Methods   []*Method
	MethodSet map[string]*Method
}

type Server struct {
	Package  string
	FileName string
	Imports  []protogen.GoImportPath
	Services []*Service
}

type RouterGenerator struct {
	Package  string
	FileName string // route_gen.go
	Imports  []protogen.GoImportPath
	Servers  []*Server
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
