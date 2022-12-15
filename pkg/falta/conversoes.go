package falta

import (
	"math"
	"math/cmplx"
)

var a complex128 = cmplx.Rect(1, 120*math.Pi/180)

func Sequencia_para_fase(componente Componente_de_sequencia) Componente_de_fase {

	componente_fase := Componente_de_fase{
		A: componente.Sequencia_zero + componente.Sequencia_positiva + componente.Sequencia_negativa,
		B: componente.Sequencia_zero + a*a*componente.Sequencia_positiva + a*componente.Sequencia_negativa,
		C: componente.Sequencia_zero + a*componente.Sequencia_positiva + a*a*componente.Sequencia_negativa,
	}

	return componente_fase
}
