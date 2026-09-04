# 💰 Controle de investimento — Back-end

API REST desenvolvida em Go para gerenciamento e acompanhamento de
investimentos e patrimônio financeiro.

## 🚀 Tecnologias

- **Go:**  (Golang): utilizado no desenvolvimento do back-end, implementando as regras de negócio e o processamento dos dados da aplicação.
- **Gin:** framework utilizado para criação do servidor HTTP e gerenciamento das rotas da API REST.
- **PostgreSQL:** banco de dados relacional utilizado para armazenamento e consulta dos registros de investimentos.
- **REST API:** arquitetura utilizada para comunicação entre o back-end e o front-end, através de requisições HTTP e endpoints para consulta e manipulação dos dados.
- **Clean Architecture:** utilizada para organizar o projeto em camadas e separar responsabilidades, facilitando a manutenção, evolução e testabilidade da aplicação.
- **AWS EC2:** serviço utilizado para hospedar e executar a API desenvolvida em Go.
- **AWS RDS:** serviço utilizado para hospedar o banco de dados PostgreSQL.
- **Cloudflare Tunnel:** utilizado para disponibilizar a API através de uma conexão HTTPS segura.
- **Git:** utilizado para controle de versão e gerenciamento do histórico de desenvolvimento do projeto.

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

## ☁️ Fluxo da comunicação

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

## 🎯 Objetivo do Projeto

O projeto foi desenvolvido com o objetivo de aplicar e aprimorar conhecimentos em desenvolvimento **Full Stack e Cloud**, integrando front-end, back-end, banco de dados e infraestrutura em nuvem.

Além do desenvolvimento da aplicação, o projeto também teve como objetivo colocar em prática conhecimentos de **AWS**, realizando o deploy e a configuração da infraestrutura necessária para executar a aplicação em ambiente de nuvem.

Durante o desenvolvimento e deploy foram aplicados conceitos como:

- Desenvolvimento de APIs REST com **Go + Gin**
- **Clean Architecture**
- **PostgreSQL**
- Deploy de aplicações Go em **Amazon EC2**
- Hospedagem do PostgreSQL no **Amazon RDS**
- Configuração de **VPC**
- Utilização de **subnets públicas e privadas**
- Configuração de **Security Groups**
- Comunicação entre **EC2 e RDS**
- Configuração de regras de entrada e saída de rede
- Utilização de **Internet Gateway e Route Tables**
- Acesso e gerenciamento de instâncias EC2 via **SSH**
- Configuração de variáveis de ambiente no servidor
- Utilização do **Cloudflare Tunnel** para disponibilização da API via HTTPS
- Integração entre **Flutter Web, API REST e infraestrutura AWS**

Dessa forma, o projeto também serviu como ambiente prático para compreender como uma aplicação pode ser **implantada, configurada e executada na AWS**, abrangendo conceitos de infraestrutura, rede, segurança e computação em nuvem.
