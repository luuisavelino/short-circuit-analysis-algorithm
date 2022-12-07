package barra

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Barra struct {
	Key		string
	Color	string
} 

type Linha struct {
    From 	string
	To 		string
}

// { key: "Barra", color: "black" }
// { from: "Barra_de", to: "Barra_para" }

func Linhas_sistema(barras []string, elementos_2_3 map[string]Dados_de_linha) {
    var dados_linhas []Linha
	var dados_barras []Barra

	for _, barra := range barras {
		dados_barras = append(dados_barras, Barra{Key: barra, Color: "grey"})
	}

	for _, linha := range elementos_2_3 {
		dados_linhas = append(dados_linhas, Linha{From: linha.De , To: linha.Para})
	}

    a, _ := json.Marshal(dados_linhas)
    b, _ := json.Marshal(dados_barras)

    fmt.Println(strings.ToLower(string(b)))
    fmt.Println(strings.ToLower(string(a)))
}