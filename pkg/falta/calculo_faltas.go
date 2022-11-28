package falta

import (
	"math"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)


type Componente_de_sequencia struct {
	Sequencia_positiva	float64
	Sequencia_negativa	float64
	Sequencia_zero		float64
}


type Componente_de_fase struct {
	A	complex128
	B	complex128
	C	complex128
}


var a = complex(-(1.0/2.0), math.Sqrt(3.0/2.0))


func Tensoes_de_sequencia(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, corrente Componente_de_sequencia) (map[string]Componente_de_sequencia) {

	var tensoes_sequencia = make(map[string]Componente_de_sequencia)

	for barras, posicao := range barras_sistema {
		tensoes_sequencia[barras] = Componente_de_sequencia{
			Sequencia_positiva:	1 - zbus_positiva[posicao.Posicao][posicao.Posicao] * corrente.Sequencia_positiva,
			Sequencia_negativa:	0 - zbus_positiva[posicao.Posicao][posicao.Posicao] * corrente.Sequencia_negativa,
			Sequencia_zero:		0 - zbus_zero[posicao.Posicao][posicao.Posicao] 	* corrente.Sequencia_zero,
		}
	}

	return tensoes_sequencia
}


func Tensoes_de_fase(barras_sistema map[string]zbus.Posicao_zbus, tensoes_sequencia map[string]Componente_de_sequencia) (map[string]Componente_de_fase) {

	var tensoes_fase = make(map[string]Componente_de_fase)

	for barras := range barras_sistema {
		tensoes_fase[barras] = Componente_de_fase{
			A:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
			B:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + a*a * complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + a * complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
			C:	complex(tensoes_sequencia[barras].Sequencia_zero, 0) + a * complex(tensoes_sequencia[barras].Sequencia_positiva, 0) + a*a * complex(tensoes_sequencia[barras].Sequencia_negativa, 0),
		}
	}

	return tensoes_fase
}


func Correntes_de_sequencia_nas_linhas(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, tensoes_sequencia map[string]Componente_de_sequencia, elementos_tipo_2_3 map[string]barra.Dados_de_linha, barras_sistema map[string]zbus.Posicao_zbus) (map[string]Componente_de_sequencia) {

	var corrente_de_sequencia_na_linha = make(map[string]Componente_de_sequencia)

	for nome_linha, linha := range elementos_tipo_2_3 {
		posicao_de := barras_sistema[linha.De].Posicao
		posicao_para := barras_sistema[linha.Para].Posicao

		corrente_de_sequencia_na_linha[nome_linha] = Componente_de_sequencia{
			Sequencia_positiva: (tensoes_sequencia[linha.De].Sequencia_positiva - tensoes_sequencia[linha.Para].Sequencia_positiva) / zbus_positiva[posicao_de][posicao_para],
			Sequencia_negativa: (tensoes_sequencia[linha.De].Sequencia_negativa - tensoes_sequencia[linha.Para].Sequencia_negativa) / zbus_positiva[posicao_de][posicao_para],
			Sequencia_zero: 	(tensoes_sequencia[linha.De].Sequencia_zero - tensoes_sequencia[linha.Para].Sequencia_zero) / zbus_zero[posicao_de][posicao_para],
		}
    }

	return corrente_de_sequencia_na_linha
}


func Corrente_na_linha(corrente_de_sequencia_na_linha map[string]Componente_de_sequencia) (map[string]Componente_de_fase) {

	var correntes = make(map[string]Componente_de_fase)

	for nome_linha, linha := range corrente_de_sequencia_na_linha {
		correntes[nome_linha] = Componente_de_fase{
			A:	complex(linha.Sequencia_zero, 0) + complex(linha.Sequencia_positiva, 0) + complex(linha.Sequencia_negativa, 0),
			B:	complex(linha.Sequencia_zero, 0) + a*a * complex(linha.Sequencia_positiva, 0) + a * complex(linha.Sequencia_negativa, 0),
			C:	complex(linha.Sequencia_zero, 0) + a * complex(linha.Sequencia_positiva, 0) + a*a * complex(linha.Sequencia_negativa, 0),
		}
	}

	return correntes
}