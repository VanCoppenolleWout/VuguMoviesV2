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

RUN go get -d -v ./...

RUN go get github.com/vugu/vugu/cmd/vugugen

RUN go generate .

RUN GOOS=js GOARCH=wasm go build -o dist/main.wasm .

CMD [ "." ]

#final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=vugumoviesv2 Version=0.0.1
EXPOSE 8844
