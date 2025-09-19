package models

import "strconv"

func ConverteParaFloat(valor string) float64 {
	valorConvertido, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		panic(err.Error())
	}
	return valorConvertido
}

func ConverteParaInt(valor string) int {
	valorConvertido, err := strconv.Atoi(valor)
	if err != nil {
		panic(err.Error())
	}
	return valorConvertido
}
