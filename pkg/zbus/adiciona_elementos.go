package zbus

func Adiciona_elemento_tipo_1_na_zbus(zbus Matrix, posicao int, impedancia complex128) Matrix {
	zbus[posicao][posicao] = impedancia

	return zbus
}

func Adiciona_elemento_tipo_2_na_zbus(zbus Matrix, posicao_barra_conectada int, posicao int, impedancia complex128) Matrix {

	for x := 0; x < posicao; x++ {
		zbus[x][posicao] = zbus[x][posicao_barra_conectada]
		zbus[posicao][x] = zbus[posicao_barra_conectada][x]
	}

	zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

	return zbus
}

func Adiciona_elemento_tipo_3_com_reducao_de_kron(zbus Matrix, posicao_barra_de int, posicao_barra_para int, impedancia complex128, tamanho_do_sistema int) Matrix {

	var matriz_B []complex128
	var matriz_C []complex128
	var matriz_D complex128
	var zbus_reduzida = Preenche_matriz_com_zeros(tamanho_do_sistema)

	for x := 0; x < tamanho_do_sistema; x++ {
		matriz_B = append(matriz_B, zbus[x][posicao_barra_de]-zbus[x][posicao_barra_para])
		matriz_C = append(matriz_C, zbus[posicao_barra_de][x]-zbus[posicao_barra_para][x])
	}

	matriz_D = zbus[posicao_barra_de][posicao_barra_de] + zbus[posicao_barra_para][posicao_barra_para] - (2 * zbus[posicao_barra_de][posicao_barra_para]) + impedancia

	for x := 0; x < tamanho_do_sistema; x++ {
		for y := 0; y < tamanho_do_sistema; y++ {
			zbus_reduzida[x][y] = (zbus[x][y] - ((matriz_B[x] * matriz_C[y]) / matriz_D))
		}
	}

	return zbus_reduzida
}
