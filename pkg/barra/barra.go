package barra

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/geral"
	"github.com/xuri/excelize/v2"
)


type Dados_de_linha struct {
    De	                string
    Para	            string
    Nome	            string
    Impedancia_positiva float64
    Impedancia_zero     float64
}


// Entrada:     Tabela do excel
//
// Processo:    Pega a tabela do 3 (Dados dos Transformadores) do Excel
//              Retira as duas primeiras linhas (informações)
//              Armazena todos os dados em uma variável do tipo Dados_de_linha (struct)
//
// Saida:       Retorna um map contendo os dados de todos transformadores
func transformadores(tabela_excel *excelize.File) map[string]Dados_de_linha {

    dados_transformadores, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[3])
    dados_transformadores = dados_transformadores[2:]
    geral.Valida_erro(err)

    elementos_transformadores := make(map[string]Dados_de_linha)
    for x := 0; x < len(dados_transformadores); x ++ {

        transformador := dados_transformadores[x][0] + "-" +dados_transformadores[x][1]
        impedancia_atual := elementos_transformadores[transformador].Impedancia_positiva

        elementos_transformadores[transformador] = Dados_de_linha{
            De:	                    dados_transformadores[x][0],
            Para:	                dados_transformadores[x][1],
            Nome:	                dados_transformadores[x][2],
            Impedancia_positiva:    geral.Impedancia(dados_transformadores[x][6], dados_transformadores[x][7], impedancia_atual),
            Impedancia_zero:        0.0,
        }
    }

    return elementos_transformadores
}


// Entrada:     Tabela do excel
//
// Processo:    Pega a tabela do 2 (Dados dos Geradores) do Excel
//              Retira a primeira linha (informações)
//              Armazena todos os dados em uma variável do tipo Dados_de_linha (struct)
//
// Saida:       Retorna um map contendo os dados de todos os elementos tipo 1
func Elementos_tipo_1(tabela_excel *excelize.File) []Dados_de_linha {
    var elementos_tipo_1 []Dados_de_linha

    dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[2])
    dados_linhas = dados_linhas[1:]
    geral.Valida_erro(err)

    for x := 0; x < len(dados_linhas); x ++ {
        elementos_tipo_1 = append(elementos_tipo_1, Dados_de_linha{
            De:	                    dados_linhas[x][0],
            Nome:	                dados_linhas[x][1],
            Impedancia_positiva:	geral.String_para_float(dados_linhas[x][3]) / 100,
            Impedancia_zero:        0,
        })
    }

    return elementos_tipo_1
}


// Entrada:     Tabela do excel
//
// Processo:    Pega a tabela do 1 (Dados de Linha) do Excel
//              Retira as duas primeiras linhas (informações)
//              Armazena todos os dados em uma variável do tipo Dados_de_linha (struct)
//
// Obs:         Pegamos os dados dos transformadores e passamos para cá, pois são elementos do tipo 2
//
// Saida:       Retorna um map contendo os dados de todos os elementos tipo 1
func Elementos_tipo_2_3(tabela_excel *excelize.File) []Dados_de_linha {
    var elementos_tipo_2_3 []Dados_de_linha

    dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
    dados_linhas = dados_linhas[2:]
    geral.Valida_erro(err)

    for _, dado_do_transformador := range transformadores(tabela_excel) {
        elementos_tipo_2_3 = append(elementos_tipo_2_3, dado_do_transformador)
    }

    for x := 0; x < len(dados_linhas); x ++ {
        elementos_tipo_2_3 = append(elementos_tipo_2_3, Dados_de_linha{
            De:	                    dados_linhas[x][0],
            Para:	                dados_linhas[x][1],
            Nome:	                dados_linhas[x][2],
            Impedancia_positiva:    geral.Impedancia(dados_linhas[x][5], dados_linhas[x][6], 0),
            Impedancia_zero:        geral.Impedancia(dados_linhas[x][10], dados_linhas[x][11], 0),
        })
    }

    return elementos_tipo_2_3
}


