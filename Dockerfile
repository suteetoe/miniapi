# syntax=docker/dockerfile:1

##
## Builder
##
FROM golang:1.20.2-alpine3.17 AS builder

RUN apk add alpine-sdk
RUN apk add librdkafka=1.9.2-r0
RUN apk add build-base

WORKDIR /go/app

COPY go.mod /go/app
COPY go.sum /go/app
RUN go mod download

ADD . /go/app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -o go-app -tags musl main.go  

##
## Deploy
##
FROM siamecosystems/gl_app-filebeat:1.0
WORKDIR /root/
COPY --from=builder /go/app/go-app .
COPY ./startup.sh /root/startup.sh
RUN chmod +x /root/startup.sh 

# COPY ./tdict-std.txt /root/tdict-std.txt

# ENTRYPOINT ["filebeat", "-c", "/usr/share/filebeat/filebeat.yml"]
CMD ["filebeat", "-c", "/usr/share/filebeat/filebeat.yml"]
ENTRYPOINT /root/go-app
# ENTRYPOINT startup.sh
# CMD ["startup.sh"]