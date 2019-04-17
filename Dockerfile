FROM golang:1.11-alpine

RUN apk update
RUN apk add git ca-certificates
ADD main.go main.go
RUN go build main.go
EXPOSE 9000
CMD ["./main"]

