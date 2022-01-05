package mexec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	CMD        = "cmd"
	POWERSHELL = "powershell"
	BASH       = "/bin/bash"
)

func ExecCommand(env string, cmdName string, args []string) ([]byte, error) {
	cmd := exec.Command(cmdName, args...)
	if env != "" {
		cmd.Path = env
	}
	//var stdout bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	//if err := cmd.Run(); err != nil {
	//	return nil, []byte(err.Error())
	//}
	return cmd.CombinedOutput()
}

func ExecCommandContext(env string, cmdName string, args []string, ctx context.Context) ([]byte, error) {
	cmd := exec.CommandContext(ctx, cmdName, args...)
	if env != "" {
		cmd.Path = env
	}
	//var stdout bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &stdout
	//cmd.Stderr = &stderr
	//if err := cmd.Run(); err != nil {
	//	return nil, []byte(err.Error())
	//}
	return cmd.CombinedOutput()
}

type MExec struct {
	Cmd *exec.Cmd
}

type ExecFunc func(e *MExec)

func WinCmd() ExecFunc {
	return func(e *MExec) {
		e.Cmd = exec.Command("cmd")
	}
}

func Powershell() ExecFunc {
	return func(e *MExec) {
		e.Cmd = exec.Command("powershell")
	}

}

func CmdName(cmd string, args ...string) ExecFunc {
	return func(e *MExec) {
		e.Cmd = exec.Command(cmd, args...)
	}
}

func Bash() ExecFunc {
	return func(e *MExec) {
		e.Cmd = exec.Command("/bin/bash", "-c")
	}
}

func NewMExec(cmds ...ExecFunc) *MExec {
	e := &MExec{}
	for _, cmd := range cmds {
		cmd(e)
	}
	return e
}

func (e *MExec) SetEevPath(env string) {
	e.Cmd.Dir = env
}

func (e *MExec) Start(cmd string) {
	stdin := bytes.NewBuffer(nil)
	e.Cmd.Stdin = stdin
	stdin.WriteString(cmd)
	stdoutIn, _ := e.Cmd.StdoutPipe()
	stderrIn, _ := e.Cmd.StderrPipe()
	var stdoutBuf, stderrBuf bytes.Buffer
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := e.Cmd.Start()
	if err != nil {
		return
	}

	go func() {
		_, _ = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, _ = io.Copy(stderr, stderrIn)
	}()
	err = e.Cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
}
