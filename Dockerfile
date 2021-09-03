FROM golang:1.17.0-alpine3.13 as aprendaApp
WORKDIR /go/src/aprendamais
COPY ./src .
RUN go get -d -v ./...
RUN go install ./...
CMD ["aprendamais"]

FROM mysql:8.0 as aprendaDb
WORKDIR
COPY ./sql .

FROM node:16-alpine3.11 as aprendaTest
WORKDIR /testes
COPY testes/__tests__ .
RUN apk update && npm i


FROM node:16-alpine3.11 as aprendaFront
WORKDIR /front
COPY frontend .
RUN apk update && npm i