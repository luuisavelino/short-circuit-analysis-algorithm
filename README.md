# short-circuit-analysis-algorithm

Development of the activity to complement the workload of Power Systems Analysis 2

## Estrutura de pastas

Aqui será um resumo da estrutura de pasta criada

### Main

Pasta de comando
    main -> arquivo main, onde chama todo o código

### Data

arquivos xlsx que serão analisados

### Internal

possui arquivos com funcionalidades internas
    funcoes_gerais -> possui as funções gerais utilizadas por todo o código

### Pkg

possui as funções de construção do aplicativo

#### Analise

contem o codigo que realizará a análise do sistema

- analise -> Analise de faltas
- tempo_critico -> Analise de tempo de critico do sistema

#### Barra

Pega as informações do arquivo xlsx e coloca dentro de estruturas

#### Falta

realiza os calculos de falta do sistema

- calculo_faltas -> realiza o calculo de tensão e corrente de sequencia
- conversoes -> Realiza a transformação de fase para sequencia e o inverso
- corrente_de_falta -> realiza o calculo das correntes de falta para diversos tipos de falta

#### Zbus

Realiza a desenvolvimento da matriz zbus do sistema

- adiciona_elementos -> Adiciona os elementos tipo 1, 2 e 3 na matriz zbus
- cria_zbus -> recebe os arquivos de linhas e barras e constrói a matriz zbus
