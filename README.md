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
│   ├── access.go
│   ├── negativacao.go
│   ├── tokendetails.go
│   ├── user.go
├── infrastructure
│   ├── router.go
│   └── sqlhandler.go
├── interfaces
│   ├── cryptohandler.go
│   ├── jwt_auth.go
│   ├── mainframe_controller.go
│   ├── negativacao_controller.go
│   ├── negativacao_repository.go
│   ├── router.go
│   ├── sqlhandler.go
│   ├── user_controller.go
│   ├── user_repository.go
├── middleware
│   ├── jwt_middleware.go
├── usecases
│   ├── cryptohandler.go
│   ├── jwt_auth.go
│   ├── mainframe_service.go
│   ├── negativacao_service.go
│   ├── negativacao_repository.go
│   ├── user_repository.go
│   ├── user_service.go
├── .env_example
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

- O banco de dados escolhido para esta aplicacao foi o MySQL 8.0. O usuario e senha do banco de dados eh definido nas variaveis de ambiente. (DB_USER e DB_PASSWORD).
- Criptografia de Dados: A criptografia escolhida foi a AES e utiliza uma variavel de ambiente como chave secreta(ENCRYPT_KEY.)
- GORM: A escolha dessa ferramenta foi pela simplicidade de criacao de tabelas e manuseamento de dados JSON onde sao relacionados com as structs. Essa tecnologia eh utilizada nos repositories (CRUD das entities) e na captura dos dados via json server.
- GIN: O Gin foi uma tecnologia escolhida para manuseamento das requisicoes HTTP, pois facilita a manipulacao das requests,torna o codigo mais legivel e mais curto. Seus usos na aplicacao estao relacionados a criacao de rotas, captura de parametros e no start do servidor.
- JWT: JWT foi escolhido como ferramenta para autenticacao da API. Para a autenticacao, eh necessario criar uma conta na rota /signup, e logar na aplicacao pela rota /login, que gerara um token que deve ser enviado em toda requisicao na forma de Authorization tipo 'Bearer Token'. Esse token tem duracao de 15 minutos, porem eh definido um RefreshToken de 1 semana. Se o usuario deslogar atraves da rota /logout, o token nao sera mais valido, pois o mesmo eh removido do BD pela rota /logout. Os tokens sao gerados a partir de variaveis de ambientes (JWT_SECRET e JWT_REFRESH_SECRET).

# API

# Tests