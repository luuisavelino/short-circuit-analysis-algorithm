package zbus

import (
	"fmt"
	"math"
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
	
	var matriz_B [6][1]float64
	var matriz_C [1][6]float64
	var matriz_D float64
	var matriz_BxDxC matrix
	var zbus_reduzida matrix	

	for x := 0; x < 6; x++ {
		matriz_B[x][0] = zbus[x][posicao_barra_de] - zbus[x][posicao_barra_para]
		matriz_C[0][x] = zbus[posicao_barra_de][x] - zbus[posicao_barra_para][x]
	}

	fmt.Println(matriz_B)
	fmt.Println(matriz_C)
	
	matriz_D = zbus[posicao_barra_de][posicao_barra_de] + zbus[posicao_barra_para][posicao_barra_para] + (2 * zbus[posicao_barra_de][posicao_barra_para]) + impedancia
	
	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {
			matriz_BxDxC[x][y] = (matriz_B[x][0] * matriz_C[0][y]) / matriz_D
		}
	}

	for x := 0; x < 6; x++ {
		for y := 0; y < 6; y++ {
			zbus_reduzida[x][y] = math.Round((zbus[x][y] - matriz_BxDxC[x][y])*10000)/10000
		}
	}

	return zbus_reduzida
}