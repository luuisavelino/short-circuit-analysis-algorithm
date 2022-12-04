package falta

import (
	"fmt"
	"math"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)


const (
	Vf complex128 = 1.0 //pu
)


func Corrente_falta_monofasica(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Componente_de_fase, Componente_de_sequencia) {

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao

	Icc_a := (3 * Vf) / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus])

	// Circuito aberto, não há corrente de falta monofasica
	if imag(Icc_a) < math.Pow(10, -6) {
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

	return Icc_fase, Icc_sequencia
}


func Corrente_falta_bifasica(zbus_positiva zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Componente_de_fase, Componente_de_sequencia) {

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao
	Icc_a_positivo := Vf / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus])

	Icc_sequencia := Componente_de_sequencia{
		Sequencia_positiva:	Icc_a_positivo,
		Sequencia_negativa:	-Icc_a_positivo,
		Sequencia_zero:		0,
	}

	Icc_fase := Sequencia_para_fase(Icc_sequencia)

	return Icc_fase, Icc_sequencia
}


func Corrente_falta_bifasico_terra(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Componente_de_fase, Componente_de_sequencia) {

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao

	// Zeq = Z+  +  (Z- * Z0)/(Z- + Z0)
	impedancia_equivalente := zbus_positiva[posicao_na_zbus][posicao_na_zbus] + ((zbus_zero[posicao_na_zbus][posicao_na_zbus] * zbus_positiva[posicao_na_zbus][posicao_na_zbus]) / (zbus_zero[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus]))
	Icc := Vf / impedancia_equivalente

	Icc_sequencia := Componente_de_sequencia{
		Sequencia_positiva:	Icc,
		Sequencia_negativa:	Icc * (zbus_zero[posicao_na_zbus][posicao_na_zbus]) / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus]),			// Divisor de corrente
		Sequencia_zero:		Icc * (zbus_positiva[posicao_na_zbus][posicao_na_zbus]) / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus]),		// Divisor de corrente	
	}

	Icc_fase := Sequencia_para_fase(Icc_sequencia)

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
		tensao := Vf - (zbus[x][posicao] * corrente_curto_circuito)
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