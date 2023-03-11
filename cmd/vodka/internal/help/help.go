package help

const (
	InitComment    = ""
	UpdateComment  = ""
	VersionComment = ""
	UpgradeComment = ""
	HelpComment    = ""
)

const (
	InitIDLUsage = ""
	InitOutUsage = ""

	UpdateIDLUsage = ""
)

// 打印args提示，并中断进程
func InvalidArgs() {

}

func InvalidInitArgs() {

}

func NotProjectDir() {
	// 没有vodka目录或者没有vodka.meta,当前目录不是一个vodka项目
}

func ExistProject() {
	// 项目已经存在，请使用update更新，不能使用init
}

func InternalError() {

}
