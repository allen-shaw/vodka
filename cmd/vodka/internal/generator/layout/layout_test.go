package layout

import (
	"fmt"
	"strings"
	"testing"
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
