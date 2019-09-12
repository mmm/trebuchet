FROM golang:1.12-stretch

RUN apt-get -qq update \
  && apt-get -yqq install \
    stress \
  && apt-get -yqq clean

RUN go get github.com/pkg/errors \
  && go get golang.org/x/net/context \
  && go get google.golang.org/grpc

ENV SORCERER_HOST=localhost
ENV SORCERER_PORT=8080
ENV SORCERER_THREADS=1
ENV SORCERER_TIMEOUT=10

ADD . /go/src/app
WORKDIR /go/src/app

CMD ["go", "run", "."]
