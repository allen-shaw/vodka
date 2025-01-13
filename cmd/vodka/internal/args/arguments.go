package args

import (
	"flag"
	"strings"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/meta"
)

type Argument interface {
}

type Init struct {
	IDLs []string
	Out  string
}

type Update struct {
	IDLs []string
}

type Version struct {
	Gin   bool
	Vodka bool
}

type Upgrade struct {
}

type Help struct {
	All bool
	Cmd Cmd
}

type flagSet struct {
	*flag.FlagSet
	comment string
}

func parseInitArgs(args []string) *Init {
	initCmd := &flagSet{
		FlagSet: flag.NewFlagSet("init", flag.ExitOnError),
		comment: help.InitComment,
	}

	var (
		idl string
		out string
	)

	initCmd.StringVar(&idl, "idl", "", help.InitIDLUsage)
	initCmd.StringVar(&out, "out", ".", help.InitOutUsage)

	initCmd.Parse(args[2:])

	idls := strings.Split(idl, ",")
	if len(idls) == 0 {
		help.InvalidInitArgs()
	}

	initArgs := &Init{
		IDLs: idls,
		Out:  out,
	}

	return initArgs
}

func parseUpdateArgs(args []string) *Update {
	updateCmd := &flagSet{
		FlagSet: flag.NewFlagSet("update", flag.ExitOnError),
		comment: help.UpdateComment,
	}

	var (
		idl string
	)

	updateCmd.StringVar(&idl, "idl", "", help.UpdateIDLUsage)
	updateCmd.Parse(args[2:])

	idls := strings.Split(idl, ",")
	md := meta.MustGet()
	if len(idls) == 0 {
		idls = md.IDLs
	}

	updateArgs := &Update{
		IDLs: idls,
	}

	return updateArgs
}

func parseVersionArgs(args []string) *Version {
	if len(args) <= 2 {
		// 用户输入 vodka version
		return &Version{true, true}
	}

	v := trim(args[2])
	switch v {
	case meta.Vodka:
		return &Version{Vodka: true}
	case meta.Gin:
		return &Version{Gin: true}
	default:
		return &Version{true, true}
	}
}

func parseUpgradeArgs(args []string) *Upgrade {
	return &Upgrade{}
}

func ParseHelpArgs(args []string) *Help {
	if len(args) <= 2 {
		return &Help{All: true}
	}

	cmd := parseCmd(args[2])
	return &Help{Cmd: cmd}
}
