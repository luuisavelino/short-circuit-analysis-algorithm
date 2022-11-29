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
        Ponto:  0,
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
    
    zbus_positiva, zbus_zero, barras_sistema := zbus.Zbus(elementos_tipo_1, elementos_tipo_2_3, tamanho_do_sistema)

    fmt.Println("=====================================================")
    falta.Falta_trifasica(zbus_positiva, barras_sistema, barra_curto_circuito, barra.Elementos_tipo_2_3(input.Tabela_excel(), curto_circuito))
    fmt.Println("=====================================================")
    _, Icc_bifasica_sequencia := falta.Corrente_falta_bifasica(zbus_positiva, barras_sistema, barra_curto_circuito)
    fmt.Println("\nAs tensões de sequencia são:")
    tensoes_sequencia := falta.Tensoes_de_sequencia(zbus_positiva, zbus_zero, barras_sistema, Icc_bifasica_sequencia)
    fmt.Println(tensoes_sequencia)
    fmt.Println("\nAs tensões de fase são:")
    tensoes_fase := falta.Tensoes_de_fase(barras_sistema, tensoes_sequencia)
    fmt.Println(tensoes_fase)
    fmt.Println("\nAs correntes de sequencia nas linhas são:")
    corrente_de_sequencia_na_linha := falta.Correntes_de_sequencia_nas_linhas(zbus_positiva, zbus_zero, tensoes_sequencia, barra.Elementos_tipo_2_3(input.Tabela_excel(), curto_circuito), barras_sistema)
    fmt.Println(corrente_de_sequencia_na_linha)
    fmt.Println("\nAs correntes de fase nas linhas são:")
    corrente_de_fase_na_linha := falta.Corrente_na_linha(corrente_de_sequencia_na_linha)
    fmt.Println(corrente_de_fase_na_linha)

    fmt.Println("\nO tempo de execução foi de", time.Since(start))
}