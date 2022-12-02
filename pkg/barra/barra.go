package barra

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/input"
	"github.com/xuri/excelize/v2"
)


type Dados_de_linha struct {
    De	                string
    Para	            string
    Nome	            string
    Impedancia_positiva complex128
    Impedancia_zero     complex128
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
            Impedancia_zero:        0,
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
            Impedancia_positiva:	complex(0, geral.String_para_float(dados_linhas[x][3]) / 100),
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
func Elementos_tipo_2_3(tabela_excel *excelize.File, curto_circuito input.Ponto_curto_circuito) map[string]Dados_de_linha {
    var elementos_tipo_2_3 = make(map[string]Dados_de_linha)
    var elemento_barra Dados_de_linha

    dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
    dados_linhas = dados_linhas[2:]
    geral.Valida_erro(err)

    for _, dado_do_transformador := range transformadores(tabela_excel) {
        elementos_tipo_2_3[dado_do_transformador.De+"-"+dado_do_transformador.Para] = dado_do_transformador
    }

    for x := 0; x < len(dados_linhas); x ++ {
        _, elemento_ja_existe := elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]]

        if elemento_ja_existe {
            elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]] = Dados_de_linha{
                De:	                    dados_linhas[x][0],
                Para:	                dados_linhas[x][1],
                Nome:	                dados_linhas[x][2],
                Impedancia_positiva:    geral.Impedancia(dados_linhas[x][5], dados_linhas[x][6], elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]].Impedancia_positiva),
                Impedancia_zero:        geral.Impedancia(dados_linhas[x][10], dados_linhas[x][11], elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]].Impedancia_positiva),
            }
        } else {
            elementos_tipo_2_3[dados_linhas[x][0]+"-"+dados_linhas[x][1]] = Dados_de_linha{
                De:	                    dados_linhas[x][0],
                Para:	                dados_linhas[x][1],
                Nome:	                dados_linhas[x][2],
                Impedancia_positiva:    geral.Impedancia(dados_linhas[x][5], dados_linhas[x][6], 0),
                Impedancia_zero:        geral.Impedancia(dados_linhas[x][10], dados_linhas[x][11], 0),
            }
        }
    }


    if curto_circuito.Ponto != 0 && curto_circuito.Ponto != 100 {
        barra_de_para, de_para := elementos_tipo_2_3[curto_circuito.De+"-"+curto_circuito.Para]
        barra_para_de, para_de := elementos_tipo_2_3[curto_circuito.Para+"-"+curto_circuito.De]

        if de_para {
            delete(elementos_tipo_2_3, curto_circuito.De+"-"+curto_circuito.Para)
            elemento_barra = barra_de_para
        } else if para_de {
            delete(elementos_tipo_2_3, curto_circuito.Para+"-"+curto_circuito.De)
            elemento_barra = barra_para_de
    
        }

        // Adiciona o elemento tipo 2 (Barra curto-circuitada até De)
        elementos_tipo_2_3["barra_curto_circuito-"+curto_circuito.De] = Dados_de_linha{
            De:	                    curto_circuito.De,
            Para:	                "barra_curto_circuito",
            Nome:	                "Elemento tipo 2 do curto-circuito",
            Impedancia_positiva:    elemento_barra.Impedancia_positiva * complex(float64(curto_circuito.Ponto), 0) / 100,
            Impedancia_zero:        elemento_barra.Impedancia_zero * complex(float64(curto_circuito.Ponto), 0) / 100,
        }

        // Adciona o elemento tipo 3 (Linha entre a barra cc até Para)
        elementos_tipo_2_3["barra_curto_circuito-"+curto_circuito.Para] = Dados_de_linha{
            De:	                    "barra_curto_circuito",
            Para:	                curto_circuito.Para,
            Nome:	                "Elemento tipo 3 do curto-circuito",
            Impedancia_positiva:    elemento_barra.Impedancia_positiva * complex(float64(100 - curto_circuito.Ponto), 0) / 100,
            Impedancia_zero:        elemento_barra.Impedancia_zero * complex(float64(100 - curto_circuito.Ponto), 0) / 100,
        }
    }

    return elementos_tipo_2_3
}


