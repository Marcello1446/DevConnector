FROM golang:1.23.4 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO=0 GOOS=linux go build -o /DevConnector

CMD ["/DevConnector"]