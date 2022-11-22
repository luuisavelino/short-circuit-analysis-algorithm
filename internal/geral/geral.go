package geral

import (
	"log"
    "math"
	"strconv"
)

// Entrada:     Resistencia, Reatancia, e a impedância atual
// Processo:    Realiza o calculo da impedância
//              Caso já exista uma impedância atual, ele irá realizar o paralelo entre eles
// Saída:       Retorna o valor da impedância
func Impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual float64) float64 {
    resistencia, _ := strconv.ParseFloat(resistencia_linha, 64)
    reatancia, _ := strconv.ParseFloat(reatancia_linha, 64)

    impedancia := math.Sqrt(math.Pow(resistencia, 2) + math.Pow(reatancia, 2))
    
    if impedancia_atual != 0 {
        impedancia = (impedancia * impedancia_atual) / (impedancia + impedancia_atual)
    }

    return impedancia
}


// Irá converter o tipo string para o tipo float64
func String_para_float(grandeza_str string) float64 {
    grandeza, _ := strconv.ParseFloat(grandeza_str, 64)

    return grandeza
}


// Realiza a validação do erro
func Valida_erro(err error) {
    if err != nil {
        log.Fatal(err)
        return
    }
}


func Round(valor float64, casas float64) float64 {

	return math.Round( valor * (math.Pow(10, casas))) / (math.Pow(10, casas))
}