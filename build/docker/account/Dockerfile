FROM golang:1.16 as builder
LABEL maintainer="Huang Kai<h1770360848@outlook.com>"

WORKDIR /app
ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main github.com/Huangkai1008/mall/cmd/account

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/main .
ENTRYPOINT ["./main"]
CMD ["-f", "account"]