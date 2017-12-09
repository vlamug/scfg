FROM golang

WORKDIR /migrations

RUN go get -u github.com/pressly/goose/cmd/goose && goose postgres "user=postgres password=root dbname=scfg sslmode=disable" up
