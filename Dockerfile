FROM  golang:rc-alpine3.13

COPY auto-version.go . 

ENV GO111MODULE=on

RUN apk add gcc

RUN go build auto-version.go

RUN  cp auto-version /usr/bin/

RUN auto-version help

CMD [ "/bin/bash" ]