package zbus

import (
	"fmt"
)

func Adiciona_elemento_tipo_1(zbus [30][30]float64, posicao int, impedancia float64) [30][30]float64 {
	zbus[posicao][posicao] = impedancia

	return zbus
}


func Adiciona_elemento_tipo_2(barra_conectada string, zbus [30][30]float64, impedancia float64, posicao int, barras_adicionadas map[string]posicao_zbus) [30][30]float64 {

	fmt.Println(posicao)
	
	posicao_barra_conectada := barras_adicionadas[barra_conectada].Posicao
	fmt.Println("Barra Conectada: " + barra_conectada + ". Impedancia:", impedancia, ". Posicao da barra conectada:", posicao_barra_conectada)

    for x := 0; x < posicao; x++ {
        zbus[x][posicao] = zbus[x][posicao_barra_conectada]
        zbus[posicao][x] = zbus[posicao_barra_conectada][x]
    }

    zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

    return zbus
}

func Adiciona_elemento_tipo_3() {

	return

}