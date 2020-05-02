FROM golang:latest

ARG app_env
ENV APP_ENV $app_env
ENV CGO_ENABLED=0

RUN go get github.com/githubnemo/CompileDaemon

COPY . /go/src/app
WORKDIR /go/src/app

# Build the Go app
RUN go build

EXPOSE 3001

# Entrypoint using CompileDaemon for live reload
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build" -command="./main"