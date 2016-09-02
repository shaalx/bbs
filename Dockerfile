FROM golang

WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/

RUN go get github.com/toukii/bbs
EXPOSE 80

CMD ["/gopath/app/bin/bbs"]