package layout

var layoutTpl = defaultLayoutTpl

func SetLayout(layout string) {
	layoutTpl = layout
}

func GetLayout() string {
	return layoutTpl
}

const defaultLayoutTpl = `
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
