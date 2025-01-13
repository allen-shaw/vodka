package template

import (
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/stvp/assert"
)

var (
	method11 = &Method{
		Name:    "Hello11",
		Request: "HelloReq11",
		Reply:   "HelloResp11",
		Path:    "/hello11",
		Method:  "Get",
	}
	method12 = &Method{
		Name:    "Hello12",
		Request: "HelloReq12",
		Reply:   "HelloResp12",
		Path:    "/hello12",
		Method:  "Post",
	}
	// method21 = &Method{
	// 	Name:    "Hello21",
	// 	Request: "HelloReq21",
	// 	Reply:   "HelloResp21",
	// 	Path:    "/hello21",
	// 	Method:  "Get",
	// }

	service1 = &Service{
		Name:        "Hello1Service",
		PrivateName: "hello1Service",
		FullName:    "hello.hello1Service",
		FilePath:    "./hello1.proto",

		Methods: []*Method{method11, method12},
		MethodSet: map[string]*Method{
			method11.Name: method11,
			method12.Name: method12,
		},
	}
	// service2 = &Service{
	// 	Name:        "Hello2Service",
	// 	PrivateName: "hello2Service",
	// 	FullName:    "hello.hello2Service",
	// 	FilePath:    "./hello2.proto",
	// 	Methods:     []*Method{method21},
	// 	MethodSet: map[string]*Method{
	// 		method21.Name: method21,
	// 	},
	// }
	// server = NewServer([]*Service{service1, service2})
)

// func TestServerTpl(t *testing.T) {
// 	file := "server.tpl"

// 	t1, err := template.ParseFiles(file)
// 	assert.Nil(t, err)

// 	t1.Execute(os.Stdout, server)
// }

// func TestRouterTpl(t *testing.T) {
// 	file := "router.tpl"

// 	t1, err := template.ParseFiles(file)
// 	assert.Nil(t, err)

// 	t1.Execute(os.Stdout, server)
// }

func TestServiceTpl(t *testing.T) {
	file := "service.tpl"

	t1, err := template.ParseFiles(file)
	assert.Nil(t, err)

	t1.Execute(os.Stdout, service1)
}

func TestApiTpl(t *testing.T) {
	file := "api.tpl"

	t1, err := template.ParseFiles(file)
	assert.Nil(t, err)
	t1.Execute(os.Stdout, service1)
}

func TestGoImport(t *testing.T) {
	id := ginPkg.Ident("g")
	fmt.Println(id.String())
}

func TestEmbed(t *testing.T) {
	fmt.Println("123123123")
	fmt.Println("11111111111111", apiTpl)
}
