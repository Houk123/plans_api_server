FROM golang:1.21-alpine as builder

WORKDIR /server
RUN go build -o plans-rest-api-server ./cmd/plans-rest-api-server

ENV CONFIG_PATH=/server/local.yml

CMD ["./plans-rest-api-server"]