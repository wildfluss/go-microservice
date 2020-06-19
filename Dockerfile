FROM golang:1.14.4-alpine as builder
# $GOPATH/go.mod exists but should not
WORKDIR /a/microservice
COPY main.go .
COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on 
RUN go install 

FROM alpine:3.10
CMD ["./microservice","--host","0.0.0.0"]
COPY --from=builder /go/bin/microservice .
