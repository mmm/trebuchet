FROM ubuntu:xenial
MAINTAINER Mark Mims <mark.mims@gmail.com>

RUN apt-get -qq update \
  && apt-get -yqq install \
    stress \
  && apt-get -yqq clean

COPY sorcerer /usr/bin/

ENV SORCERER_PORT=5000
ENV SORCERER_THREADS=1
ENV SORCERER_TIMEOUT=10

CMD ["sorcerer"]
