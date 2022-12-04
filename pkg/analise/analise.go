package analise

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/falta"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
)

func Mostra_matriz_zbus(zbus zbus.Matrix, tamanho_do_sistema int) {
    for x := 0; x < tamanho_do_sistema; x++ {
        for y := 0; y < tamanho_do_sistema; y++ {
            fmt.Printf("    %v\t", geral.Round_cmplx(zbus[x][y], 4))
        }
        fmt.Println("")
    }
	fmt.Println("")
}


func Analise_curto_circuito(zbus_positiva zbus.Matrix, zbus_zero zbus.Matrix, elementos_tipo_2_3 map[string]barra.Dados_de_linha, barras_sistema map[string]zbus.Posicao_zbus, Icc_fase falta.Componente_de_fase, Icc_sequencia falta.Componente_de_sequencia, curto_circuito barra.Ponto_curto_circuito) {

	tensoes_sequencia := falta.Tensoes_de_sequencia_nas_barras(zbus_positiva, zbus_zero, barras_sistema, Icc_sequencia)
	correntes_sequencia := falta.Correntes_de_sequencia_nas_linhas(zbus_positiva, zbus_zero, tensoes_sequencia, elementos_tipo_2_3, barras_sistema)


	//========================= CORRENTE DE FALTA =========================
	fmt.Println("\nCorrente de fase da falta: \nA:", geral.Retangular_To_Polar(Icc_fase.A), "\nB:", geral.Retangular_To_Polar(Icc_fase.B), "\nC:", geral.Retangular_To_Polar(Icc_fase.C))
	fmt.Println("\nCorrente de sequencia da falta: \n+:", geral.Retangular_To_Polar(Icc_sequencia.Sequencia_positiva), "\n-:", geral.Retangular_To_Polar(Icc_sequencia.Sequencia_negativa), "\n0:", geral.Retangular_To_Polar(Icc_sequencia.Sequencia_zero))

	// ========================= TENSÃO NAS BARRAS =========================
	fmt.Println("\nAs tensões de sequencia nas barras são:")
	fmt.Println("\tBARRA\t |   \tSEQUENCIA +  \t|   \tSEQUENCIA -  \t|   \tSEQUENCIA 0  \t|")
	for barra, componentes := range tensoes_sequencia {
		fmt.Println("\t", barra, "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_positiva), "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_negativa), "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_zero))
	}

	fmt.Println("\nAs tensões de fase nas barras são:")
	fmt.Println("\tBARRA\t |   \tFASE A  \t|   \tFASE B  \t|   \tFASE C  \t|")
	for barra, componentes_sequencia := range tensoes_sequencia {
		tensao_fase := falta.Sequencia_para_fase(componentes_sequencia)
		fmt.Println("\t", barra, "\t  ", geral.Retangular_To_Polar(tensao_fase.A), "\t  ", geral.Retangular_To_Polar(tensao_fase.B), "\t  ", geral.Retangular_To_Polar(tensao_fase.C))
	}


	// ========================= CORRENTE NAS LINHAS =========================
	fmt.Println("\nAs correntes de sequencia nas linhas são:")
	fmt.Println("\tLINHA\t |   \tSEQUENCIA +  \t|   \tSEQUENCIA -  \t|   \tSEQUENCIA 0  \t|")
	for linha, componentes := range correntes_sequencia {
		fmt.Println("\t", linha, "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_positiva), "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_negativa), "\t  ", geral.Retangular_To_Polar(componentes.Sequencia_zero))
	}
	
	fmt.Println("\nAs correntes de fase nas linhas são:")
	fmt.Println("\tLINHA\t |   \tFASE A  \t|   \tFASE B  \t|   \tFASE C  \t|")
	for linha, componentes := range correntes_sequencia {
		tensao_fase := falta.Sequencia_para_fase(componentes)
		fmt.Println("\t", linha, "\t  ", 	geral.Retangular_To_Polar(tensao_fase.A), "\t  ", 	geral.Retangular_To_Polar(tensao_fase.B), "\t  ", 	geral.Retangular_To_Polar(tensao_fase.C))
	}
}
