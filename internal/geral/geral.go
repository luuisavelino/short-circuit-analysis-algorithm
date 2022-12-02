package geral

import (
	"log"
    //"math"
	"strconv"
)

// Entrada:     Resistencia, Reatancia, e a impedância atual
// Processo:    Realiza o calculo da impedância
//              Caso já exista uma impedância atual, ele irá realizar o paralelo entre eles
// Saída:       Retorna o valor da impedância
func Impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual complex128) complex128 {
    var resistencia, _ = strconv.ParseFloat(resistencia_linha, 64)
    var reatancia, _ = strconv.ParseFloat(reatancia_linha, 64)

    impedancia := complex(resistencia, reatancia)

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
        log.Fatal(err.Error())
        return
    }
}


// Arredondamento para uma determinada quantidade de casas decimais
func Round(valor complex128, casas float64) {

    //var convesao complex128 = complex(math.Pow(10, casas), 0)

	//return math.Round( (valor * convesao) / convesao)
}