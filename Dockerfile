FROM golang:1.16.2

ENV GO111MODULE=on

ENV GOLANG_CI_LINT_VERSION=v1.13.2

RUN apt-get update && \
    apt-get install -y graphviz

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}


WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum

EXPOSE 80

RUN go mod download