package generator

type Gin struct {
	root string
	idls []string
}

func NewGin(root string, idls []string) *Gin {
	g := &Gin{
		root: root,
		idls: idls,
	}
	return g
}

func (g *Gin) Gen() {}
