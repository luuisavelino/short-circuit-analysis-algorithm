package barra

import (
    "github.com/xuri/excelize/v2"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/geral"
)

type Dados_de_barra struct {
    n_bar   int
    Barra   string
    Cidade  string
    Tipo    string
    Tensao  string
    Max     string
    Min     string
}

type Dados_de_linha struct {
    De	                string
    Para	            string
    Nome	            string
    Impedancia_positiva float64
    Impedancia_zero     float64
}


func barras_de_geracao(tabela_excel *excelize.File) []string {
    var barra_geradores []string

    colunas_dados_de_geradores, err := tabela_excel.Cols(tabela_excel.GetSheetList()[2])
    geral.Valida_erro(err)

    for colunas_dados_de_geradores.Next() {
        col, err := colunas_dados_de_geradores.Rows()
        geral.Valida_erro(err)

        barra_geradores = col[1:len(col)]
        break
    }

    return barra_geradores
}


func Barras_do_sistema(tabela_excel *excelize.File) ([]string, map[string]Dados_de_barra) {
    var barras_do_sistema []string

    // Pega a coluna de barras da tabela de dados das barras
    colunas_dados_de_barra, err := tabela_excel.Cols(tabela_excel.GetSheetList()[0])
    geral.Valida_erro(err)

    // Coloca em uma lista todas as barras do sistema, excluindo as barras dos geradores
    for colunas_dados_de_barra.Next() {
        col, err := colunas_dados_de_barra.Rows()
        geral.Valida_erro(err)
        barras_do_sistema = geral.Difference(col[2:len(col)-1], barras_de_geracao(tabela_excel))
        
        break
    }

    // Pega as linhas dos dados de barra
    dados_de_barra_bruto, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[0])
    geral.Valida_erro(err)

    // Filtra os dados de barra, exclindo o cabeçalho
    dados_de_barra_bruto = dados_de_barra_bruto[2:len(dados_de_barra_bruto)]
    
    // Desenvolve um dicinário de dados de barra, onde teremos todas as barras do sistema
    barras := make(map[string]Dados_de_barra)
    for x := 0; x < len(dados_de_barra_bruto); x ++ {
        barras[dados_de_barra_bruto[x][0]] = Dados_de_barra{
            n_bar:  x+1,
            Barra:  dados_de_barra_bruto[x][0],
            Cidade: dados_de_barra_bruto[x][1], 
            Tipo:   dados_de_barra_bruto[x][2], 
            Tensao: dados_de_barra_bruto[x][3],
            Max:    dados_de_barra_bruto[x][4],
            Min:    dados_de_barra_bruto[x][5],
        }
    }

    return barras_do_sistema, barras
}


func Elementos_tipo_2_3(tabela_excel *excelize.File) []Dados_de_linha {

    dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
    dados_linhas = dados_linhas[2:len(dados_linhas)]
    geral.Valida_erro(err)

    var elementos_tipo_2_3 []Dados_de_linha
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


func Elementos_tipo_1(tabela_excel *excelize.File) map[string]Dados_de_linha {

    dados_transformadores, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[3])
    dados_transformadores = dados_transformadores[2:len(dados_transformadores)]
    geral.Valida_erro(err)

    elementos_tipo_1 := make(map[string]Dados_de_linha)
    for x := 0; x < len(dados_transformadores); x ++ {
        for _, y := range barras_de_geracao(tabela_excel) {
            if y == dados_transformadores[x][0] {
                elementos_tipo_1[dados_transformadores[x][1]] = Dados_de_linha{
                    De:	                    dados_transformadores[x][0],
                    Para:	                dados_transformadores[x][1],
                    Nome:	                dados_transformadores[x][2],
                    Impedancia_positiva:    geral.Impedancia(dados_transformadores[x][6], dados_transformadores[x][7], elementos_tipo_1[y].Impedancia_positiva),
                }
            }
            if y == dados_transformadores[x][1] {
                elementos_tipo_1[dados_transformadores[x][0]] = Dados_de_linha{
                    De:	                    dados_transformadores[x][0],
                    Para:	                dados_transformadores[x][1],
                    Nome:	                dados_transformadores[x][2],
                    Impedancia_positiva:    geral.Impedancia(dados_transformadores[x][6], dados_transformadores[x][7], elementos_tipo_1[y].Impedancia_positiva),
                }
            }
        }
    }

    return elementos_tipo_1
}
