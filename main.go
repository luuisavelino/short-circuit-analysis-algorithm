package main

import (
	"math"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"

	"github.com/luuisavelino/short-circuit-analysis-algorithm/internal/geral"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/analise"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/barra"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/falta"
	"github.com/luuisavelino/short-circuit-analysis-algorithm/pkg/zbus"

	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var ponto_cc_string, barra_curto_circuito, nome_do_arquivo, barra_de, barra_para string
	var ponto_cc int
	var escolha string
	var arquivos []string

	for {
		fmt.Println("Escolha uma das opções:")
		fmt.Println("1 - Definir o arquivo do sistema a ser analisado\n2 - Sair")
		fmt.Scanln(&escolha)
		if escolha == "2" {
			break
		} else if escolha != "1" {
			fmt.Println("Opção inválida")
			continue
		}

		// Pegando o arquivo que contem os dados do sistema
		fmt.Println("Escolha o arquivo a ser analisado:")
		files, _ := os.ReadDir("./data/")
		for pos, file := range files {
			fmt.Printf("(%v) - %v\n", pos, file.Name())
			arquivos = append(arquivos, file.Name())
		}

		fmt.Scanln(&escolha)
		numero_escolhido, _ := strconv.Atoi(escolha)
		nome_do_arquivo = arquivos[numero_escolhido]

		tabela_dados, err := excelize.OpenFile("../data/" + nome_do_arquivo)
		if err != nil {
			fmt.Println("Arquivo não encontrado, tente novamente")
			continue
		}

		tamanho_do_sistema, _ := geral.Quantidade_de_barras(tabela_dados)

		for {
			fmt.Println("Escolha o tipo de analise que deseja realizar:")
			fmt.Println("1 - Realizar analise de curto-circuito\n2 - Voltar")
			fmt.Scanln(&escolha)
			if escolha == "1" {
				// Curto circuito do sistema
				fmt.Println("Definindo onde que ocorrerá o curto-circuito")
				fmt.Println("Barra De:")
				fmt.Scanln(&barra_de)
				fmt.Println("Barra Para:")
				fmt.Scanln(&barra_para)
				fmt.Println("Ponto do curto circuito (0 a 100):")
				fmt.Scanln(&ponto_cc_string)
				ponto_cc, _ = strconv.Atoi(ponto_cc_string)

				curto_circuito := barra.Ponto_curto_circuito{
					De:    barra_de,
					Para:  barra_para,
					Ponto: ponto_cc,
				}

				if curto_circuito.Ponto == 0 {
					barra_curto_circuito = curto_circuito.De
				} else if curto_circuito.Ponto == 100 {
					barra_curto_circuito = curto_circuito.Para
				} else {
					barra_curto_circuito = "barra_curto_circuito"
					tamanho_do_sistema++
				}

				inicio_calculo_zbus := time.Now()

				elementos_tipo_1 := barra.Elementos_tipo_1(tabela_dados)
				elementos_tipo_2_3 := barra.Elementos_tipo_2_3(tabela_dados, curto_circuito)

				//barra.Linhas_sistema(barras_do_sistema ,elementos_tipo_2_3)

				// Constroi a matriz Zbus do sistema
				zbus_positiva, zbus_zero, barras_sistema := zbus.Zbus(elementos_tipo_1, elementos_tipo_2_3, tamanho_do_sistema)

				analise.Mostra_matriz_zbus(zbus_positiva, tamanho_do_sistema)
				analise.Mostra_matriz_zbus(zbus_zero, tamanho_do_sistema)

				fmt.Println("\nO tempo de calculo da Zbus foi de", time.Since(inicio_calculo_zbus))

				elementos_tipo_2_3 = barra.Elementos_tipo_2_3(tabela_dados, curto_circuito)

				for {
					fmt.Println("Escolha o tipo curto-circuito a ser analisado:")
					fmt.Println("\nEscolha o tipo de falta:\n  (1) - Monofasica\n  (2) - Bifasica\n  (3) - Bifasico Terra\n  (4) - Trifasica\n  (5) - Tempo crítico\n  (6) - Voltar")
					fmt.Scanln(&escolha)

					inicio_calculo_falta := time.Now()
					if escolha == "1" {
						Icc_monofasica_fase, Icc_monofasica_sequencia := falta.Corrente_falta_monofasica(zbus_positiva, zbus_zero, barras_sistema, barra_curto_circuito)
						analise.Analise_curto_circuito(zbus_positiva, zbus_zero, elementos_tipo_2_3, barras_sistema, Icc_monofasica_fase, Icc_monofasica_sequencia, curto_circuito)

					} else if escolha == "2" {
						Icc_bifasica_fase, Icc_bifasica_sequencia := falta.Corrente_falta_bifasica(zbus_positiva, barras_sistema, barra_curto_circuito)
						analise.Analise_curto_circuito(zbus_positiva, zbus_zero, elementos_tipo_2_3, barras_sistema, Icc_bifasica_fase, Icc_bifasica_sequencia, curto_circuito)

					} else if escolha == "3" {
						Icc_bifasica_fase, Icc_bifasica_sequencia := falta.Corrente_falta_bifasico_terra(zbus_positiva, zbus_zero, barras_sistema, barra_curto_circuito)
						analise.Analise_curto_circuito(zbus_positiva, zbus_zero, elementos_tipo_2_3, barras_sistema, Icc_bifasica_fase, Icc_bifasica_sequencia, curto_circuito)

					} else if escolha == "4" {
						falta.Falta_trifasica(zbus_positiva, barras_sistema, barra_curto_circuito, barra.Elementos_tipo_2_3(tabela_dados, curto_circuito))

					} else if escolha == "5" {

						curto_circuito.Ponto = 999
						elementos_tipo_2_3 := barra.Elementos_tipo_2_3(tabela_dados, curto_circuito)

						zbus_pos_retirada_da_linha, _, barras_sistema_pos := zbus.Zbus(elementos_tipo_1, elementos_tipo_2_3, tamanho_do_sistema)

						fmt.Println("\nAnáise de tempo crítico para abertura dos disjuntores")
						fmt.Println("Gerador \t Angulo Crítico (°) \t\t Tempo máximo (s)")
						for _, gerador := range elementos_tipo_1 {
							posicao_barra_gerador_pre := barras_sistema[gerador.De].Posicao
							posicao_barra_cc_pre := barras_sistema[barra_curto_circuito].Posicao

							posicao_barra_gerador_pos := barras_sistema_pos[gerador.De].Posicao
							posicao_barra_cc_pos := barras_sistema_pos[barra_curto_circuito].Posicao

							geracao := analise.Tempo_critico(
								zbus_positiva[posicao_barra_gerador_pre][posicao_barra_cc_pre],
								zbus_pos_retirada_da_linha[posicao_barra_gerador_pos][posicao_barra_cc_pos],
								gerador.Impedancia_positiva)

							fmt.Println(gerador.De, "\t\t", geral.Round(geracao.Angulo_delta_critico*180/math.Pi, 4), "\t\t\t", geral.Round(geracao.Tempo_maximo, 4))

						}
					} else if escolha == "6" {
						break

					} else {
						fmt.Println("Tipo de falta não encontrada")
					}

					fmt.Println("\nO tempo de calculo da falta foi de", time.Since(inicio_calculo_falta))
				}
			} else if escolha == "2" {
				break

			} else {
				fmt.Println("Tipo de analise não encontrada")
			}
		}
	}

	fmt.Println("\nO tempo de execução do programa foi de", time.Since(start))
}
