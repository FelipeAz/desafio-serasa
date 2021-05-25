FROM golang:alpine

COPY .env_example .env

# Copia o diretorio atual para o work directory
COPY . /go/src/github.com/FelipeAz/desafio-serasa

# Instalacao do Git e Bash
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/FelipeAz/desafio-serasa/cmd/desafio-serasa

# Copy go mod and sum files
COPY go.mod go.sum /go/src/github.com/FelipeAz/desafio-serasa/

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

## Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#RUN go run main.go

EXPOSE 8080

CMD ["/go/src/github.com/FelipeAz/desafio-serasa/cmd/desafio-serasa/main"]