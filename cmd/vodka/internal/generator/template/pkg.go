package template

import "google.golang.org/protobuf/compiler/protogen"

const (
	contextPkg = protogen.GoImportPath("context")
	ginPkg     = protogen.GoImportPath("github.com/gin-gonic/gin")
)

const (
	PkgInternal = "internal"
)
