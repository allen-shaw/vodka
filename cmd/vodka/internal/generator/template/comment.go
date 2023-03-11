package template

import (
	"fmt"
	"strings"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/version"
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
