package random

import (
	"math/rand"
)

// GStrRandomCod Estrutura com recursos para gerar números aleatórios
type GStrRandomCod struct {
	Seed     int64
	RangeCod int
}

// New Disponibiliza uma nova instância da estrutura GStrRandomCod
func NewGStrRandomCod(mSeed int64, rangeCode int) *GStrRandomCod {

	return &GStrRandomCod{
		Seed:     mSeed,
		RangeCod: rangeCode,
	}

}

// Init Inicializa recursos internos necessários ao processamento
func (gr GStrRandomCod) Init() {
	rand.Seed(gr.Seed)
}

// GenerateStrCode Gera uma string contendo representações numérias e de letras
func (gr *GStrRandomCod) GenerateStrCode() string {
	var strRef string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var codStream []byte = make([]byte, gr.RangeCod)

	for index := range codStream {
		codStream[index] = strRef[rand.Intn(len(strRef))]
	}

	return string(codStream)
}

// Gera uma string contendo apenas representações numéricas
func (gr *GStrRandomCod) GenerateStrNumberCode() string {
	var strRef string = "0123456789"
	var codStream []byte = make([]byte, gr.RangeCod)

	for index := range codStream {
		codStream[index] = strRef[rand.Intn(len(strRef))]
	}

	return string(codStream)
}
