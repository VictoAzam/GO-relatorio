# 🚀 Exercícios e Conceitos em Go

Aprendizado prático da linguagem Go por meio de 26 exercícios fundamentais e exemplos de conceitos (estruturas, funções, ponteiros, interfaces, enums etc.).

**Período:** 19/08/2025 — 28/11/2025
**Autor:** Victor Hugo Azambuja
**Assistente:** GitHub Copilot (GPT-5)

</div>

---

## 📦 Visão Geral

Este repositório organiza cada exercício em uma pasta separada (isolamento de `main.go`) e agrupa exemplos de conceitos da linguagem em `conceitos/`. O objetivo é ter código simples, direto e comentado em português para estudo e referência rápida.

---

## 🗂 Estrutura

```
exercicios/
	ex1/  (Soma de inteiros)
	ex2/  (Divisão inteira / resto)
	...
	ex26/ (Algoritmo genético simples)
conceitos/
	hello_world/
	values/
	variables/
	constants/
	for_loop/
	if_else/
	switch_case/
	arrays/
	slices/
	maps/
	functions/
	multiple_returns/
	variadic_functions/
	closures/
	recursion/
	range_builtin/
	pointers/
	strings_runes/
	structs/
	methods/
	interfaces/
	enums_iota/
RELATORIO.md
RELATORIO_RENTRY.md
RELATORIO_UNIFICADO.md
README.md
go.mod
```

---

## ▶️ Execução Rápida (PowerShell)

```powershell
# Exercício 1 (Soma)
go run .\exercicios\ex1\main.go

# Torre de Hanoi (Ex10)
go run .\exercicios\ex10\main.go

# Algoritmo genético (Ex26)
go run .\exercicios\ex26\main.go

# Exemplo de conceito: Interfaces
go run .\conceitos\interfaces\main.go
```

Se algum exercício solicitar entrada, digite conforme instruções no terminal.

---

## 🧪 Tabela dos Exercícios


| Nº | Pasta | Tema                    | Foco Principal             |
| --- | ----- | ----------------------- | -------------------------- |
| 1   | ex1   | Soma de inteiros        | Entrada/saída básica     |
| 2   | ex2   | Divisão + resto        | Operações aritméticas   |
| 3   | ex3   | Sucessor/antecessor     | Aritmética simples        |
| 4   | ex4   | Classificação número | Condicionais               |
| 5   | ex5   | Número primo           | Laços e eficiência       |
| 6   | ex6   | Ordenar números        | sort.Ints                  |
| 7   | ex7   | Ordenar caracteres      | Manipulação de string    |
| 8   | ex8   | Árvore de decisão     | Lógica binária           |
| 9   | ex9   | Valor e endereço       | Ponteiros                  |
| 10  | ex10  | Torre de Hanoi          | Recursão                  |
| 11  | ex11  | Dia da semana (Zeller)  | Fórmula matemática       |
| 12  | ex12  | Igualdade               | Booleano                   |
| 13  | ex13  | Moda                    | Map/frequência            |
| 14  | ex14  | Palíndromo             | Normalização texto       |
| 15  | ex15  | Área retângulo        | Função pura              |
| 16  | ex16  | Conversão temperatura  | Fórmulas C/F/K            |
| 17  | ex17  | Jogo adivinhação      | Aleatoriedade              |
| 18  | ex18  | Imagem RGB              | image/color/jpeg           |
| 19  | ex19  | Vogais/consoantes       | Unicode                    |
| 20  | ex20  | Contar palavra          | strings.Fields             |
| 21  | ex21  | Fatorial big.Int        | Inteiros grandes           |
| 22  | ex22  | Olá Mundo              | Saída básica             |
| 23  | ex23  | IMC                     | Cálculo + classificação |
| 24  | ex24  | MMC                     | Euclides (MDC)             |
| 25  | ex25  | Média                  | Loop e divisão            |
| 26  | ex26  | Algoritmo genético     | Seleção + mutação      |

---

## 📚 Conceitos Demonstrados

- Tipos básicos, constantes e variáveis
- Controle de fluxo (for, if/else, switch)
- Arrays, slices, maps
- Funções (retornos múltiplos, variádicas), closures, recursão
- Ponteiros e manipulação de memória indireta
- Strings, runes e Unicode
- Structs e métodos
- Interfaces (polimorfismo) e enums com `iota`

Em planejamento: generics, concorrência (goroutines/channels), testes e benchmarks.

---

## 💡 Destaques Técnicos

- Uso de `sort`, `unicode`, `image`, `math/big`
- Recursão aplicada (Hanoi, fatorial)
- Geração de imagem RGB sintética
- Algoritmo genético com seleção por torneio

---

## 🛠 Pré-requisitos

- Go 1.22+ instalado
- Terminal PowerShell (Windows) ou outro shell compatível

Verifique a versão:

```powershell
go version
```

---

## 🧭 Próximos Passos Sugeridos

- Adicionar exemplos de goroutines e channels
- Implementar testes (`go test`) para primalidade, moda e palíndromo
- Benchmarks (`go test -bench .`) para comparar abordagens
- Introduzir generics em funções utilitárias (ex: Min/Max)

---

## 📖 Relatórios

- `RELATORIO.md`: resumo tabular inicial
- `RELATORIO_RENTRY.md`: versão narrativa expandida
- `RELATORIO_UNIFICADO.md`: documento completo integrado

---

## 🙌 Créditos

Projeto educacional desenvolvido por Victor Hugo Azambuja com apoio do GitHub Copilot.

---

## 🔗 Referências

- Documentação Go: https://go.dev
- Pacotes padrão: `fmt`, `sort`, `unicode`, `image`, `math/big`
- Fórmula de Zeller (cálculo de dia da semana)
- WHO (categorias IMC)

---

## ✅ Execução Rápida (Resumo)

```powershell
# Listar versão Go
go version

# Rodar exercício específico
go run .\exercicios\ex5\main.go

# Rodar conceito
go run .\conceitos\closures\main.go
```

Bom estudo! ✨
