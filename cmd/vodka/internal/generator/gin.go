package generator

type ProtocGenGin struct {
}

func newProtocGenGin() *ProtocGenGin {
	p := &ProtocGenGin{}

	// 查找protoc-gen-gin,如果没有则安装
	p.CheckOrInstall()

	return p
}

func (p *ProtocGenGin) Run() {
	// 调用protoc-gen-gin生成gin代码
	
}

func (p *ProtocGenGin) CheckOrInstall() {

}
