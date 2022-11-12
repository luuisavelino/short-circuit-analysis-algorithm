package main

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

func main() {
    barra.Barras_do_sistema(barra.Tabela_excel())

    tipo_1 := barra.Elementos_tipo_1(barra.Tabela_excel())
    tipo_2_3 := barra.Elementos_tipo_2_3(barra.Tabela_excel())

    barra.Zbus(tipo_1, tipo_2_3)

}