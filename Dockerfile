FROM golang:1.11.2

# git installation
RUN apt -y update \
    && apt -y remove git \
    && apt -y install \
        libssl-dev \
        libghc-zlib-dev \
        libcurl4-gnutls-dev \
        libexpat1-dev \
        gettext \
        unzip
WORKDIR /usr/src
RUN wget https://github.com/git/git/archive/v2.20.0.tar.gz -O git.tar.gz \
    && tar -xf git.tar.gz \
    && cd git-* \
    && make prefix=/usr/local all \
    && make prefix=/usr/local install \
    && git --version \
    && apt -y remove \
        libssl-dev \
        libghc-zlib-dev \
        libcurl4-gnutls-dev \
        libexpat1-dev \
        gettext \
        unzip

WORKDIR /go/src/20181209sun-go-jwt/server
COPY . .
ENV GO111MODULE=on

RUN go mod vendor \
    && go mod download \
    && go get github.com/pilu/fresh
CMD ["fresh"]

EXPOSE 9000