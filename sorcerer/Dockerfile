FROM golang:1.12-stretch

RUN apt-get -qq update \
  && apt-get -yqq install \
    stress \
  && apt-get -yqq clean

ADD . /go/src/app
WORKDIR /go/src/app

ENV SORCERER_PORT=5000
ENV SORCERER_THREADS=1
ENV SORCERER_TIMEOUT=10

CMD ["go", "run", "main.go"]
