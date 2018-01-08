FROM golang:1.9

USER root
ENV TZ Asia/Tokyo
ENV GOPATH /go
ENV APPROOT ${GOPATH}/src/github.com/fujimisakari/go-crawler
ENV APPENV production

RUN apt-get update
RUN apt-get install -y mysql-client vim net-tools telnet curl

WORKDIR ${APPROOT}

COPY . ${APPROOT}

RUN set -x \
   && curl -L https://github.com/mattes/migrate/releases/download/v3.0.1/migrate.linux-amd64.tar.gz | tar xvz \
   && mv migrate.linux-amd64 /go/bin/migrate \
   && go get -u github.com/golang/dep/cmd/dep \
   && dep ensure \
   && go build -o bin/go-crawler

CMD bin/go-crawler server
