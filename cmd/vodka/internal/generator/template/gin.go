package template

import "google.golang.org/protobuf/compiler/protogen"

type Method struct {
	Name    string // SayHello
	Group   string // greeter  // TODO:还没支持
	Num     int    // 一个 rpc 方法可以对应多个 http 请求
	Request string // SayHelloReq
	Reply   string // SayHelloResp
	// http_rule
	Path         string // 路由 /hello
	Method       string // HTTP Post/Get
	Body         string
	ResponseBody string
	// 元数据
	ProtoMethod *protogen.Method
}

type Service struct {
	Name        string // GreeterService
	PrivateName string // greeterService
	FullName    string // helloworld.Greeter
	FilePath    string // api/helloworld/helloworld.proto
	Prefix      string
	// HTTP properties
	Methods    []*Method
	MethodSet  map[string]*Method
	Deprecated bool
	// 元数据
	ProtoFile    *protogen.File
	ProtoService *protogen.Service
}

type Server struct {
	Services []*Service
}
