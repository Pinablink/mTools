package command

import "fmt"

type CClearCmdl struct {
}

func NewCClearCmdl() *CClearCmdl {
	return &CClearCmdl{}
}

//
func (ref *CClearCmdl) ExecuteCommand(param ...string) {
	fmt.Println("Executando comando de limpeza do Linux")
}
