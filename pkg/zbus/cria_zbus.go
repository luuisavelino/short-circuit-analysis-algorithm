package zbus

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)



type posicao_zbus struct {
    Posicao   int
}

const Tamanho_do_sistema = 6

type matrix [Tamanho_do_sistema][Tamanho_do_sistema]float64



func Zbus(elementos_tipo_1 []barra.Dados_de_linha, elementos_tipo_2_3 []barra.Dados_de_linha) {

    var zbus_positiva matrix
    var zbus_zero matrix
    var elementos_tipo_3 []barra.Dados_de_linha

    barras_adicionadas := make(map[string]posicao_zbus)
    posicao := 0

    // Adiciona os elementos do tipo 1
    // Loop passando por todos os elementos do tipo 1, e adicionando cada um na matriz Zbus
    for _, dados_linha := range elementos_tipo_1 {

        zbus_positiva = Adiciona_elemento_tipo_1_na_zbus(zbus_positiva, posicao, dados_linha.Impedancia_positiva)
        zbus_zero = Adiciona_elemento_tipo_1_na_zbus(zbus_zero, posicao, dados_linha.Impedancia_zero)


        fmt.Println("Adicionado elemento tipo 1 -> Barra: " + dados_linha.De + "\t\tImpedancia:", dados_linha.Impedancia_positiva)

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

                fmt.Println("Adicionado elemento tipo 2 -> Linha: " + linha.De + "-" + linha.Para + "\tImpedancia:", linha.Impedancia_positiva)

                barras_adicionadas[linha.Para] = posicao_zbus{
                    Posicao:    posicao,
                }

                elementos_tipo_2_3 = RemoveIndex(elementos_tipo_2_3, x)
                posicao++

            } else if existe_para {
                zbus_positiva = Adiciona_elemento_tipo_2_na_zbus(zbus_positiva, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_positiva)
                zbus_zero = Adiciona_elemento_tipo_2_na_zbus(zbus_zero, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_zero)

                fmt.Println("Adicionado elemento tipo 2 -> Linha: " + linha.De + "-" + linha.Para + "\tImpedancia:", linha.Impedancia_positiva)

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

        fmt.Println("Adicionado elemento tipo 3 -> Linha: " + linha.De + "-" + linha.Para + " \tImpedancia:", linha.Impedancia_positiva, " \tRealizando redução de Kron")

        zbus_positiva = Adiciona_elemento_tipo_3_com_reducao_de_kron(zbus_positiva, barras_adicionadas[linha.De].Posicao, barras_adicionadas[linha.Para].Posicao, linha.Impedancia_positiva)
        zbus_zero = Adiciona_elemento_tipo_3_com_reducao_de_kron(zbus_zero, barras_adicionadas[linha.De].Posicao, barras_adicionadas[linha.Para].Posicao, linha.Impedancia_zero)

    }

    fmt.Println("\nA matriz Zbus do sistema é: ")
    for x := 0; x < Tamanho_do_sistema; x++ {
        for y := 0; y < Tamanho_do_sistema; y++ {
            fmt.Printf("\t%v\t", geral.Round(zbus_positiva[x][y], 4))
        }
        fmt.Println("")
    }
}

func RemoveIndex(s []barra.Dados_de_linha, index int) []barra.Dados_de_linha {
	return append(s[:index], s[index+1:]...)
}