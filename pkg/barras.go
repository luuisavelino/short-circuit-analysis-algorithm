package barras

import (
    "fmt"
    "log"
    "strconv"
    "math"
    "github.com/xuri/excelize/v2"
)

type dados_de_barra struct {
    n_bar   int
    Barra   string
    Cidade  string
    Tipo    string
    Tensao  string
    Max     string
    Min     string
}

type dados_de_linha struct {
    De	                string
    Para	            string
    Nome	            string
    Impedancia_positiva float64
    Impedancia_zero     float64
}

type posicao_zbus struct {
    Posicao   int
}


func Zbus(elementos_tipo_1 map[string]dados_de_linha, elementos_tipo_2_3 [26]dados_de_linha) {
    var zbus_positiva [10][10]float64
    var zbus_zero [10][10]float64

    fmt.Println(elementos_tipo_1, elementos_tipo_2_3)

    barras_adicionadas := make(map[string]posicao_zbus)

    posicao := 0

    for barra, dados_linha := range elementos_tipo_1 {
        fmt.Println(barra, dados_linha.Impedancia_positiva)
        fmt.Println(posicao)
        zbus_positiva[posicao][posicao] = dados_linha.Impedancia_positiva
        zbus_zero[posicao][posicao] = dados_linha.Impedancia_zero

        barras_adicionadas[barra] = posicao_zbus{
            Posicao:    posicao,
        }
        posicao++
    }

    //while(elementos_tipo_2_3) {

    //}
    
    for x := 0; x < len(elementos_tipo_2_3); x++ {

        barra_de, existe_de := barras_adicionadas[elementos_tipo_2_3[x].De]
        barra_para, existe_para := barras_adicionadas[elementos_tipo_2_3[x].Para]

        fmt.Println(barra_de, existe_de, barra_para, existe_para)

        fmt.Println(elementos_tipo_2_3)

        if existe_de == true && existe_para == true {
            
            fmt.Println("Elemento do tipo 3")
            fmt.Println("Removendo a linha ", x, "dos elementos 2 e 3")
            fmt.Println(elementos_tipo_2_3[x])


        } else if existe_de == true && existe_para == false {

            fmt.Println("Elemento do tipo 2")
            fmt.Println("Removendo a linha ", x, "dos elementos 2 e 3")
            fmt.Println(elementos_tipo_2_3[x])

        } else if existe_de == false && existe_para == true {

            fmt.Println("Elemento do tipo 2")
            fmt.Println("Removendo a linha ", x, "dos elementos 2 e 3")
            fmt.Println(elementos_tipo_2_3[x])

        } else {

            fmt.Println("Elemento não pertence ao sistema")

        }

        


    }
    
    fmt.Println(zbus_positiva)



    return

}



func adiciona_elemento_2(barra string, barra_conectada string, zbus [50][50]float64, impedancia float64, posicao int, barras_adicionadas map[string]posicao_zbus) [50][50]float64 {

    posicao_barra_conectada := barras_adicionadas[barra_conectada].Posicao

    for x := 0; x < posicao; x++ {
        zbus[x][posicao] = zbus[x][posicao_barra_conectada]
        zbus[posicao][x] = zbus[posicao_barra_conectada][x]
    }

    zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

    return zbus
}




func tabela_excel() *excelize.File {
    tabela_excel, err := excelize.OpenFile("../../ACCH/dados_acch.xlsx")
    valida_erro(err)

    return tabela_excel
}


func barras_de_geracao(tabela_excel *excelize.File) []string {
    var barra_geradores []string

    // Pega a coluna de barras da tabela de geradores
    colunas_dados_de_geradores, err := tabela_excel.Cols(tabela_excel.GetSheetList()[2])
    valida_erro(err)

    // Coloca em uma lista todas as barras dos geradores
    // As barras que se conectarem a ela, serão do tipo 1
    for colunas_dados_de_geradores.Next() {
        col, err := colunas_dados_de_geradores.Rows()
        if err != nil {
            log.Fatal(err)
        }

        barra_geradores = col[1:len(col)]
        break
    }

    return barra_geradores
}


