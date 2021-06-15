FROM ubuntu:latest

RUN apt-get update

RUN apt install -y wget

RUN wget https://github.com/Full-CI-Mode/auto-vesion/releases/download/1.2.0-alpha/autover-1.2.0-alpha-linux-amd64.tar.gz

RUN  tar xzvf autover-1.2.0-alpha-linux-amd64.tar.gz

RUN  cp autover /usr/bin/

CMD [ "/bin/bash" ]