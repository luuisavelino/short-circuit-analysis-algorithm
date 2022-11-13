package zbus

import (
	"fmt"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

type posicao_zbus struct {
    Posicao   int
}

func Zbus(elementos_tipo_1 map[string]barra.Dados_de_linha, elementos_tipo_2_3 []barra.Dados_de_linha) {

    var zbus_positiva [30][30]float64
    var zbus_zero [30][30]float64

    barras_adicionadas := make(map[string]posicao_zbus)

    posicao := 0

    for barra, dados_linha := range elementos_tipo_1 {

        zbus_positiva = Adiciona_elemento_tipo_1(zbus_positiva, posicao, dados_linha.Impedancia_positiva)
        zbus_zero = Adiciona_elemento_tipo_1(zbus_zero, posicao, dados_linha.Impedancia_zero)

        barras_adicionadas[barra] = posicao_zbus{
            Posicao:    posicao,
        }

        posicao++

    }

    fmt.Println(zbus_positiva)

    for len(elementos_tipo_2_3) != 0 {
        for x := 0; x < len(elementos_tipo_2_3); x++ {

            _, existe_de := barras_adicionadas[elementos_tipo_2_3[x].De]
            _, existe_para := barras_adicionadas[elementos_tipo_2_3[x].Para]
   
            if existe_de == true && existe_para == true {
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
    
            } else if existe_de == true && existe_para == false {
                zbus_positiva = Adiciona_elemento_tipo_2(elementos_tipo_2_3[x].De, zbus_positiva, elementos_tipo_2_3[x].Impedancia_positiva, posicao, barras_adicionadas)
                zbus_zero = Adiciona_elemento_tipo_2(elementos_tipo_2_3[x].De, zbus_zero, elementos_tipo_2_3[x].Impedancia_zero, posicao, barras_adicionadas)

                barras_adicionadas[elementos_tipo_2_3[x].Para] = posicao_zbus{
                    Posicao:    posicao,
                }
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++
                
            } else if existe_de == false && existe_para == true {
                zbus_positiva = Adiciona_elemento_tipo_2(elementos_tipo_2_3[x].Para, zbus_positiva, elementos_tipo_2_3[x].Impedancia_positiva, posicao, barras_adicionadas)
                zbus_zero = Adiciona_elemento_tipo_2(elementos_tipo_2_3[x].Para, zbus_zero, elementos_tipo_2_3[x].Impedancia_zero, posicao, barras_adicionadas)

                barras_adicionadas[elementos_tipo_2_3[x].De] = posicao_zbus{
                    Posicao:    posicao,
                }
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++
            }

            
        }
    }

    fmt.Println(zbus_positiva)
    fmt.Println(barras_adicionadas)

    return

}

func RemoveIndex(s []barra.Dados_de_linha, index int) []barra.Dados_de_linha {
	return append(s[:index], s[index+1:]...)
}