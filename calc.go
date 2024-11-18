package main

import "fmt"

func main() {
	var query, conversionOption int
	var radius float64
	var side float64
	var shape string

	fmt.Print("Por favor escolha qual forma deseja calcular: ")
	fmt.Scanf("%s", &shape)

	switch shape {
	case "square", "quadrado":

		fmt.Print("Insira o um dos lados (em metros): ")
		fmt.Scanf("%f", &side)
		fmt.Printf("Escolha:\n 1 - Área \n 2 - Perímetro \n Opção Desejada: ")
		fmt.Scanf("%d", &query)

		fmt.Printf("Escolha a unidade de conversão:\n 1 - Metros (m) \n 2 - Centímetros (cm) \n 3 - Quilômetros (km) \n Opção Desejada: ")
		fmt.Scanf("%d", &conversionOption)

		square := square{side: side}
		conversion := getConversion(conversionOption)
		switch query {
		case 1:
			printResult(square.Area(), conversion)
		case 2:
			printResult(square.Perimeter(), conversion)
		default:
			fmt.Println("Opção inválida")
		}
	case "circulo", "circle":

		fmt.Print("Insira o Raio do Circulo (em metros): ")
		fmt.Scanf("%f", &radius)
		fmt.Printf("Escolha:\n 1 - Área \n 2 - Perímetro \n 3 - Diâmetro \n Opção Desejada: ")
		fmt.Scanf("%d", &query)

		fmt.Printf("Escolha a unidade de conversão:\n 1 - Metros (m) \n 2 - Centímetros (cm) \n 3 - Quilômetros (km) \n Opção Desejada: ")
		fmt.Scanf("%d", &conversionOption)

		circle := Circle{Radius: radius}
		conversion := getConversion(conversionOption)
		switch query {
		case 1:
			printResult(circle.Area(), conversion)
		case 2:
			printResult(circle.Perimeter(), conversion)
		case 3:
			printResult(circle.Diameter(), conversion)
		default:
			fmt.Println("Opção inválida")
		}
	}

}
