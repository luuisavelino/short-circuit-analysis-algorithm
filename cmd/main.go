package main

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barras"
)

func main() {
    barras.barras_do_sistema(barras.tabela_excel())

    tipo_1 := barras.Elementos_tipo_1(barras.tabela_excel())
    tipo_2_3 := barras.Elementos_tipo_2_3(barras.tabela_excel())

    barras.Zbus(tipo_1, tipo_2_3)

}