package mexec

import (
	"testing"
)

func TestStart(t *testing.T) {
	Start()
}

func TestExecCommandRun(t *testing.T) {

	ExecCommandRun("cmd", []string{"ls"}, "", nil)

}
