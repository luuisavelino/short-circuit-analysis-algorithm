package main

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/internal/input"

    "fmt"
    "time"
)

func main() {

    start := time.Now()


    //barra.Barras_do_sistema(input.Tabela_excel())

    tipo_1 := barra.Elementos_tipo_1(input.Tabela_excel())
    tipo_2_3 := barra.Elementos_tipo_2_3(input.Tabela_excel())

    zbus.Zbus(tipo_1, tipo_2_3)


    fmt.Println("\nO tempo de execução foi de", time.Since(start))
}