package geral

import (
	"testing"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
)

func TestDivisaoPorZero(t *testing.T) {
	if geral.Valida_divisao_por_0(1,0) != 0 {
		t.Errorf(`geral.Valida_divisao_por_0(1,0) = 0`)
	}
}

