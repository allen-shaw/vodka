package handler

import (
	"github.com/allen-shaw/vodka/cmd/vodka/internal/args"
	"github.com/allen-shaw/vodka/cmd/vodka/internal/help"
)

func Handle(cmd args.Cmd, arg args.Argument) {
	switch cmd {
	case args.CmdInit:
		HandleInit(cmd, arg)
	case args.CmdUpdate:
		HandleUpdate(cmd, arg)
	case args.CmdVersion:
		HandleUpdate(cmd, arg)
	case args.CmdUpgrade:
		HandleUpgrade(cmd, arg)
	case args.CmdHelp:
		HandleHelp(cmd, arg)
	default:
		help.InvalidArgs()
	}
}
