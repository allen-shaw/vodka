package meta

import (
	"os"
	"path/filepath"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/util"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/version"
	"github.com/pelletier/go-toml/v2"
)

const (
	metaFile = "vodka.meta"
	metaDir  = ".vodka"
)

var (
	gMeta *GlobalMeta
	pMeta *Meta
)

type writer struct {
	dir      string
	filename string
	file     *os.File
	data     interface{}
}

func (w *writer) Save() error {
	return toml.NewEncoder(w.file).Encode(w.data)
}

type GlobalMeta struct {
	*writer
	AutoUpgrade bool   `toml:"auto_upgrade"`
	CreatTime   string `toml:"create_time"`
	UpdateTime  string `toml:"update_time"`
	UpgradeTime string `toml:"upgrade_time"`
}

type Meta struct {
	*writer
	IDLs          []string `toml:"idls"`
	Out           string   `toml:"out"`
	CreateVersion string   `toml:"create_version"`
	CreatTime     string   `toml:"create_time"`
	UpdateTime    string   `toml:"update_time"`
}

func newMeta(idls []string, out string) *Meta {
	now := util.Now()
	m := &Meta{
		IDLs:          idls,
		Out:           out,
		CreateVersion: version.Version,
		CreatTime:     now,
		UpdateTime:    now,
	}
	w := &writer{
		dir:      filepath.Join(out, metaDir),
		filename: metaFile,
		data:     m,
	}
	var err error
	w.file, err = os.Create(filepath.Join(w.dir, w.filename))
	checkError(err, "create meta file")

	m.writer = w

	return m
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

func CreateMeta(root string, idls []string) bool {
	m := newMeta(idls, root)
	err := m.Save()
	checkError(err, "save meta file fail")
	return true
}

func UpdateMeta(meta *Meta) bool {

	return true
}

func checkError(err error, prompts string) {
	if err == nil {
		return
	}
}
