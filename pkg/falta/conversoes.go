package falta


import (
	"math"
)


var a = complex(-(1.0/2.0), math.Sqrt(3.0/2.0))


func Sequencia_para_fase(componente Componente_de_sequencia) (Componente_de_fase) {

	componente_fase := Componente_de_fase{
		A:	componente.Sequencia_zero + componente.Sequencia_positiva + componente.Sequencia_negativa,
		B:	componente.Sequencia_zero + a*a * componente.Sequencia_positiva + a * componente.Sequencia_negativa,
		C:	componente.Sequencia_zero + a * componente.Sequencia_positiva + a*a * componente.Sequencia_negativa,
	}

	return componente_fase
}
