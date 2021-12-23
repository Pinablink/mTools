package oslabel

import (
	"errors"
	"runtime"
)

type OSLabel int
type OSCommand byte

const (
	LINUX   OSLabel = 1
	WINDOWS OSLabel = 2
)

const CLEAR_CMD OSCommand = 0x0A

//
func QOS() (OSLabel, error) {
	strCurr := runtime.GOOS

	if strCurr == "windows" {
		return WINDOWS, nil
	} else if strCurr == "linux" {
		return LINUX, nil
	} else {
		return -1, errors.New("Plataforma em execução não implementa o comando solicitado")
	}
}
