FROM golang:1.23.4-alpine3.20
WORKDIR /app
COPY . .
RUN go build -o stress-test-cli ./cmd/stresstestcli/main.go
ENTRYPOINT ["./stress-test-cli"]