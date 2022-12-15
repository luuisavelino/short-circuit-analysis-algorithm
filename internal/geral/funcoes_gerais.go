package geral

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// Entrada:     Resistencia, Reatancia, e a impedância atual
// Processo:    Realiza o calculo da impedância
//              Caso já exista uma impedância atual, ele irá realizar o paralelo entre eles
// Saída:       Retorna o valor da impedância
func Impedancia(resistencia_linha string, reatancia_linha string, impedancia_atual complex128) complex128 {
	var resistencia, _ = strconv.ParseFloat(resistencia_linha, 64)
	var reatancia, _ = strconv.ParseFloat(reatancia_linha, 64)

	impedancia := complex(resistencia, reatancia)

	if impedancia_atual != 0 {
		impedancia = (impedancia * impedancia_atual) / (impedancia + impedancia_atual)
	}

	return impedancia
}

// Irá converter o tipo string para o tipo float64
func String_para_float(grandeza_str string) float64 {
	grandeza, _ := strconv.ParseFloat(grandeza_str, 64)

	return grandeza
}

// Realiza a validação do erro
func Valida_erro(err error) {
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

// Arredondamento para uma determinada quantidade de casas decimais
func Round(valor float64, casas float64) float64 {
	ratio := math.Pow(10, float64(casas))
	return math.Round(valor*ratio) / ratio
}

// Arredonda numeros complexos
func Round_cmplx(valor complex128, casas float64) complex128 {

	valor_real := Round(real(valor), casas)
	valor_imag := Round(imag(valor), casas)

	return complex(valor_real, valor_imag)
}

func Retangular_To_Polar(valor complex128) string {
	modulo, angulo := cmplx.Polar(valor)

	modulo = Round(modulo, 4)
	angulo = Round(angulo*180/math.Pi, 4)

	if modulo == 0 {
		angulo = 0
	}

	return fmt.Sprintf("%g ∠%g", modulo, angulo)
}

func Quantidade_de_barras(tabela_excel *excelize.File) (int, []string) {
	barras, _ := tabela_excel.GetRows(tabela_excel.GetSheetList()[0])

	tamanho_do_sistema := len(barras) - 2
	var barras_do_sistema []string

	for x := 2; x < len(barras); x++ {
		barras_do_sistema = append(barras_do_sistema, (barras[x][0]))
	}

	return tamanho_do_sistema, barras_do_sistema
}

func Valida_divisao_por_0(num complex128, den complex128) complex128 {

	if den == 0 {
		return 0
	}

	return num / den
}
