FROM golang:1.15.1 AS builder
RUN apt-get update
WORKDIR /opt
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM alpine
RUN apk update \
 && apk add --update --no-cache tzdata sshpass openssh openssl
RUN cp /usr/share/zoneinfo/Europe/Istanbul /etc/localtime  \
 && echo "Europe/Istanbul" > /etc/timezone
WORKDIR /app
COPY --from=builder /opt/authentication /app/authentication
COPY --from=builder /opt/configmap /app/configmap
COPY --from=builder /opt/main /app/main

CMD ["/app/main"]
