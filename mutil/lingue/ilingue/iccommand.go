package ilingue

// ICCommand Disponiliza uma interface para implementação e uniformização dos comandos
type ICCommand interface {
	ExecuteCommand(param ...string)
}
