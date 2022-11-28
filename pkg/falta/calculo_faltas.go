package falta

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)


func Falta_trifasica(zbus zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string, elementos_tipo_2_3 map[string]barra.Dados_de_linha) {

	var tensoes_barras []float64
	var corrente_nos_ramos = make(map[string]float64)
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
	fmt.Println("As correntes nos ramos são:")
	for nome_linha, barra := range elementos_tipo_2_3 {
		posicao_de := barras_sistema[barra.De].Posicao
		posicao_para := barras_sistema[barra.Para].Posicao
		corrente := (tensoes_barras[posicao_de] - tensoes_barras[posicao_para]) / barra.Impedancia_positiva

		corrente_nos_ramos[nome_linha] = corrente

		fmt.Println("\tBarra", barra.De, "para", barra.Para, "=", geral.Round(corrente, 4), "pu")
    }
}


func Corrente_falta_monofasica(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (float64, float64, float64, float64) {

	Vf := 1.0 //pu

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao

	// Corrente de falta na fase A
	If_sequencia := Vf / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus])

	// Circuito aberto, não há corrente de falta monofasica
	if zbus_zero[posicao_na_zbus][posicao_na_zbus] == 0 {
		If_sequencia = 0
	}

	fmt.Println("A corrente de falta na fase A é de", 3*If_sequencia, "pu")

	// Retornando a corrente de falta na fase A
	// Componente de sequencia positiva
	// Componente de sequencia negativa
	// Componente de sequencia zero
	return 3*If_sequencia, If_sequencia, If_sequencia, If_sequencia
}

