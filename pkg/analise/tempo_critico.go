package analise

import (
	"math"
	"math/cmplx"
)

const (
	frequencia_da_rede = 60
	H = 5.0
	V1 = 1.05
	V2 = 1.0
	potencia_media = 0.8
)


type Dados_geracao struct {
	Modulo_corrente_de_saida_gerador	float64
	Angulo_corrente_de_saida_gerador	float64
	Modulo_tensao_de_saida_gerador		float64
	Angulo_tensao_de_saida_gerador		float64
	Angulo_tensao_barra_curto			float64
	Potencia_media_pre_barra_gerador	float64
	Potencia_media_pre_barra_curto		float64
	Potencia_media_pos_barra_curto		float64
	Angulo_delta_maximo					float64
	Angulo_delta_critico				float64
	Tempo_maximo						float64
}


func Tempo_critico(X_1_2_pre_falta, X_1_2_pos_falta, Xd_gerador complex128) (Dados_geracao) {

	var geracao Dados_geracao

	var Ws float64 = 2 * math.Pi * frequencia_da_rede
	var X_equivalente_pre_falta float64 = cmplx.Abs(X_1_2_pre_falta + Xd_gerador)
	var X_equivalente_pos_falta float64 = cmplx.Abs(X_1_2_pos_falta + Xd_gerador)

	geracao.Angulo_tensao_barra_curto = math.Asin((potencia_media * cmplx.Abs(X_1_2_pre_falta))/(V1 * V2))

	corrente_saida_gerador := ((cmplx.Rect(V1, geracao.Angulo_tensao_barra_curto) - complex(V2, 0)) / X_1_2_pre_falta)
	
	geracao.Modulo_corrente_de_saida_gerador, geracao.Angulo_corrente_de_saida_gerador = cmplx.Polar(corrente_saida_gerador)

	geracao.Modulo_tensao_de_saida_gerador, geracao.Angulo_tensao_de_saida_gerador = cmplx.Polar(cmplx.Rect(V1, geracao.Angulo_tensao_barra_curto) + Xd_gerador * corrente_saida_gerador)

	geracao.Potencia_media_pre_barra_gerador = geracao.Modulo_tensao_de_saida_gerador * V1 / cmplx.Abs(Xd_gerador)
	geracao.Potencia_media_pre_barra_curto = geracao.Modulo_tensao_de_saida_gerador * V2 / X_equivalente_pre_falta

	geracao.Potencia_media_pos_barra_curto = geracao.Modulo_tensao_de_saida_gerador * V2 / X_equivalente_pos_falta

	geracao.Angulo_delta_maximo = math.Pi - math.Asin(potencia_media / geracao.Potencia_media_pos_barra_curto)

	geracao.Angulo_delta_critico = math.Acos((potencia_media * (geracao.Angulo_delta_maximo - geracao.Angulo_tensao_de_saida_gerador) + geracao.Potencia_media_pos_barra_curto * math.Cos(geracao.Angulo_delta_maximo)) / geracao.Potencia_media_pos_barra_curto)

	geracao.Tempo_maximo = math.Sqrt((4 * H * (geracao.Angulo_delta_critico - geracao.Angulo_tensao_de_saida_gerador)) / (Ws * potencia_media))

	return geracao
}
