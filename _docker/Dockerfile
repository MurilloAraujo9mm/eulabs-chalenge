FROM golang:1.21.4-bullseye

ENV TZ=America/Sao_Paulo
ENV DEBIAN_FRONTEND=noninteractive

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update -yqq && \
  apt-get upgrade -y && \
  apt-get install -y --no-install-recommends \
  git \
  wget \
  curl \
  unzip \
  nano \
  vim \
  zsh

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN groupadd -g 1000 dev && \
  useradd -m -u 1000 -g dev -s $(which zsh) dev

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

WORKDIR /workspace/src

COPY --chown=dev:dev ./src /workspace/src

USER dev:dev

RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 8080

CMD ["CompileDaemon", "--build=go build -o main ./cmd/api", "--command=./main"]
