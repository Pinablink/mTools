package lingue

import (
	"github.com/Pinablink/mTools/mutil/lingue/command"
	"github.com/Pinablink/mTools/mutil/lingue/ilingue"
	"github.com/Pinablink/mTools/mutil/lingue/oslabel"
)

// Disponibiliza recursos que uniformiza alguns comandos de Windows e Linux.
type Lingue struct {
	mapResource map[oslabel.OSLabel]map[oslabel.OSCommand]ilingue.ICCommand
}

// NewPoliglota Disponibiliza uma instância da estrutura Lingue
func NewLingue() *Lingue {
	mapRef := make(map[oslabel.OSLabel]map[oslabel.OSCommand]ilingue.ICCommand)

	// Criando as implementações com os comandos
	var cclearLinux ilingue.ICCommand = command.NewCClearCmdl()
	var cclearWindows ilingue.ICCommand = command.NewCClearCmdw()

	mapCmmLinux := make(map[oslabel.OSCommand]ilingue.ICCommand)
	mapCmmWindows := make(map[oslabel.OSCommand]ilingue.ICCommand)

	mapCmmLinux[oslabel.CLEAR_CMD] = cclearLinux
	mapCmmWindows[oslabel.CLEAR_CMD] = cclearWindows
	//

	mapRef[oslabel.LINUX] = mapCmmLinux
	mapRef[oslabel.WINDOWS] = mapCmmWindows

	return &Lingue{mapResource: mapRef}
}

// ExecCommand Executa o comando solicitado de acordo com o Sistema Operacional provedor
func (lingue *Lingue) ExecCommand(refCommand oslabel.OSCommand) error {

	key, err := oslabel.QOS()

	if err != nil {
		return err
	}

	(lingue.mapResource[key])[refCommand].ExecuteCommand()

	return nil
}
