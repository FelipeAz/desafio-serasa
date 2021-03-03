FROM golang:1.6.2-alpine

# Copia o diretorio atual para o work directory
COPY . /go/src/github.com/FelipeAz/desafio-serasa

# Instalacao do Git e Bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/FelipeAz/desafio-serasa

RUN go get gorm.io/gorm
RUN go get gorm.io/driver/mysql
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/garyburd/redigo/redis

# Build the Go app
RUN go run app/main.go

EXPOSE 8080