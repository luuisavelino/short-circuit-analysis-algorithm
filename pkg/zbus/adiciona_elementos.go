package zbus

func adiciona_elemento_2(barra string, barra_conectada string, zbus [50][50]float64, impedancia float64, posicao int, barras_adicionadas map[string]posicao_zbus) [50][50]float64 {

    posicao_barra_conectada := barras_adicionadas[barra_conectada].Posicao

    for x := 0; x < posicao; x++ {
        zbus[x][posicao] = zbus[x][posicao_barra_conectada]
        zbus[posicao][x] = zbus[posicao_barra_conectada][x]
    }

    zbus[posicao][posicao] = impedancia + zbus[posicao_barra_conectada][posicao_barra_conectada]

    return zbus
}