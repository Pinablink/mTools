package mtime

import "time"

// DateCurrentDefault Retorna a Data Corrente do Sistema no formato YYYYMMDD
func DateCurrentDefault() string {
	const formatDate = "20060102"
	var time = time.Now()
	return time.Format(formatDate)
}
