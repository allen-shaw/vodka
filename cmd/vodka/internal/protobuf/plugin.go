package protobuf

import (
	"flag"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/protobuf/log"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
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
		// 1. 生成Server
		// 1.1. 当前有哪些source files
		// 1.2. 有哪些Service
		// 2. 

		for _, f := range gen.Files {
			log.Debug("file:%v\n", f.Desc.Path())
			if !f.Generate {
				continue
			}
			genfile(gen, f)
		}

		return nil
	})

	return 0
}

type ServerGenerator struct {
}

type RouterGenerator struct {
}

type ServiceGenerator struct {
}

type ModelGenerator struct {
}

type ApiGenerator struct {
}
