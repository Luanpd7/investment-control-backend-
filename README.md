# 💰 Controle de Investimentos — Back-end

API REST desenvolvida em **Go (Golang)** para gerenciamento e acompanhamento de investimentos e patrimônio financeiro.

O back-end é responsável pelo processamento das regras de negócio, gerenciamento dos registros de investimentos e comunicação com o banco de dados PostgreSQL.

## 🚀 Tecnologias

- **Go (Golang):** utilizado no desenvolvimento do back-end, implementando as regras de negócio e o processamento dos dados da aplicação.
- **Gin:** framework utilizado para criação do servidor HTTP e gerenciamento das rotas da API REST.
- **PostgreSQL:** banco de dados relacional utilizado para armazenamento e consulta dos registros de investimentos.
- **REST API:** arquitetura utilizada para comunicação entre o back-end e o front-end através de requisições HTTP.
- **Clean Architecture:** utilizada para organizar o projeto em camadas e separar responsabilidades, facilitando a manutenção e evolução da aplicação.
- **AWS EC2:** serviço utilizado para hospedar e executar a API desenvolvida em Go.
- **AWS RDS:** serviço utilizado para hospedar o banco de dados PostgreSQL.
- **Cloudflare Tunnel:** utilizado para disponibilizar a API através de uma conexão HTTPS segura.
- **Git / GitHub:** utilizados para controle de versão e gerenciamento do código-fonte.
- **JIRA:** Para organização nas demandas.

## 📋 Funcionalidades

- Cadastro de registros de investimentos
- Consulta do patrimônio
- Consulta do histórico de investimentos
- Filtro por ano e mês
- Cálculo da variação do patrimônio
- Consulta por categoria de investimento
- Consulta da evolução anual do patrimônio
- Consulta dos anos disponíveis
- Recuperação do último registro de investimento
- Fornecimento dos dados utilizados pelo dashboard

## 🏗️ Arquitetura

O projeto utiliza **Clean Architecture** para separar as responsabilidades da aplicação e facilitar sua manutenção e evolução.

```text
investment-control-backend/
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

## ☁️ Infraestrutura

A API está hospedada em uma instância **Amazon EC2** e utiliza um banco de dados **PostgreSQL hospedado no Amazon RDS**.

O acesso externo à API é realizado através do **Cloudflare Tunnel**, permitindo que o front-end realize requisições HTTPS para o back-end.

```text
Flutter Web
     │
     │ Dio / HTTPS
     ▼
Cloudflare Tunnel
     │
     ▼
AWS EC2
Go + Gin REST API
     │
     │ PostgreSQL
     │ Porta 5432
     ▼
AWS RDS
PostgreSQL
```

- **Cloudflare Tunnel:** recebe as requisições HTTPS provenientes do front-end e as encaminha para a API.
- **AWS EC2:** hospeda e executa a API REST desenvolvida em Go + Gin.
- **AWS RDS:** hospeda o banco de dados PostgreSQL utilizado para persistência dos registros.
- **Security Groups:** controlam o tráfego de rede entre os recursos da AWS.
- **PostgreSQL (porta 5432):** utilizado na comunicação entre a API executada na EC2 e o banco hospedado no RDS.

## 🔄 Fluxo do Back-end

```text
Requisição HTTP
      │
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
      │
      ▼
 PostgreSQL
  AWS RDS
```

Cada camada possui uma responsabilidade específica:

- **Routes:** define os endpoints disponíveis na API.
- **Handlers:** recebe e trata as requisições HTTP.
- **Use Cases:** contém as regras de negócio da aplicação.
- **Repository:** realiza o acesso e manipulação dos dados.
- **PostgreSQL:** realiza a persistência dos registros.

## 🔌 Endpoints

A API disponibiliza os seguintes endpoints:

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/saveInvestment` | Salva um novo registro de investimento |
| `GET` | `/getAllInvestment` | Retorna os registros de investimentos |
| `GET` | `/dataDashboard` | Retorna os dados utilizados no dashboard |
| `GET` | `/assetGrowth` | Retorna os dados de crescimento do patrimônio |
| `GET` | `/categoryGrowth` | Retorna o crescimento dos investimentos por categoria |
| `GET` | `/availableYears` | Retorna os anos disponíveis nos registros |
| `GET` | `/lastInvestmentRecord` | Retorna o último registro de investimento |

## 🔗 Front-end

O front-end da aplicação foi desenvolvido utilizando **Flutter Web** e está hospedado na **Vercel**.

Repositório:

https://github.com/Luanpd7/investment-control-frontend

## 🎯 Objetivo do Projeto

O projeto foi desenvolvido para aplicar e aprimorar conhecimentos em desenvolvimento **Full Stack e Cloud**, utilizando uma arquitetura separada entre front-end, back-end e banco de dados.

Além do desenvolvimento da API, o projeto permitiu colocar em prática conceitos como:

- Desenvolvimento de APIs REST com Go
- Clean Architecture
- PostgreSQL
- Deploy de aplicações Go
- Amazon EC2
- Amazon RDS
- VPC e Subnets
- Security Groups
- Comunicação entre EC2 e RDS
- Cloudflare Tunnel
- HTTPS
- Integração entre Flutter Web e API REST
