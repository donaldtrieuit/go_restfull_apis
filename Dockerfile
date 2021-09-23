FROM golang:1.14.2-alpine

WORKDIR /go/src/app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download

COPY . .
RUN go build -v -o app main.go

########
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
RUN echo "Asia/Ho_Chi_Minh" > /etc/timezone

WORKDIR /go/src/app

COPY --from=0 /go/src/app/config ./config
COPY --from=0 /go/src/app/app .
CMD ["./app"]