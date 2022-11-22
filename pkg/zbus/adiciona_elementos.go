package zbus

import (
	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)

func Adiciona_elemento_tipo_1_na_zbus(zbus matrix, posicao int, impedancia float64) matrix {
	zbus[posicao][posicao] = impedancia

	return zbus
}

func Adiciona_elemento_tipo_2_na_zbus(zbus matrix, posicao_barra_conectada int, posicao int, impedancia float64) matrix {

	for x := 0; x < posicao; x++ {
		zbus[x][posicao] = zbus[x][posicao_barra_conectada]
		zbus[posicao][x] = zbus[posicao_barra_conectada][x]
	}

	zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

	return zbus
}

func Adiciona_elemento_tipo_3_na_zbus(zbus matrix, posicao_barra_de int, posicao_barra_para int, posicao int, impedancia float64) matrix {

	for x := 0; x < posicao; x++ {
		zbus[x][posicao] = zbus[x][posicao_barra_de] - zbus[x][posicao_barra_para]
		zbus[posicao][x] = zbus[posicao_barra_de][x] - zbus[posicao_barra_para][x]
	}

	zbus[posicao][posicao] = zbus[posicao_barra_de][posicao_barra_de] + zbus[posicao_barra_para][posicao_barra_para] + (2 * zbus[posicao_barra_de][posicao_barra_para]) + impedancia

	return zbus
}


func Adiciona_elemento_tipo_3_com_reducao_de_kron(zbus matrix, posicao_barra_de int, posicao_barra_para int, impedancia float64) matrix {
	
	var matriz_B [Tamanho_do_sistema]float64
	var matriz_C [Tamanho_do_sistema]float64
	var matriz_D float64
	var zbus_reduzida matrix	

	for x := 0; x < Tamanho_do_sistema; x++ {
		matriz_B[x] = zbus[x][posicao_barra_de] - zbus[x][posicao_barra_para]
		matriz_C[x] = zbus[posicao_barra_de][x] - zbus[posicao_barra_para][x]
	}

	matriz_D = geral.Round(zbus[posicao_barra_de][posicao_barra_de] + zbus[posicao_barra_para][posicao_barra_para] - (2 * zbus[posicao_barra_de][posicao_barra_para]) + impedancia, 5)

	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {
			zbus_reduzida[x][y] = geral.Round((zbus[x][y] - ((matriz_B[x] * matriz_C[y]) / matriz_D)), 5)
		}
	}

	return zbus_reduzida
}
