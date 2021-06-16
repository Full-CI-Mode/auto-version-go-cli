FROM  golang:rc-alpine3.13

COPY . . 

ENV GO111MODULE=on

RUN apk add gcc


RUN go build auto-version.go


RUN ls

RUN  cp auto-version /usr/bin/

RUN auto-version help

CMD [ "/bin/bash" ]