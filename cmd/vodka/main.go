package main

import (
	"os"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/args"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/handler"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/meta"
)

func main() {
	// 1. 初始化
	meta.Init()

	// 2. 解析命令行参数
	cmd, args := args.Parse(os.Args)

	// 3. 根据不同命令调用不同handler
	handler.Handle(cmd, args)
}
