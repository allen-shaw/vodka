package protobuf

import "google.golang.org/protobuf/compiler/protogen"

var (
	vodkaModeOption = Options{
		DefualtHttpMethod: false,
	}

	defaultOption = Options{
		DefualtHttpMethod: true,
	}
)

type GenOptions struct {
	Imports []protogen.GoImportPath
	Output  string
}

type Options struct {
	DefualtHttpMethod bool
	Api               GenOptions
	Service           GenOptions
	Server            GenOptions
	Router            GenOptions
}
