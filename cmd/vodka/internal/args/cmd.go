package args

import (
	"strings"

	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
)

type Cmd string

const (
	CmdInit    Cmd = "init"
	CmdUpdate  Cmd = "update"
	CmdVersion Cmd = "version"
	CmdUpgrade Cmd = "upgrade"
	CmdHelp    Cmd = "help"
)

var shorthand = map[string]Cmd{
	"h": CmdHelp,
	"v": CmdVersion,
}

func Parse(args []string) (cmd Cmd, arg Argument) {
	checkArgs(args)

	cmd = parseCmd(args[1])

	switch cmd {
	case CmdInit:
		arg = parseInitArgs(args)
	case CmdUpdate:
		arg = parseUpdateArgs(args)
	case CmdVersion:
		arg = parseVersionArgs(args)
	case CmdUpgrade:
		arg = parseUpgradeArgs(args)
	case CmdHelp:
		arg = ParseHelpArgs(args)
	default:
		help.InvalidArgs()
	}

	return cmd, arg
}

func checkArgs(args []string) {
	if len(args) < 2 {
		help.InvalidArgs()
	}
}

func parseCmd(s string) Cmd {
	raw := trim(s)
	if cmd, ok := shorthand[raw]; ok {
		return cmd
	}
	return Cmd(raw)
}

func trim(s string) string {
	t := strings.TrimPrefix(s, "-")
	return strings.ToLower(t)
}
