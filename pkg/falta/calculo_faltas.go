package falta

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)

type Componente_de_sequencia struct {
	Sequencia_positiva complex128
	Sequencia_negativa complex128
	Sequencia_zero     complex128
}

type Componente_de_fase struct {
	A complex128
	B complex128
	C complex128
}

func Tensoes_de_sequencia_nas_barras(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, barras_sistema map[string]zbus.Posicao_zbus, corrente Componente_de_sequencia) map[string]Componente_de_sequencia {

	var tensoes_sequencia = make(map[string]Componente_de_sequencia)

	for barras, posicao := range barras_sistema {
		tensoes_sequencia[barras] = Componente_de_sequencia{
			Sequencia_positiva: 1 - zbus_positiva[posicao.Posicao][posicao.Posicao]*corrente.Sequencia_positiva,
			Sequencia_negativa: 0 - zbus_positiva[posicao.Posicao][posicao.Posicao]*corrente.Sequencia_negativa,
			Sequencia_zero:     0 - zbus_zero[posicao.Posicao][posicao.Posicao]*corrente.Sequencia_zero,
		}
	}

	return tensoes_sequencia
}

func Correntes_de_sequencia_nas_linhas(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, tensoes_sequencia map[string]Componente_de_sequencia, elementos_tipo_2_3 map[string]barra.Dados_de_linha) map[string]Componente_de_sequencia {

	var corrente_de_sequencia_na_linha = make(map[string]Componente_de_sequencia)

	for nome_linha, linha := range elementos_tipo_2_3 {
		impedancia_na_linha_positiva := elementos_tipo_2_3[linha.De+"-"+linha.Para].Impedancia_positiva
		impedancia_na_linha_zero := elementos_tipo_2_3[linha.De+"-"+linha.Para].Impedancia_zero

		corrente_de_sequencia_na_linha[nome_linha] = Componente_de_sequencia{
			Sequencia_positiva: (tensoes_sequencia[linha.De].Sequencia_positiva - tensoes_sequencia[linha.Para].Sequencia_positiva) / impedancia_na_linha_positiva,
			Sequencia_negativa: (tensoes_sequencia[linha.De].Sequencia_negativa - tensoes_sequencia[linha.Para].Sequencia_negativa) / impedancia_na_linha_positiva,
			Sequencia_zero:     geral.Valida_divisao_por_0(tensoes_sequencia[linha.De].Sequencia_zero-tensoes_sequencia[linha.Para].Sequencia_zero, impedancia_na_linha_zero),
		}
	}

	return corrente_de_sequencia_na_linha
}
