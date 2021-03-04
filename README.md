# desafio-serasa
Teste tecnico de codificacao e arquitetura de software proposto pela empresa Serasa.

# Go Directories
A estrutura de pastas segue o layout padrao de projeto do golang onde as pastas sao dividas em:

### `/cmd`
Contem a aplicacao principal do projeto (arquivo main do projeto desafio-serasa) onde se comunica com o codigo presente em `/internal` e `/pkg`.

### `/internal`
Contem codigo e bibliotecas restrita a aplicacao principal (desafio-serasa). Caso o codigo ou a biblioteca possa ser compartilhado entre outras aplicacoes, alocamos em `/pkg`

### `/vendor`
Dependencias da aplicacao geradas pelo `Go Mod`

Referencia: https://github.com/golang-standards/project-layout#cmd

# Instalacao

- Requisitos: Docker

Para a instalacao do projeto basta executar o comando `make run`.

# Arquitetura
A Arquitetura escolhida para este desafio foi a Clean Architecture. Essa arquitetura se baseia em interfaces e contem, para cada entidade, um Controller, um Servico e um Repository. O motivo da escolha dessa arquitetura foi, a facilidade de implementacao de testes, a independencia de uma Interface, de um banco de dados e de outras tecnologias. Essa independencia vem do isolamento dessas ferramentas em arquivos separados, portanto fica facil por exemplo alterar o banco de dados MySQL para um PostgresDB.

O Fluxo da arquitetura segue o padrao:
`
HTTP Request -> Router -> Controller -> Service -> Repository -> JSON Output
`

```
./desafio-serasa/
├── cmd
│   └── main.go
├── config
│   └── infrastructure
│       ├── redis.go
│       ├── router.go
│       └── sqlhandler.go
├── docker
│       ├── json-server
│       │    └── db.json
│       └── Dockerfile.yml
├── internal
│   └── pkg
│       └── app
│           └── database
│           │   └── migrations
│           │       └── database.sql     
│           ├── entity
│           │   ├── access.go
│           │   ├── negativacao.go
│           │   ├── tokendetails.go
│           │   ├── user.go
│           │   └── sqlhandler.go
│           ├── interfaces
│           │   ├── cryptohandler_test.go
│           │   ├── cryptohandler.go
│           │   ├── jwt_auth_test.go
│           │   ├── jwt_auth.go
│           │   ├── mainframe_controller.go
│           │   ├── negativacao_controller.go
│           │   ├── negativacao_repository.go
│           │   ├── redis.go
│           │   ├── router.go
│           │   ├── sqlhandler.go
│           │   ├── user_controller.go
│           │   └── user_repository.go
│           ├── middleware
│           │   └── jwt_middleware.go
│           ├── usecases
│           │   ├── cryptohandler.go
│           │   ├── jwt_auth.go
│           │   ├── mainframe_service.go
│           │   ├── negativacao_service.go
│           │   ├── negativacao_repository.go
│           │   ├── user_repository.go
│           └────── user_service.go
├── scripts
│   ├── redis.go
│   ├── router.go
│   └── sqlhandler.go
├── vendor
│   └── infrastructure
│       ├── redis.go
│       ├── router.go
│       └── sqlhandler.go
├── .env_example
├── docker-compose.yml
├── go.mod
│   └── go.sum
├── makefile
└── README.ms
```

| Camada |Conteudo|
| --- | --- |
| Database | Contem um Arquivo SQL que sera utilizado para a criacao do Banco de Dados. As tabelas serao geradas com auxilio do GORM. |
| Entity | Contem os Modelos que serao utilizados como referencia para criacao da tabela no banco de dados e operacoes de GET, UPDATE, CREATE e DELETE utilizando o GORM. |
| Infrastructure | Contem os Drivers e Frameworks da Aplicacao. |
| Usecases | Contem as Regras de Negocio (Logica da Aplicacao). |
| Interface | Sao responsaveis por transformar Data em Entidade e transformar use cases em um formato 'easy-to-use'. Contem Controllers, Gateways e Presenters |

# Dependencias
Para o funcionamento dessa aplicacao sao necessarios: Curl, JWT-GO, Redis,GORM, GORM MySQL Driver e GIN
O GORM e o GIN pode ser instalado utilizando `go tool` pelos comandos:

`go get -u gorm.io/gorm`
`go get -u gorm.io/driver/mysql`
`go get -u github.com/gin-gonic/gin`
`go get -u github.com/dgrijalva/jwt-go`
`go get -u github.com/garyburd/redigo/redis`

- O banco de dados escolhido para esta aplicacao foi o MySQL 8.0. O usuario e senha do banco de dados eh definido nas variaveis de ambiente. (DB_USER e DB_PASSWORD).
- Criptografia de Dados: A criptografia escolhida foi a AES e utiliza uma variavel de ambiente como chave secreta(ENCRYPT_KEY.)
- GORM: A escolha dessa ferramenta foi pela simplicidade de criacao de tabelas e manuseamento de dados JSON onde sao relacionados com as structs. Essa tecnologia eh utilizada nos repositories (CRUD das entities) e na captura dos dados via json server.
- GIN: O Gin foi uma tecnologia escolhida para manuseamento das requisicoes HTTP, pois facilita a manipulacao das requests,torna o codigo mais legivel e mais curto. Seus usos na aplicacao estao relacionados a criacao de rotas, captura de parametros e no start do servidor.
- JWT: JWT foi escolhido como ferramenta para autenticacao da API. Para a autenticacao, eh necessario criar uma conta na rota /signup, e logar na aplicacao pela rota /login, que gerara um token que deve ser enviado em toda requisicao na forma de Authorization tipo 'Bearer Token'. Esse token tem duracao de 15 minutos, porem eh definido um RefreshToken de 1 semana. Se o usuario deslogar atraves da rota /logout, o token nao sera mais valido, pois o mesmo eh removido do BD pela rota /logout. Os tokens sao gerados a partir de variaveis de ambientes (JWT_SECRET e JWT_REFRESH_SECRET).
- REDIS: Redis foi escolhido para o cacheamento das requisicoes de Negociacoes para aumentar a performance de resposta do sistema.
- JSON-SERVER: O Json Server esta alocado na porta 3000 e recebe o arquivo negativacoes.json (renomeado para mainframe_db.json) recebido na proposta do desafio.

# Tests
Testes Unitarios: Foram testadas as funcoes de Criptohandler e Jwt Auth

- Criptohandler: Testes de Criptografia e Descriptografia do CustomerDocument
- JWT Auth: Testes de Geracao de token de Access & Refresh; Teste de Extracao de Token; Testes de Validacao dos tokens de Access e Refresh.

Para executar os testes basta rodar o comando : `make tests`

# API

A documentacao completa da API se encontra nesse link: https://documenter.getpostman.com/view/7958753/Tz5iA17X
