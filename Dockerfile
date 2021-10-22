# syntax=docker/dockerfile:1
FROM golang:1.17.1-alpine

ENV AUTH_HOST="0.0.0.0"
ENV AUTH_PORT="5000"
ENV SIGNING_KEY="MYSERCRETKEY"

WORKDIR $GOPATH/src/czwr-mailing-auth/

COPY . .

RUN go mod download

RUN go build -o ./bin/czwr-mailing-auth ./cmd/auth/main.go

EXPOSE 5000

ENTRYPOINT ["./bin/czwr-mailing-auth"]