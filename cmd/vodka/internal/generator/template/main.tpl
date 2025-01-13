func main() {
    addr := ":8080"
	s := internal.NewHttpServer()
	if err := s.Run(addr); err != nil {
		panic(err)
	}
}