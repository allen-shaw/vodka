package main

import "github.com/allen-shaw/vodka/demo/internal"

func main() {
	addr := ":8080"
	s := internal.NewHttpServer()
	if err := s.Run(addr); err != nil {
		panic(err)
	}
}
