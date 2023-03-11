package meta

import (
	"time"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
)

var (
	gMeta *GlobalMeta
	pMeta *Meta
)

type writer struct {
	dir      string
	filename string
}

func (w *writer) Save() error {

	return nil
}

type GlobalMeta struct {
	writer
	AutoUpgrade bool      `toml:"auto_upgrade"`
	CreatTime   time.Time `toml:"create_time"`
	UpdateTime  time.Time `toml:"update_time"`
	UpgradeTime time.Time `toml:"upgrade_time"`
}

type Meta struct {
	writer
	IDLs          []string  `toml:"idls"`
	Out           string    `toml:"out"`
	CreateVersion string    `toml:"create_version"`
	CreatTime     time.Time `toml:"create_time"`
	UpdateTime    time.Time `toml:"update_time"`
}

func Init() {
	initGlobal()
	initMeta()

	cleanLog()
}

func GetGlobal() *GlobalMeta {
	return gMeta
}

func Get() *Meta {
	return pMeta
}

func MustGet() *Meta {
	if pMeta == nil {
		help.NotProjectDir()
	}
	return pMeta
}

// 读取home/.vodka/vodka,如果不存在则创建
func initGlobal() {

}

// 读取当前./.vodla/vodka.meta,不会自动创建，不存在则返回nil
func initMeta() {

}

// 删除 .vodka目录下的日志
func cleanLog() {

}

func CreateMeta(dir string, meta *Meta) bool {

	return true
}

func UpdateMeta(meta *Meta) bool {

	return true
}
