From ubuntu:latest

RUN apt update

RUN apt install -y wget

RUN wget https://github.com/Full-CI-Mode/auto-vesion/releases/download/1.0.0/autover-1.0.0-linux-amd64.tar.gz

RUN  tar xzvf autover-1.0.0-linux-amd64.tar.gz

RUN  cp autover /usr/bin/

CMD [ "/bin/bash" ]