FROM golang:1.11-alpine as builder
WORKDIR /usr/build
ADD main.go .
RUN go build -o app .

FROM alpine:latest

WORKDIR /usr/src

COPY --from=builder /usr/build/app .
EXPOSE 8080

CMD ["/usr/src/app"]
