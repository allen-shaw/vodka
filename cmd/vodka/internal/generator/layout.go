package generator

type Layout struct {
}

func NewLayout() *Layout {

	l := &Layout{}
	return l
}

func (l *Layout) Gen() {}
