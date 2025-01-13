package handler

import (
	"github.com/allen-shaw/vodka/cmd/vodka/internal/args"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/generator"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/meta"
)

type initHandler struct {
	root string
	idls []string

	layout *generator.Layout
	code   *generator.Code
}

func newInitHandler(arg *args.Init) *initHandler {
	h := &initHandler{
		root: arg.Out,
		idls: arg.IDLs,
	}
	h.layout = generator.NewLayout(h.root)
	h.code = generator.NewCode(h.root, h.idls)

	h.checkFiles()
	return h
}

func (h *initHandler) checkFiles() {
	// TODO: to implement
	// 检查root是否存在，如果存在且不是目录，则报错

	// 检查idls文件是否存在，不存在则报错

}

// 1. layout manager 在指定目录创建layout
// 2. protoc-gen-gin 创建gin的go代码
func (h *initHandler) Run() {
	// 1. 创建.vodka目录
	meta.CreateMeta(h.root, h.idls)

	// 2. 在out目录下创建layout
	h.layout.Gen()

	// 3. 根据模板创建文件
	h.code.Gen()
}

func HandleInit(cmd args.Cmd, arg args.Argument) {
	// 检查项目目录下的vodka目录，存在则说明是已有项目，应该使用update
	if meta.Get() != nil {
		help.ExistProject()
	}

	initArg, ok := arg.(*args.Init)
	if !ok {
		help.InternalError()
	}

	h := newInitHandler(initArg)
	h.Run()
}
