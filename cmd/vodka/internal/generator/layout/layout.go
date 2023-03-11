package layout

type DirType int

const (
	DirTypeRoot DirType = iota + 1
	DirTypeMeta
	DirTypeIDL
	DirTypeAPI
	DirTypeModel
	DirTypeScript
)

type Node struct {
	Parent   *Node
	Children []*Node
	IsDir    bool
	Name     string
	Type     DirType
}

type Layout struct {
	Root *Node
}

func FromTemplate(tpl string) *Layout {
	l := &Layout{}
	return l
}
