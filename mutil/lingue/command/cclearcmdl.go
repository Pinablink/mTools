package command

import (
	"os"
	"os/exec"
)

type CClearCmdl struct {
}

func NewCClearCmdl() *CClearCmdl {
	return &CClearCmdl{}
}

// ExecuteCommand Implementa a forma como o linux interpreta a solicitação de comando, para limpeza do cmd
func (ref *CClearCmdl) ExecuteCommand(param ...string) {
	refcmd := exec.Command("clear")
	refcmd.Stdout = os.Stdout
	refcmd.Run()
}
