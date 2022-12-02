package falta

import (
	"fmt"
	"math"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)

func Corrente_falta_monofasica(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Componente_de_fase, Componente_de_sequencia) {

	Vf := 1.0 + 0i //pu

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao

	Icc_a := (3 * Vf) / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus])

	// Circuito aberto, não há corrente de falta monofasica
	if zbus_zero[posicao_na_zbus][posicao_na_zbus] == 0 {
		Icc_a = 0
	}

	Icc_fase := Componente_de_fase{
		A:	Icc_a,
		B:	0,
		C:	0,
	}

	Icc_sequencia := Componente_de_sequencia{
		Sequencia_positiva:	Icc_a / 3,
		Sequencia_negativa:	Icc_a / 3,
		Sequencia_zero:		Icc_a / 3,
	}

	fmt.Println("As correntes de falta monofasica são:")
	fmt.Println(Icc_fase)
	fmt.Println(Icc_sequencia)

	return Icc_fase, Icc_sequencia
}


func Corrente_falta_bifasica(zbus_positiva zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Componente_de_fase, Componente_de_sequencia) {
	var Icc_f complex128
	var Vf complex128 = 1.0 //pu

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao
	Icc_a_positivo := Vf / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus])

	Icc_fase := Componente_de_fase{
		A:	0,
		B:	complex(-math.Sqrt(3), 0) * Icc_f,
		C:	complex(math.Sqrt(3), 0) * Icc_f,
	}

	Icc_sequencia := Componente_de_sequencia{
		Sequencia_positiva:	Icc_a_positivo,
		Sequencia_negativa:	-Icc_a_positivo,
		Sequencia_zero:		0,
	}

	return Icc_fase, Icc_sequencia
}

func Falta_trifasica(zbus zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string, elementos_tipo_2_3 map[string]barra.Dados_de_linha) {

	var tensoes_barras []complex128
	var corrente_nos_ramos = make(map[string]complex128)
	var tamanho_do_sistema int = len(zbus)


	posicao := barras_sistema[barra_curto_circuito].Posicao
	corrente_curto_circuito := 1 / zbus[posicao][posicao]

	fmt.Printf("\nA corrente de curto circuito na barra %v é %v pu\n", barra_curto_circuito, corrente_curto_circuito)

	for x := 0; x < tamanho_do_sistema; x++ {
		tensao := 1 - (zbus[x][posicao] * corrente_curto_circuito)
		tensoes_barras = append(tensoes_barras, tensao)
    }

	fmt.Println("As tensões nas barras são:")
	for barra, posicao := range barras_sistema {
		fmt.Println("\tBarra " + barra + " =", tensoes_barras[posicao.Posicao], "pu")
    } 

	// Encontrando a corrente em cada ramo:
	fmt.Println("As correntes nos ramos são:")
	for nome_linha, linha := range elementos_tipo_2_3 {
		posicao_de := barras_sistema[linha.De].Posicao
		posicao_para := barras_sistema[linha.Para].Posicao
		corrente := (tensoes_barras[posicao_de] - tensoes_barras[posicao_para]) / linha.Impedancia_positiva

		corrente_nos_ramos[nome_linha] = corrente

		fmt.Println("\tBarra", linha.De, "para", linha.Para, "=", corrente, "pu")
    }
}