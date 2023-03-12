package layout

type DirType int

const (
	DirTypeRoot DirType = iota + 1
	DirTypeMeta
	DirTypeIDL
	DirTypeInternal
	DirTypeAPI
	DirTypeDeploy
	DirTypeModel
	DirTypeScript
)

type Node struct {
	Name     string  `json:"name"`
	Type     DirType `json:"type"`
	IsDir    bool    `json:"is_dir"`
	Parent   *Node   `json:"parent"`
	Children []*Node `json:"children"`
}

type Layout struct {
	Root *Node `json:"root"`
}

func FromTemplate(tpl string) *Layout {
	l := &Layout{}
	return l
}
