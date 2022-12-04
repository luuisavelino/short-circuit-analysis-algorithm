package analise

import (
	"fmt"
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

func Tempo_critico(X_1_2_pre_falta, X_1_2_pos_falta, Xd_gerador float64) {

	var Ws float64 = 2 * math.Pi * frequencia_da_rede

		
	Xd := complex(0, X_1_2_pre_falta)

	var X_equivalente_pre_falta float64 = X_1_2_pre_falta + Xd_gerador
	var X_equivalente_pos_falta float64 = X_1_2_pos_falta + Xd_gerador

	

	angulo_tensao_barra_gerador := math.Asin((potencia_media * cmplx.Abs(Xd))/(V1 * V2))

	modulo_corrente_saida_gerador, angulo_corrente_saida_gerador := cmplx.Polar((cmplx.Rect(V1, angulo_tensao_barra_gerador) - complex(V2, 0)) / Xd)
	fmt.Println("A corrente de saída do gerador é de", modulo_corrente_saida_gerador,"A, com angulo de", angulo_corrente_saida_gerador * 180 / math.Pi, "(°)")
	
	modulo_tensao_gerador, angulo_tensao_gerador := cmplx.Polar(cmplx.Rect(V1, angulo_tensao_barra_gerador) + complex(0, Xd_gerador) * ((cmplx.Rect(V1, angulo_tensao_barra_gerador) - complex(V2, 0)) / Xd))
	fmt.Println("O angulo da tensão do gerador é de", angulo_tensao_barra_gerador * 180 / math.Pi, "(°)")
	
	potencia_media_pre_falta_barra_1 := modulo_tensao_gerador * V1 / Xd_gerador
	potencia_media_pre_falta_barra_2 := modulo_tensao_gerador * V2 / X_equivalente_pre_falta

	fmt.Println("A potencia de pré-falta na barra 1 é", potencia_media_pre_falta_barra_1,"sin(delta)")
	fmt.Println("A potencia de pré-falta na barra 2 é", potencia_media_pre_falta_barra_2,"sin(delta)")

	potencia_media_pos_falta_barra_2 := modulo_tensao_gerador * V2 / X_equivalente_pos_falta
	fmt.Println("A potencia de pós-falta na barra 2 é", potencia_media_pos_falta_barra_2,"sin(delta)")


	detla_max := math.Pi - math.Asin(potencia_media / potencia_media_pos_falta_barra_2)
	fmt.Println("O Angulo delta zero é de", angulo_tensao_gerador,"(°)")
	fmt.Println("O Angulo delta máximo é de", detla_max,"(°)")

	delta_c := math.Acos((potencia_media * (detla_max - angulo_tensao_gerador) + potencia_media_pos_falta_barra_2 * math.Cos(detla_max)) / potencia_media_pos_falta_barra_2)
	fmt.Println("O Angulo delta crítico é de", delta_c,"(°)")
	
	tempo := math.Sqrt((4 * H * (delta_c - angulo_tensao_gerador)) / (Ws * potencia_media))
	fmt.Println("O tempo máximo de abertura é de", tempo,"(s)")


	//return delta_c, tempo
}
