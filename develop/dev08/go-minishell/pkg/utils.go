package pkg

import (
	"minishell/pkg/builtins"
)

func FindBuiltins(cmd builtins.CmdContext) bool {
	if cmd.Cmd == "pwd" || cmd.Cmd == "cd" || cmd.Cmd == "kill" || cmd.Cmd == "echo" {
		return true
	}
	return false
}

func ExecuteBuiltins(cmd *builtins.CmdContext) error {
	if cmd.Cmd == "pwd" {
		return builtins.MyPWD(cmd)
	} else if cmd.Cmd == "cd" {
		return builtins.MyCD(cmd)
	} else if cmd.Cmd == "kill" {
		return builtins.MyKill(cmd)
	} else if cmd.Cmd == "echo" {
		return builtins.MyEcho(cmd)
	}
	return nil
}
