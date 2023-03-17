package main

import (
	"flag"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/protobuf"
)

func main() {
	vodka := flag.Bool("vodka", false, "run by vodka")
	flag.Parse()

	p := protobuf.NewPlugin(vodka)
	p.Run()
}
