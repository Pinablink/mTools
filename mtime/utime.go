package mtime

import "time"

// DateCurrentDefault Retorna a Data Corrente do Sistema no formato YYYYMMDD
func DateCurrentDefault() string {
	const formatDate = "2006-01-02"
	var time = time.Now()
	return time.Format(formatDate)
}

// DateTimeCurrentDefault Retorna a Data e Hora Corrente do Sistema no formato RFC3339
func DateTimeCurrentDefault() string {
	const formatDate = time.RFC3339
	var time = time.Now()
	return time.Format(formatDate)
}
