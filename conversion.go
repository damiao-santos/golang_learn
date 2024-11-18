package main

import "fmt"

type Conversion struct {
	Factor float64
	Unit   string
}

func (conv Conversion) Apply(value float64) float64 {
	return value * conv.Factor
}

func getConversion(option int) Conversion {
	conversions := map[int]Conversion{
		1: {1, "m"},
		2: {100, "cm"},
		3: {0.001, "km"},
	}

	if conv, ok := conversions[option]; ok {
		return conv
	} else {
		fmt.Println("Opção de conversão inválida")
		return Conversion{1, "m"}
	}
}

func printResult(value float64, conversion Conversion) {
	result := conversion.Apply(value)
	fmt.Printf("Resultado: %.2f %s\n", result, conversion.Unit)
	fmt.Println("Obrigado")
}
