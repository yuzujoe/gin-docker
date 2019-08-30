#使用golangのversion
FROM golang:1.12.9-alpine3.10



COPY src/api /go/src/api
COPY .env /go/src/api/.env

WORKDIR /go/src/api/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u github.com/joho/godotenv \
  && go get -u github.com/joho/godotenv/cmd/godotenv \
  && dep ensure

#ライブロードの実行
CMD  gin -i run
