package mexec

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"server/common/mencoding"
)

func ExecCommandRun(cmdName string, args []string, env string, ctx context.Context) {
	cmd := exec.Command(cmdName, args...)
	if ctx != nil {
		cmd = exec.CommandContext(ctx, cmdName, args...)
	}
	if env != "" {
		cmd.Path = env
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	decoderBytes, err := mencoding.Byte2Utf8(stdout.Bytes())
	//stdoutReader := bufio.NewReader(bytes.NewReader(stdout.Bytes()))
	//decoderBytes, err := mencoding.Decoder(stdoutReader)
	//decoderBytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(stdout.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(decoderBytes))
	fmt.Println(stderr.String())
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

func Env() ExecFunc {
	return func(e *MExec) {

	}
}

func NewMExec(cmds ...ExecFunc) *MExec {
	e := &MExec{}
	for _, cmd := range cmds {
		cmd(e)
	}
	return e
}

func Start() {
	mexec := NewMExec(WinCmd())
	mexec.Cmd.Dir = ""
	stdout, err := mexec.Cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	stdin := bytes.NewBuffer(nil)
	mexec.Cmd.Stdin = stdin
	stdin.WriteString("ls")
	if err := mexec.Cmd.Start(); err != nil {
		fmt.Println(err)
	}
	if by, err := ioutil.ReadAll(stdout); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(by))
	}

}
