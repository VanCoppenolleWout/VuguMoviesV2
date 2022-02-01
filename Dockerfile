#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /dist

COPY /dist/wasm_exec.js ./
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN go get github.com/vugu/vugu/cmd/vugugen

RUN go build -o bin/server ./cmd/server

RUN go generate ./ui

RUN GOARCH=wasm GOOS=js go build -o /index.wasm ./wasm/index

RUN go get -u github.com/vugu/vgrun
RUN vgrun -install-tools

EXPOSE 8844

