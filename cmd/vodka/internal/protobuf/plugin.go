package protobuf

import (
	"flag"
	"net/http"
	"strings"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/generator/template"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/protobuf/log"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/util"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	srv        *template.Server
	methodSets map[string]int

	options Options
}

func NewPlugin(vodka bool) *Plugin {
	p := &Plugin{
		methodSets: make(map[string]int, 0),
	}
	p.srv = template.NewServer(make([]*template.Service, 0))
	return p
}

func (p *Plugin) Run() int {
	var flags flag.FlagSet
	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		// 0. 版本
		// 0.1 protoc版本
		// 0.2 protoc-gen-gin 版本

		// 1. 遍历所有文件
		// 每个文件生成Service对象

		// 2. 生成Server、router
		// 1.1. 当前有哪些source files
		// 1.2. 有哪些Service

		for _, f := range gen.Files {
			log.Debug("file:%v\n", f.Desc.Path())
			if !f.Generate {
				continue
			}
			svcs := p.analysisFile(f)
			p.srv.Services = append(p.srv.Services, svcs...)
		}

		p.genFiles(gen)

		return nil
	})

	return 0
}

func (p *Plugin) analysisFile(file *protogen.File) []*template.Service {
	log.Debug("analysis file: %v\n", file.Desc.Path())

	if len(file.Services) == 0 {
		return nil
	}

	ss := make([]*template.Service, 0)

	for _, service := range file.Services {
		s := p.analysisService(file, service)
		ss = append(ss, s)
	}

	return ss
}

func (p *Plugin) analysisService(file *protogen.File, s *protogen.Service) *template.Service {
	log.Debug("analysis service: %+v .\n", s)

	svc := new(template.Service)
	if s.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		svc.Deprecated = true
	}

	svc.Name = s.GoName
	svc.PrivateName = util.FirstLower(svc.Name)
	svc.FullName = string(s.Desc.FullName())
	svc.FilePath = file.Desc.Path()

	for _, method := range s.Methods {
		m := p.analysisMethod(method)
		svc.Methods = append(svc.Methods, m...)
	}

	log.Debug("[Done] analysis service: %+v .\n", svc)

	return svc
}

func (p *Plugin) analysisMethod(m *protogen.Method) []*template.Method {
	log.Debug("analysis method: %+v .\n", m)

	methods := make([]*template.Method, 0)

	// 解析http rule
	rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
	if ok && rule != nil {
		for _, bind := range rule.AdditionalBindings {
			methods = append(methods, p.buildHTTPRule(m, bind))
		}
		methods = append(methods, p.buildHTTPRule(m, rule))
		return methods
	}

	// 不配置rule则使用默认配置
	methods = append(methods, p.defaultMethod(m))

	log.Debug("[Done] analysis method: %+v .\n", methods)

	return methods
}

// defaultMethod 根据函数名生成 http 路由
// 例如: GetBlogArticles ==> get: /blog/articles
// 如果方法名首个单词不是 http method 映射，那么默认返回 POST
func (p *Plugin) defaultMethod(m *protogen.Method) *template.Method {
	log.Debug("default method: %+v .\n", m)

	names := strings.Split(util.SnakeCase(m.GoName), "_")
	var (
		paths      []string
		httpMethod string
		path       string
	)

	switch strings.ToUpper(names[0]) {
	case http.MethodGet, "FIND", "QUERY", "LIST", "SEARCH":
		httpMethod = http.MethodGet
	case http.MethodPost, "CREATE":
		httpMethod = http.MethodPost
	case http.MethodPut, "UPDATE":
		httpMethod = http.MethodPut
	case http.MethodPatch:
		httpMethod = http.MethodPatch
	case http.MethodDelete:
		httpMethod = http.MethodDelete
	default:
		httpMethod = "POST"
		paths = names
	}

	if len(paths) > 0 {
		path = strings.Join(paths, "/")
	}

	if len(names) > 1 {
		path = strings.Join(names[1:], "/")
	}

	md := p.buildMethodDesc(m, httpMethod, path)
	md.Body = "*"

	log.Debug("[Done] analysis A Method: %+v .\n", md)

	return md
}

func (p *Plugin) buildHTTPRule(m *protogen.Method, rule *annotations.HttpRule) *template.Method {
	log.Debug("build http rule method: %+v .\n", m)

	var (
		path   string
		method string
	)

	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = "GET"
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = "PUT"
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = "POST"
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = "DELETE"
	case *annotations.HttpRule_Patch:
		path = pattern.Patch
		method = "PATCH"
	case *annotations.HttpRule_Custom:
		path = pattern.Custom.Path
		method = pattern.Custom.Kind
	}

	md := p.buildMethodDesc(m, method, path)

	log.Debug("[Done] analysis A Method: %+v .\n", md)

	return md
}

func (p *Plugin) buildMethodDesc(m *protogen.Method, httpMethod, path string) *template.Method {
	defer func() {
		p.methodSets[m.GoName]++
	}()

	md := &template.Method{
		Name:    m.GoName,
		Group:   "",
		Num:     p.methodSets[m.GoName],
		Request: m.Input.GoIdent.GoName,
		Reply:   m.Output.GoIdent.GoName,
		Path:    path,
		Method:  util.PascalCase(httpMethod),
	}

	return md
}

func (p *Plugin) genFiles(gen *protogen.Plugin) {
	// 1. 生成api -> ./internal/api/${serviceShortName}/${service_name}.go  // 允许用户修改
	// 2. 生成model -> ./internal/model/api/${serviceShortName}/${service_name}.pb.go
	// 3. 生成service -> ./internal/${serviceShortName}_gin.pb.go

	for _, s := range p.srv.Services {
		s.GenApi()
		s.GenService()
	}

	// 4. 生成router -> ./internal/router_gen.go
	// 5. 生成server -> ./internal/server_gen.go
	p.srv.GenRouter()
	p.srv.GenServer()
}
