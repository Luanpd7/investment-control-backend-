# 💰 Controle de investimento — Back-end

API REST desenvolvida em Go para gerenciamento e acompanhamento de
investimentos e patrimônio financeiro.

## 🚀 Tecnologias

- **Go:**  (Golang): utilizado no desenvolvimento do back-end, implementando as regras de negócio e o processamento dos dados da aplicação.
- **Gin:** framework utilizado para criação do servidor HTTP e gerenciamento das rotas da API REST.
- **PostgreSQL:** banco de dados relacional utilizado para armazenamento e consulta dos registros de investimentos.
- **REST API:** arquitetura utilizada para comunicação entre o back-end e o front-end, através de requisições HTTP e endpoints para consulta e manipulação dos dados.
- **Clean Architecture:** utilizada para organizar o projeto em camadas e separar responsabilidades, facilitando a manutenção, evolução e testabilidade da aplicação.
- **Git:** utilizado para controle de versão e gerenciamento do histórico de desenvolvimento do projeto.

## 📋 Funcionalidades

- Cadastro de registros de investimentos
- Consulta do patrimônio
- Consulta do histórico de investimentos
- Filtro por ano e mês
- Cálculo de variação do patrimônio
- Consulta por categoria de investimento

## 🏗️ Arquitetura

O projeto utiliza Clean Architecture para separar as responsabilidades
da aplicação.

```text
financial-independence/
│
├── data/
│   ├── database/
│   │
│   └── repository/     
│
├── domain/
│   ├── entities/
│   │
│   ├── repositories/
│   │
│   └── usecase/
│
├── handlers/
│
└── routes/
```
## 🔌 Endpoints

A API disponibiliza os seguintes endpoints:

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/saveInvestment` | Salva um novo registro de investimento |
| `GET` | `/getAllInvestment` | Retorna todos os registros de investimentos |
| `GET` | `/dataDashboard` | Retorna os dados utilizados no dashboard |
| `GET` | `/assetGrowth` | Retorna os dados de crescimento do patrimônio |
| `GET` | `/categoryGrowth` | Retorna o crescimento dos investimentos por categoria |
| `GET` | `/availableYears` | Retorna os anos disponíveis nos registros |
| `GET` | `/lastInvestmentRecord` | Retorna o último registro de investimento |
