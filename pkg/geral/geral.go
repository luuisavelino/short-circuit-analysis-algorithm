package geral

import (
	"log"
	"math"
	"strconv"
	//"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

func Impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual float64) float64 {
    resistencia, _ := strconv.ParseFloat(resistencia_linha, 64)
    reatancia, _ := strconv.ParseFloat(reatancia_linha, 64)

    impedancia := math.Sqrt(math.Pow(resistencia, 2) + math.Pow(reatancia, 2))
    
    if impedancia_atual != 0 {
        impedancia = (impedancia * impedancia_atual) / (impedancia + impedancia_atual)
    }

    return impedancia
}


func String_para_float(grandeza_str string) float64 {
    grandeza, _ := strconv.ParseFloat(grandeza_str, 64)

    return grandeza
}


func Difference(a, b []string) []string {
    mb := make(map[string]struct{}, len(b))

    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []string
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}


func Valida_erro(err error) {
    if err != nil {
        log.Fatal(err)
        return
    }
}
