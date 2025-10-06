FROM golang:1.25.1-alpine3.22

WORKDIR /tests

COPY go.mod go.sum ./
RUN go mod download

COPY tests .

CMD ["go", "test", "-v", "/tests"]