func barras_do_sistema(tabela_excel *excelize.File) ([]string, map[string]dados_de_barra) {
    var barras_do_sistema []string

    // Pega a coluna de barras da tabela de dados das barras
    colunas_dados_de_barra, err := tabela_excel.Cols(tabela_excel.GetSheetList()[0])
    valida_erro(err)

    // Coloca em uma lista todas as barras do sistema, excluindo as barras dos geradores
    for colunas_dados_de_barra.Next() {
        col, err := colunas_dados_de_barra.Rows()
        valida_erro(err)
        barras_do_sistema = difference(col[2:len(col)-1], barras_de_geracao(tabela_excel))
        
        break
    }

    // Pega as linhas dos dados de barra
    dados_de_barra_bruto, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[0])
    valida_erro(err)

    // Filtra os dados de barra, exclindo o cabeçalho
    dados_de_barra_bruto = dados_de_barra_bruto[2:len(dados_de_barra_bruto)]
    
    // Desenvolve um dicinário de dados de barra, onde teremos todas as barras do sistema
    barras := make(map[string]dados_de_barra)
    for x := 0; x < len(dados_de_barra_bruto); x ++ {
        barras[dados_de_barra_bruto[x][0]] = dados_de_barra{
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


func Elementos_tipo_2_3(tabela_excel *excelize.File) [26]dados_de_linha {

    dados_linhas, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[1])
    dados_linhas = dados_linhas[2:len(dados_linhas)]
    valida_erro(err)


    var elementos_tipo_1_2 [26]dados_de_linha
    for x := 0; x < len(dados_linhas); x ++ {
        elementos_tipo_1_2[x] = dados_de_linha{
            De:	                    dados_linhas[x][0],
            Para:	                dados_linhas[x][1],
            Nome:	                dados_linhas[x][2],
            Impedancia_positiva:    impedancia(dados_linhas[x][5], dados_linhas[x][6], 0),
            Impedancia_zero:        impedancia(dados_linhas[x][10], dados_linhas[x][11], 0),
        }
    }

    return elementos_tipo_1_2
}


func Elementos_tipo_1(tabela_excel *excelize.File) map[string]dados_de_linha {

    dados_transformadores, err := tabela_excel.GetRows(tabela_excel.GetSheetList()[3])
    dados_transformadores = dados_transformadores[2:len(dados_transformadores)]
    valida_erro(err)

    elementos_tipo_1 := make(map[string]dados_de_linha)
    for x := 0; x < len(dados_transformadores); x ++ {
        for _, y := range barras_de_geracao(tabela_excel) {
            if y == dados_transformadores[x][0] {
                elementos_tipo_1[dados_transformadores[x][1]] = dados_de_linha{
                    De:	                    dados_transformadores[x][0],
                    Para:	                dados_transformadores[x][1],
                    Nome:	                dados_transformadores[x][2],
                    Impedancia_positiva:    impedancia(dados_transformadores[x][6], dados_transformadores[x][7], elementos_tipo_1[y].Impedancia_positiva),
                }
            }
            if y == dados_transformadores[x][1] {
                elementos_tipo_1[dados_transformadores[x][0]] = dados_de_linha{
                    De:	                    dados_transformadores[x][0],
                    Para:	                dados_transformadores[x][1],
                    Nome:	                dados_transformadores[x][2],
                    Impedancia_positiva:    impedancia(dados_transformadores[x][6], dados_transformadores[x][7], elementos_tipo_1[y].Impedancia_positiva),
                }
            }
        }
    }

    return elementos_tipo_1
}


func impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual float64) float64 {
    resistencia, _ := strconv.ParseFloat(resistencia_linha, 64)
    reatancia, _ := strconv.ParseFloat(reatancia_linha, 64)

    impedancia := math.Sqrt(math.Pow(resistencia, 2) + math.Pow(reatancia, 2))
    
    if impedancia_atual != 0 {
        impedancia = (impedancia * impedancia_atual) / (impedancia + impedancia_atual)
    }

    return impedancia
}


func difference(a, b []string) []string {
    mb := make(map[string]struct{}, len(b))

    for _, x := range b {
        mb[x] = struct{}{}
    }
    var diff []string
    for _, x := range a {
        if _, found := mb[x]; !found {
            diff = append(diff, x)
        }
    }
    return diff
}


func valida_erro(err error) {
    if err != nil {
        log.Fatal(err)
        return
    }

    return
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}