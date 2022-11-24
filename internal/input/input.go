package input

import (
	"github.com/xuri/excelize/v2"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)


type Ponto_curto_circuito struct {
    De                      string
    Para                    string
    Ponto                   int
    Impedancia_positiva     float64
    Impedancia_zero         float64
}


func Tabela_excel() *excelize.File {
    tabela_excel, err := excelize.OpenFile("../data/exemplo_de_aula_3.xlsx")
    geral.Valida_erro(err)

    return tabela_excel
}