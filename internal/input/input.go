package input

import (
	"github.com/xuri/excelize/v2"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)


func Tabela_excel(path string) *excelize.File {
    tabela_excel, err := excelize.OpenFile(path)
    geral.Valida_erro(err)

    return tabela_excel
}
