package protobuf

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestProtoGen(t *testing.T) {

	s := flag.String("out", "./test.pb.go", "out")
	flag.Parse()

	var req pluginpb.CodeGeneratorRequest
	req.FileToGenerate = []string{"./test.proto"}

	//设置参数,生成plugin
	opts := protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stderr, "a=%s\n", *s)

	// protoc 将一组文件结构传递给程序处理,包含proto中import中的文件
	for _, file := range plugin.Files {
		fmt.Println("file:", file)

		if !file.Generate { //显示传入的文件为true
			continue
		}

		genF := plugin.NewGeneratedFile(fmt.Sprintf("%s_error.pb.go", file.GeneratedFilenamePrefix), file.GoImportPath) //用来处理生成文件的对象
		generateFile(genF, file, *s)
	}

	// 生成response
	resp := plugin.Response()
	out, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}

	// 相应输出到stdout, 它将被 protoc 接收
	fmt.Printf("%v\n", string(out))
}

func generateFile(genF *protogen.GeneratedFile, file *protogen.File, out string) *protogen.GeneratedFile {

	return nil
}
