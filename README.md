# 📈 Simulador de Aposentadoria — Back-end

API REST desenvolvida em **Go (Golang)** para realizar simulações financeiras voltadas ao planejamento de aposentadoria e independência financeira.

O back-end recebe os dados informados pelo usuário, processa os cálculos financeiros e retorna os resultados da simulação para o front-end.

## 🚀 Tecnologias

- **Go (Golang):** utilizado no desenvolvimento da API e implementação das regras de negócio.
- **Gin:** framework utilizado para criação do servidor HTTP e gerenciamento das rotas.
- **REST API:** utilizada para comunicação entre o front-end e o back-end.
- **Clean Architecture:** utilizada para organizar o projeto em camadas e separar responsabilidades.
- **Repository Pattern:** utilizado para abstrair a persistência dos dados.
- **Git / GitHub:** utilizados para controle de versão e gerenciamento do código.

## 📋 Funcionalidades

- Simulação financeira para aposentadoria
- Cálculo do patrimônio ao longo do tempo
- Cálculo do patrimônio ajustado pela inflação
- Cálculo de juros reais
- Cálculo do total contribuído
- Cálculo dos rendimentos obtidos
- Simulação com aportes mensais
- Cálculo do tempo até a aposentadoria
- Geração da evolução patrimonial ao longo dos anos
- Persistência dos dados da simulação
- Retorno dos resultados através de API REST

## 🏗️ Arquitetura

O projeto utiliza **Clean Architecture** para separar as responsabilidades da aplicação.

```text
simulator-api/
│
├── data/
│   ├── database/
│   └── repository/
│
├── domain/
│   ├── entities/
│   ├── repositories/
│   └── usecase/
│
├── handlers/
│
├── routes/
│
└── main.go
```

As regras de negócio ficam concentradas na camada de **use cases**, enquanto handlers e rotas ficam responsáveis pelo fluxo das requisições HTTP.

## 🔄 Fluxo da Aplicação

```text
Flutter Web
     │
     │ REST API
     ▼
   Routes
     │
     ▼
  Handlers
     │
     ▼
 Use Cases
     │
     ▼
Repository
```

O front-end envia os dados da simulação para a API, o back-end processa as regras de negócio, realiza os cálculos financeiros e retorna os resultados ao usuário.

## 📥 Dados da Simulação

A API recebe informações como:

- Patrimônio atual
- Aporte mensal
- Taxa de rentabilidade anual
- Inflação
- Idade atual
- Idade desejada para aposentadoria
- Tempo da simulação em anos

## 📤 Resultados

A partir dos dados informados, a API retorna indicadores como:

- **Patrimônio final**
- **Patrimônio final ajustado pela inflação**
- **Taxa de juros real anual**
- **Taxa de juros real mensal**
- **Total contribuído**
- **Total de rendimentos**
- **Tempo até a aposentadoria**
- **Evolução patrimonial ao longo dos anos**

## 🔌 Endpoints

A API disponibiliza atualmente os seguintes endpoints:

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `GET` | `/test` | Verifica se a API está funcionando |
| `POST` | `/simulation` | Executa uma nova simulação financeira |

## 🎯 Objetivo do Projeto

O projeto foi desenvolvido com o objetivo de aprimorar conhecimentos em **Go, Flutter, APIs REST, Clean Architecture e desenvolvimento Full Stack**, aplicando esses conceitos em uma aplicação voltada ao planejamento financeiro.

Além dos conhecimentos técnicos, o projeto também permitiu aplicar conceitos do mercado financeiro, como:

- Juros compostos
- Juros reais
- Inflação
- Rentabilidade
- Aportes mensais
- Evolução patrimonial
- Planejamento financeiro de longo prazo

O projeto busca unir conhecimentos de **desenvolvimento de software e mercado financeiro** através de uma aplicação prática de simulação financeira.
