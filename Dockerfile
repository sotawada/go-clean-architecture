# Base image
FROM golang:1.23

WORKDIR /app

RUN go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

CMD ["air", "-c", "air.toml"]
