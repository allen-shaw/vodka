package generator

type Code struct {
	root string
	idls []string

	ginG *ProtocGenGin
}

func NewCode(root string, idls []string) *Code {
	c := &Code{
		root: root,
		idls: idls,
	}

	c.ginG = newProtocGenGin()

	return c
}

func (c *Code) Gen() {
	// 生成gomod文件
	c.genGomod()

	// 生成git.ignore
	c.genGitIgnore()

	// 生成script
	c.genScripts()

	// 生成deploy
	c.genDeploy()

	// 先拷贝proto到idl目录
	c.copyIDLs()

	// 调用proto-gen-gin生成gin代码
	c.ginG.Run()

	// 生成main.go
	c.genMain()
}

func (c *Code) copyIDLs() {

}

func (c *Code) genGomod() {

}

func (c *Code) genGitIgnore() {

}

// 生成script
func (c *Code) genScripts() {

}

// 生成deploy
func (c *Code) genDeploy() {

}

func (c *Code) genMain() {

}
