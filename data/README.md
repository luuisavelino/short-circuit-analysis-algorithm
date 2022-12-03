# Formato e especificações do tipo de entrada de dados

Descritivo detalhado sobre a entrada de dados para a execução do serviço

## Formato

O formato do arquivo deve ser .xlsx

## Tabelas

É necessário ter 4 tabelas dentro do arquivo, sendo elas:

- [0] Dados de Barra
- [1] Dados de Linha
- [2] Dados dos Geradores
- [3] Dados dos Transformadores

### Dados de Barra

Necessáriamente precisa possuir o seguinte formato

| Numero da Barra | Nome | [coluna 2] | Tensão |
|--:|---------|:--:|:----|
| Linha 2 |
| Linha 3 |

Sendo que a primeira linha deve ocupar dois espaços. Assim, os dados começam a partir da linha 2 (sendo 0 a primeira linha)

---

### Dados de Linha

Necessáriamente precisa possuir o seguinte formato

| De | Para | Nome | Tensão | [coluna 4] | R+ | X+ | [coluna 7] | [coluna 8] | [coluna 9] | R0 | X0 |
|--:|---------|:--:|:----|--:|---------|:--:|:----|--:|---------|:--:|:----|
| Linha 2 |
| Linha 3 |

Sendo que a primeira linha deve ocupar dois espaços. Assim, os dados começam a partir da linha 2 (sendo 0 a primeira linha)

---

### Dados dos Geradores

Necessáriamente precisa possuir o seguinte formato

| Numero da Barra | Nome | [coluna 2] | Xd |
|--:|---------|:--:|:----|
| Linha 1 |
| Linha 2 |

Sendo que a primeira linha deve ocupar apenas um espaço. Assim, os dados começam a partir da linha 1 (sendo 0 a primeira linha)

---

### Dados dos Transformadores

Necessáriamente precisa possuir o seguinte formato

| De | Para | Nome | [coluna 3] | [coluna 4] | [coluna 5] | R+ | X+ |
|--:|---------|:--:|:----|--:|---------|:--:|:----|
| Linha 2 |
| Linha 3 |

Sendo que a primeira linha deve ocupar dois espaços. Assim, os dados começam a partir da linha 2 (sendo 0 a primeira linha)
