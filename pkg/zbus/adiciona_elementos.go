package zbus


func Adiciona_elemento_tipo_1_na_zbus(zbus [40][40]float64, posicao int, impedancia float64) [40][40]float64 {
	zbus[posicao][posicao] = impedancia

	return zbus
}

func Adiciona_elemento_tipo_2_na_zbus(zbus [40][40]float64, posicao_barra_conectada int, posicao int, impedancia float64) [40][40]float64 {

	for x := 0; x < posicao; x++ {
		zbus[x][posicao] = zbus[x][posicao_barra_conectada]
		zbus[posicao][x] = zbus[posicao_barra_conectada][x]
	}

	zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

	return zbus
}

func Adiciona_elemento_tipo_3_na_zbus(zbus [40][40]float64, posicao_barra_de int, posicao_barra_para int, posicao int, impedancia float64) [40][40]float64 {

	for x := 0; x < posicao; x++ {
		zbus[x][posicao] = zbus[x][posicao_barra_de] - zbus[x][posicao_barra_para]
		zbus[posicao][x] = zbus[posicao_barra_de][x] - zbus[posicao_barra_para][x]
	}

	zbus[posicao][posicao] = zbus[posicao_barra_de][posicao_barra_de] + zbus[posicao_barra_para][posicao_barra_para] + (2 * zbus[posicao_barra_de][posicao_barra_para]) + impedancia

	return zbus
}
