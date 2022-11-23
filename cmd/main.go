package main

import (
    "github.com/luuisavelino/short-circuit-analysis-algorithm/internal/input"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"
    "github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/falta"

    "fmt"
    "time"
)

func main() {

    start := time.Now()

    elementos_tipo_1 := barra.Elementos_tipo_1(input.Tabela_excel())
    elementos_tipo_2_3 := barra.Elementos_tipo_2_3(input.Tabela_excel())
    
    zbus_positiva, _, barras_sistema := zbus.Zbus(elementos_tipo_1, elementos_tipo_2_3)
    
    falta.Falta_trifasica(zbus_positiva, barras_sistema, "3", barra.Elementos_tipo_2_3(input.Tabela_excel()))

    fmt.Println("\nO tempo de execução foi de", time.Since(start))
}