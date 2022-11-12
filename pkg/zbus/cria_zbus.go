package zbus

import (
	"fmt"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

type posicao_zbus struct {
    Posicao   int
}

func Zbus(elementos_tipo_1 map[string]barra.Dados_de_linha, elementos_tipo_2_3 []barra.Dados_de_linha) {
    var zbus_positiva [10][10]float64
    var zbus_zero [10][10]float64

    barras_adicionadas := make(map[string]posicao_zbus)

    posicao := 0

    for barra, dados_linha := range elementos_tipo_1 {
        zbus_positiva[posicao][posicao] = dados_linha.Impedancia_positiva
        zbus_zero[posicao][posicao] = dados_linha.Impedancia_zero

        barras_adicionadas[barra] = posicao_zbus{
            Posicao:    posicao,
        }
        posicao++
    }

    fmt.Println(barras_adicionadas)

    for len(elementos_tipo_2_3) != 0 {
        for x := 0; x < len(elementos_tipo_2_3); x++ {

            _, existe_de := barras_adicionadas[elementos_tipo_2_3[x].De]
            _, existe_para := barras_adicionadas[elementos_tipo_2_3[x].Para]
   
            if existe_de == true && existe_para == true {
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
    
            } else if existe_de == true && existe_para == false {
                barras_adicionadas[elementos_tipo_2_3[x].Para] = posicao_zbus{
                    Posicao:    posicao,
                }
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++
                
            } else if existe_de == false && existe_para == true {
                barras_adicionadas[elementos_tipo_2_3[x].De] = posicao_zbus{
                    Posicao:    posicao,
                }
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++
            }

            
        }
    }

    fmt.Println(barras_adicionadas)
    fmt.Println(len(barras_adicionadas))



    return

}

func RemoveIndex(s []barra.Dados_de_linha, index int) []barra.Dados_de_linha {
	return append(s[:index], s[index+1:]...)
}