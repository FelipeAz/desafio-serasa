# desafio-serasa
Teste tecnico de codificacao e arquitetura de software proposto pela empresa Serasa.

# Instalacao

# Arquitetura
A Arquitetura escolhida para este desafio foi a Clean Architecture. Essa arquitetura se baseia em interfaces e contem, para cada entidade, um Controller, um Servico e um Repository alem da Entity. 

O Fluxo da arquitetura segue o padrao:

`
HTTP Request -> Router -> Controller -> Service -> Repository -> JSON Output
`

```
./app/
├── database
│   ├── migrations
│       └── database.sql
├── entity
│   ├── negativacao.go
├── infrastructure
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── negativacao_controller.go
│   ├── negativacao_repository.go
│   ├── router.go
│   ├── sqlhandler.go
├── usecases
│   ├── negativacao_service
│   ├── negativacao_repository.go
└── main.go
```

| Camada |Conteudo|
| --- | --- |
| Database | Contem um Arquivo SQL que sera utilizado para a criacao do Banco de Dados. As tabelas serao geradas com auxilio do GORM. |
| Entity | Contem os Modelos que serao utilizados como referencia para criacao da tabela no banco de dados e operacoes de GET, UPDATE, CREATE e DELETE utilizando o GORM. |
| Infrastructure | Contem os Drivers da Aplicacao. |
| Usecases | Contem as Regras de Negocio (Logica da Aplicacao). |

# Dependencias
Para o funcionamento dessa aplicacao sao necessarios: Curl, JWT-GO, GORM, GORM MySQL Driver e GIN
O GORM e o GIN pode ser instalado utilizando `go tool` pelos comandos:

`go get -u gorm.io/gorm`
`go get -u gorm.io/driver/mysql`
`go get -u github.com/gin-gonic/gin`
`go get -u github.com/dgrijalva/jwt-go`

O banco de dados escolhido para esta aplicacao foi o MySQL 8.0

# API

# Tests