package zbus

import (
	"fmt"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
)

type Posicao_zbus struct {
	Posicao int
}

type Matrix = [][]complex128

func Zbus(elementos_tipo_1 []barra.Dados_de_linha, elementos_tipo_2_3 map[string]barra.Dados_de_linha, tamanho_do_sistema int) (Matrix, Matrix, map[string]Posicao_zbus) {
	var zbus_positiva = Preenche_matriz_com_zeros(tamanho_do_sistema)
	var zbus_zero = Preenche_matriz_com_zeros(tamanho_do_sistema)

	var elementos_tipo_3 []barra.Dados_de_linha
	var barras_adicionadas = make(map[string]Posicao_zbus)
	var posicao = 0

	// Adiciona os elementos do tipo 1
	// Loop passando por todos os elementos do tipo 1, e adicionando cada um na matriz Zbus
	for _, dados_linha := range elementos_tipo_1 {

		zbus_positiva = Adiciona_elemento_tipo_1_na_zbus(zbus_positiva, posicao, dados_linha.Impedancia_positiva)
		zbus_zero = Adiciona_elemento_tipo_1_na_zbus(zbus_zero, posicao, dados_linha.Impedancia_zero)

		fmt.Println("Adicionado elemento tipo 1 -> Barra: "+dados_linha.De+"\t\tImpedancia:", dados_linha.Impedancia_zero)

		barras_adicionadas[dados_linha.De] = Posicao_zbus{
			Posicao: posicao,
		}

		posicao++
	}

	// Adiciona os elementos do tipo 2
	// Valida se o elemento é do tipo 2, caso seja, adiciona na Zbus
	// Caso o elemento seja do tipo 3, ele adiciona em uma lista que será utilizada futuramente para adicionar os elementos tipo 3
	for len(elementos_tipo_2_3) != 0 {
		for nome_linha, linha := range elementos_tipo_2_3 {

			_, existe_de := barras_adicionadas[linha.De]
			_, existe_para := barras_adicionadas[linha.Para]

			if existe_de && existe_para {
				elementos_tipo_3 = append(elementos_tipo_3, linha)
				delete(elementos_tipo_2_3, nome_linha)

			} else if existe_de {
				zbus_positiva = Adiciona_elemento_tipo_2_na_zbus(zbus_positiva, barras_adicionadas[linha.De].Posicao, posicao, linha.Impedancia_positiva)
				zbus_zero = Adiciona_elemento_tipo_2_na_zbus(zbus_zero, barras_adicionadas[linha.De].Posicao, posicao, linha.Impedancia_zero)

				fmt.Println("Adicionado elemento tipo 2 -> Linha: "+linha.De+"-"+linha.Para+"\tImpedancia:", linha.Impedancia_zero)

				barras_adicionadas[linha.Para] = Posicao_zbus{
					Posicao: posicao,
				}

				delete(elementos_tipo_2_3, nome_linha)
				posicao++

			} else if existe_para {
				zbus_positiva = Adiciona_elemento_tipo_2_na_zbus(zbus_positiva, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_positiva)
				zbus_zero = Adiciona_elemento_tipo_2_na_zbus(zbus_zero, barras_adicionadas[linha.Para].Posicao, posicao, linha.Impedancia_zero)

				fmt.Println("Adicionado elemento tipo 2 -> Linha: "+linha.De+"-"+linha.Para+"\tImpedancia:", linha.Impedancia_zero)

				barras_adicionadas[linha.De] = Posicao_zbus{
					Posicao: posicao,
				}

				delete(elementos_tipo_2_3, nome_linha)
				posicao++

			}
		}
	}

	// Com a lista criada de elementos do tipo 3, adicionamos na Zbus
	for x := 0; x < len(elementos_tipo_3); x++ {
		linha := elementos_tipo_3[x]

		fmt.Println("Adicionado elemento tipo 3 -> Linha: "+linha.De+"-"+linha.Para+" \tImpedancia:", linha.Impedancia_zero, " \tRealizando redução de Kron")

		zbus_positiva = Adiciona_elemento_tipo_3_com_reducao_de_kron(
			zbus_positiva,
			barras_adicionadas[linha.De].Posicao,
			barras_adicionadas[linha.Para].Posicao,
			linha.Impedancia_positiva,
			tamanho_do_sistema)
		zbus_zero = Adiciona_elemento_tipo_3_com_reducao_de_kron(
			zbus_zero,
			barras_adicionadas[linha.De].Posicao,
			barras_adicionadas[linha.Para].Posicao,
			linha.Impedancia_zero,
			tamanho_do_sistema)
	}

	return zbus_positiva, zbus_zero, barras_adicionadas
}

func Preenche_matriz_com_zeros(tamanho int) Matrix {
	var matrix_com_zeros = make(Matrix, 0)

	// Adiciona elementos 0 na matriz zbus
	for i := 0; i < tamanho; i++ {
		temp := make([]complex128, 0)
		for j := 0; j < tamanho; j++ {
			temp = append(temp, 0)
		}

		matrix_com_zeros = append(matrix_com_zeros, Matrix{temp}...)
	}

	return matrix_com_zeros
}

func Aumenta_tamanho_da_matriz(matrix Matrix) Matrix {

	temp := make([]complex128, 0)
	for j := 0; j <= len(matrix); j++ {
		temp = append(temp, 0)
	}

	matrix = append(matrix, Matrix{temp}...)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = append(matrix[i], 0)
	}

	return matrix
}
