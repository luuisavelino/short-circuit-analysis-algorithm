package zbus

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

type posicao_zbus struct {
    Posicao   int
}

func Zbus(elementos_tipo_1 []barra.Dados_de_linha, elementos_tipo_2_3 []barra.Dados_de_linha) {

    var zbus_positiva [50][50]float64
    var zbus_zero [50][50]float64
    var elementos_tipo_3 []barra.Dados_de_linha

    barras_adicionadas := make(map[string]posicao_zbus)
    posicao := 0

    // Adiciona os elementos do tipo 1
    // Loop passando por todos os elementos do tipo 1, e adicionando cada um na matriz Zbus
    for _, dados_linha := range elementos_tipo_1 {

        zbus_positiva = Adiciona_elemento_tipo_1_na_zbus(zbus_positiva, posicao, dados_linha.Impedancia_positiva)
        zbus_zero = Adiciona_elemento_tipo_1_na_zbus(zbus_zero, posicao, dados_linha.Impedancia_zero)

        barras_adicionadas[dados_linha.De] = posicao_zbus{
            Posicao:    posicao,
        }

        posicao++
    }


    // Adiciona os elementos do tipo 2
    // Valida se o elemento é do tipo 2, caso seja, adiciona na Zbus
    // Caso o elemento seja do tipo 3, ele adiciona em uma lista que será utilizada futuramente para adicionar os elementos tipo 3
    for len(elementos_tipo_2_3) != 0 {
        for x := 0; x < len(elementos_tipo_2_3); x++ {

            linha := elementos_tipo_2_3[x]

            _, existe_de := barras_adicionadas[linha.De]
            _, existe_para := barras_adicionadas[linha.Para]
   
            if existe_de && existe_para {
                elementos_tipo_3 = append(elementos_tipo_3, linha)
                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)

            } else if existe_de {
                zbus_positiva = Adiciona_elemento_tipo_2_na_zbus(zbus_positiva, barras_adicionadas[linha.De].Posicao, posicao, linha.Impedancia_positiva)
                zbus_zero = Adiciona_elemento_tipo_2_na_zbus(zbus_zero, barras_adicionadas[linha.De].Posicao, posicao, linha.Impedancia_zero)

                barras_adicionadas[linha.Para] = posicao_zbus{
                    Posicao:    posicao,
                }

                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++

            } else if existe_para {
                zbus_positiva = Adiciona_elemento_tipo_2_na_zbus(zbus_positiva, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_positiva)
                zbus_zero = Adiciona_elemento_tipo_2_na_zbus(zbus_zero, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_zero)

                barras_adicionadas[linha.De] = posicao_zbus{
                    Posicao:    posicao,
                }

                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++
            }
        }
    }


    // Com a lista criada de elementos do tipo 3, adicionamos na Zbus
    for x := 0; x < len(elementos_tipo_3); x++ {
        linha := elementos_tipo_3[x]

        zbus_positiva = Adiciona_elemento_tipo_3_na_zbus(zbus_positiva, barras_adicionadas[linha.De].Posicao, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_positiva)
        zbus_zero = Adiciona_elemento_tipo_3_na_zbus(zbus_zero, barras_adicionadas[linha.De].Posicao, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_zero)

        posicao++
    }
}

func RemoveIndex(s []barra.Dados_de_linha, index int) []barra.Dados_de_linha {
	return append(s[:index], s[index+1:]...)
}