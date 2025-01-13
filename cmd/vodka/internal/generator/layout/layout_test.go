package layout

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/stvp/assert"
)

const tpl = `
.
├── deploy
├── idl
│   └── google
│       └── api
├── internal
│   ├── api
│   │   └── ${services}
│   └── model
│       └── api
│           └── ${services}
└── scripts
`

func TestFromTpl(t *testing.T) {
	dirs := strings.Split(tpl, "\n")
	for i, dir := range dirs {
		fmt.Println(i, dir)
	}
}

func TestToLayout(t *testing.T) {
	root := &Node{
		Parent:   nil,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     ".",
		Type:     DirTypeRoot,
	}
	deploy := &Node{
		// Parent:   root,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "deploy",
		Type:     DirTypeDeploy,
	}
	root.Children = append(root.Children, deploy)

	idl := &Node{
		// Parent:   root,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "idl",
		Type:     DirTypeIDL,
	}
	root.Children = append(root.Children, idl)

	idlgoogle := &Node{
		// Parent:   idl,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "idl/google",
		Type:     DirTypeIDL,
	}
	idl.Children = append(idl.Children, idlgoogle)

	idlgoogleapi := &Node{
		// Parent:   idlgoogle,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "idl/google/api",
		Type:     DirTypeIDL,
	}
	idlgoogle.Children = append(idlgoogle.Children, idlgoogleapi)

	internal := &Node{
		// Parent:   root,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "internal",
		Type:     DirTypeInternal,
	}
	root.Children = append(root.Children, internal)

	internalapi := &Node{
		// Parent:   internal,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "internal/api",
		Type:     DirTypeAPI,
	}
	internal.Children = append(internal.Children, internalapi)

	internalmodel := &Node{
		// Parent:   internal,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "internal/model",
		Type:     DirTypeModel,
	}
	internal.Children = append(internal.Children, internalmodel)

	internalmodelapi := &Node{
		// Parent:   internalmodel,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "internal/model/api",
		Type:     DirTypeModel,
	}
	internalmodel.Children = append(internalmodel.Children, internalmodelapi)

	script := &Node{
		// Parent:   root,
		Children: make([]*Node, 0),
		IsDir:    true,
		Name:     "script",
		Type:     DirTypeScript,
	}
	root.Children = append(root.Children, script)

	layout := &Layout{Root: root}
	j, err := json.Marshal(layout)
	assert.Nil(t, err)

	fmt.Println(string(j))
}

func TestFromTemplate(t *testing.T) {
	tpl := `{"root":{"name":".","type":1,"is_dir":true,"parent":null,"children":[{"name":"deploy","type":6,"is_dir":true,"parent":null,"children":[]},{"name":"idl","type":3,"is_dir":true,"parent":null,"children":[{"name":"idl/google","type":3,"is_dir":true,"parent":null,"children":[{"name":"idl/google/api","type":3,"is_dir":true,"parent":null,"children":[]}]}]},{"name":"internal","type":4,"is_dir":true,"parent":null,"children":[{"name":"internal/api","type":5,"is_dir":true,"parent":null,"children":[]},{"name":"internal/model","type":7,"is_dir":true,"parent":null,"children":[{"name":"internal/model/api","type":7,"is_dir":true,"parent":null,"children":[]}]}]},{"name":"script","type":8,"is_dir":true,"parent":null,"children":[]}]}}`

	layout := &Layout{}
	err := json.Unmarshal([]byte(tpl), layout)
	assert.Nil(t, err)

	queue := make(chan *Node, 1000)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		n, ok := <-queue
		if !ok {
			return
		}
		fmt.Println(n)
	}()

	handler(layout.Root, queue)
	close(queue)
	wg.Wait()
}

func handler(n *Node, queue chan *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Name)
	queue <- n
	for _, c := range n.Children {
		handler(c, queue)
	}
}
