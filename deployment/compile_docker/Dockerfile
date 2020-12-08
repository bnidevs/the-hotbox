FROM ubuntu

WORKDIR /

RUN apt-get update
RUN apt-get -y upgrade

RUN apt-get install -y wget
RUN apt-get install -y tar

RUN mkdir temp

WORKDIR /temp

RUN wget https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz

RUN tar -xvf go1.15.6.linux-amd64.tar.gz

RUN mv go /usr/local

ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# check go version (comment out if there are issues installing go)

# RUN go version

WORKDIR /

RUN apt-get install -y git
RUN apt-get install -y make

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get install -y --no-install-recommends cmake

RUN apt-get install -y sudo

RUN go get -u -d gocv.io/x/gocv

WORKDIR $GOPATH/src/gocv.io/x/gocv

RUN make install

WORKDIR /

RUN mkdir docs

WORKDIR /docs

COPY main.go .

RUN go get -u github.com/aws/aws-lambda-go/lambda

RUN go build main.go

CMD ./main
