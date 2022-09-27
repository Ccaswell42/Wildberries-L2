package builtins

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
)

type CmdContext struct {
	Cmd    string
	Args   []string
	StdIn  io.Reader
	StdOut io.Writer
	Next   *CmdContext
}

func MyPWD(ctx *CmdContext) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(ctx.StdOut, "%s\n", dir)
	if err != nil {
		return err
	}
	return nil
}

func MyCD(ctx *CmdContext) error {
	err := os.Chdir(ctx.Args[0])
	if err != nil {
		return err
	}
	return nil
}

func MyEcho(ctx *CmdContext) error {
	_, err := fmt.Fprintln(ctx.StdOut, ctx.Args)
	if err != nil {
		return err
	}
	return nil
}

func MyKill(ctx *CmdContext) error {
	pid, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return err
	}
	err = syscall.Kill(pid, syscall.SIGINT)
	if err != nil {
		return err
	}
	return nil
}
