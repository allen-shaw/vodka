package protobuf

import (
	"flag"
	"fmt"

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
		for _, f := range gen.Files {
			fmt.Println("file:", f)
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
