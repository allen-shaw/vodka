package generator

type Layout struct {
	root string
}

func NewLayout(root string) *Layout {
	l := &Layout{
		root: root,
	}
	return l
}

func (l *Layout) Gen() {}
