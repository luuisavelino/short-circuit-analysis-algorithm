package falta

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)




func Falta_trifasica(zbus zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string, elementos_tipo_2_3 []barra.Dados_de_linha) {

	var tensoes_barras []float64
	var corrente_nos_ramos []float64
	var tamanho_do_sistema int = len(zbus)


	posicao := barras_sistema[barra_curto_circuito].Posicao
	corrente_curto_circuito := 1 / zbus[posicao][posicao]

	fmt.Printf("\nA corrente de curto circuito na barra %v é %v pu\n", barra_curto_circuito, geral.Round(corrente_curto_circuito, 4))

	for x := 0; x < tamanho_do_sistema; x++ {
		tensao := 1 - (zbus[x][posicao] * corrente_curto_circuito)
		tensoes_barras = append(tensoes_barras, tensao)
    }

	fmt.Println("As tensões nas barras são:")
	for barra, posicao := range barras_sistema {
		fmt.Println("\tBarra " + barra + " =", geral.Round(tensoes_barras[posicao.Posicao], 4), "pu")
    } 

	// Encontrando a corrente em cada ramo:
	for x := 0; x < len(elementos_tipo_2_3); x++ {
		posicao_de := barras_sistema[elementos_tipo_2_3[x].De].Posicao
		posicao_para := barras_sistema[elementos_tipo_2_3[x].Para].Posicao
		corrente := (tensoes_barras[posicao_de] - tensoes_barras[posicao_para]) / elementos_tipo_2_3[x].Impedancia_positiva

		corrente_nos_ramos = append(corrente_nos_ramos, corrente)
    }


	fmt.Println("As correntes nos ramos são:")
	for posicao, linha := range elementos_tipo_2_3 {
		fmt.Println("\tBarra", linha.De, "para", linha.Para, "=", geral.Round(corrente_nos_ramos[posicao], 4), "pu")
    } 

}