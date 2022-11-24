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

    curto_circuito := input.Ponto_curto_circuito{
        De:     "2",
        Para:   "3",
        Ponto:  50,
    }

    var barra_curto_circuito string
    var tamanho_do_sistema int = 6

    if curto_circuito.Ponto == 0 {
        barra_curto_circuito = curto_circuito.De
    } else if curto_circuito.Ponto == 100 {
        barra_curto_circuito = curto_circuito.Para
    } else {
        barra_curto_circuito = "barra_curto_circuito"
        tamanho_do_sistema++
    }


    

    elementos_tipo_1 := barra.Elementos_tipo_1(input.Tabela_excel())
    elementos_tipo_2_3 := barra.Elementos_tipo_2_3(input.Tabela_excel(), curto_circuito)
    
    zbus_positiva, _, barras_sistema := zbus.Zbus(elementos_tipo_1, elementos_tipo_2_3, tamanho_do_sistema)
    
    


    falta.Falta_trifasica(zbus_positiva, barras_sistema, barra_curto_circuito, barra.Elementos_tipo_2_3(input.Tabela_excel(), curto_circuito))

    fmt.Println("\nO tempo de execução foi de", time.Since(start))
}