package falta

import (
	"fmt"
	"math"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)


type Correntes_de_falta struct {
	Fase_a				float64
	Sequencia_positiva	float64
	Sequencia_negativa	float64
	Sequencia_zero		float64
}


type Tensao_de_sequencia struct {
	Sequencia_positiva	float64
	Sequencia_negativa	float64
	Sequencia_zero		float64
}


type Corrente_de_sequencia struct {
	Sequencia_positiva	float64
	Sequencia_negativa	float64
	Sequencia_zero		float64
}


type Tensao_de_fase struct {
	A	complex128
	B	complex128
	C	complex128
}


var a = complex(-(1.0/2.0), math.Sqrt(3.0/2.0))


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
	for nome_linha, linha := range elementos_tipo_2_3 {
		posicao_de := barras_sistema[linha.De].Posicao
		posicao_para := barras_sistema[linha.Para].Posicao
		corrente := (tensoes_barras[posicao_de] - tensoes_barras[posicao_para]) / linha.Impedancia_positiva

		corrente_nos_ramos[nome_linha] = corrente

		fmt.Println("\tBarra", linha.De, "para", linha.Para, "=", geral.Round(corrente, 4), "pu")
    }
}


// Retorna:
// 		Corrente de falta na fase A
// 		Componente de sequencia positiva
// 		Componente de sequencia negativa
// 		Componente de sequencia zero
func Corrente_falta_monofasica(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, barra_curto_circuito string) (Correntes_de_falta){

	Vf := 1.0 //pu

	posicao_na_zbus := barras_sistema[barra_curto_circuito].Posicao

	// Corrente de falta na fase A
	If_sequencia := Vf / (zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_positiva[posicao_na_zbus][posicao_na_zbus] + zbus_zero[posicao_na_zbus][posicao_na_zbus])

	// Circuito aberto, não há corrente de falta monofasica
	if zbus_zero[posicao_na_zbus][posicao_na_zbus] == 0 {
		If_sequencia = 0
	}

	If_monofasica := Correntes_de_falta{
		Fase_a:				If_sequencia * 3,
		Sequencia_positiva:	If_sequencia,
		Sequencia_negativa:	If_sequencia,
		Sequencia_zero:		If_sequencia,
	}

	fmt.Println("A corrente de falta na fase A é de", If_monofasica.Fase_a, "pu")

	return If_monofasica
}


func Tensoes_de_sequencia(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, corrente Correntes_de_falta) (map[string]Tensao_de_sequencia) {

	var tensoes_sequencia = make(map[string]Tensao_de_sequencia)

	for barras, posicao := range barras_sistema {
		tensoes_sequencia[barras] = Tensao_de_sequencia{
			Sequencia_positiva:	1 - zbus_positiva[posicao.Posicao][posicao.Posicao] * corrente.Sequencia_positiva,
			Sequencia_negativa:	0 - zbus_positiva[posicao.Posicao][posicao.Posicao] * corrente.Sequencia_negativa,
			Sequencia_zero:		0 - zbus_zero[posicao.Posicao][posicao.Posicao] 	* corrente.Sequencia_zero,
		}
	}

	return tensoes_sequencia
}


func Tensoes_de_fase(barras_sistema map[string]zbus.Posicao_zbus, tensoes_sequencia map[string]Tensao_de_sequencia) (map[string]Tensao_de_fase) {

	var tensoes_fase = make(map[string]Tensao_de_fase)

	// Matriz Transformação
	//	1, 	1 	  1
	//	1, 	a*a   a 
	//	1, 	a 	  a*a
	for barras := range barras_sistema {
		tensoes_fase[barras] = Tensao_de_fase{
			A:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
			B:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + a*a * complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + a * complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
			C:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + a * complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + a*a * complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
		}
	}

	return tensoes_fase
}


func Correntes_nas_linhas(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, tensoes_sequencia map[string]Tensao_de_sequencia, elementos_tipo_2_3 map[string]barra.Dados_de_linha, barras_sistema map[string]zbus.Posicao_zbus) (map[string]Corrente_de_sequencia) {

	var corrente_na_linha = make(map[string]Corrente_de_sequencia)

	for nome_linha, linha := range elementos_tipo_2_3 {
		posicao_de := barras_sistema[linha.De].Posicao
		posicao_para := barras_sistema[linha.Para].Posicao

		corrente_na_linha[nome_linha] = Corrente_de_sequencia{
			Sequencia_positiva: (tensoes_sequencia[linha.De].Sequencia_positiva - tensoes_sequencia[linha.Para].Sequencia_positiva) / zbus_positiva[posicao_de][posicao_para],
			Sequencia_negativa: (tensoes_sequencia[linha.De].Sequencia_negativa - tensoes_sequencia[linha.Para].Sequencia_negativa) / zbus_positiva[posicao_de][posicao_para],
			Sequencia_zero: 	(tensoes_sequencia[linha.De].Sequencia_zero - tensoes_sequencia[linha.Para].Sequencia_zero) / zbus_zero[posicao_de][posicao_para],
		}
    }

	return corrente_na_linha
}