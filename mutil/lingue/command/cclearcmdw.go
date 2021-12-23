package command

import (
	"os"
	"os/exec"
)

// CClearCmd Estrutura que uniformiza o comando de limpeza do CMD nos Sistemas Operacionais, Windows e Linux
type CClearCmdw struct {
}

//
func NewCClearCmdw() *CClearCmdw {
	return &CClearCmdw{}
}

// ExecuteCommand Implementa a forma como o windows interpreta a solicitação de comando, para limpeza do cmd
func (ref *CClearCmdw) ExecuteCommand(param ...string) {
	refcmd := exec.Command("cmd", "/c", "cls")
	refcmd.Stdout = os.Stdout
	refcmd.Run()
}
