FROM golang:1.13.4-alpine
LABEL maintainer="dave rowntree"

ENV SOURCES /go/src/github.com/daveo24/cloud-native-go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go
