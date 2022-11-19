package input

import (
	"github.com/xuri/excelize/v2"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/geral"
)

func Tabela_excel() *excelize.File {
    tabela_excel, err := excelize.OpenFile("../data/dados_acch.xlsx")
    geral.Valida_erro(err)

    return tabela_excel
}