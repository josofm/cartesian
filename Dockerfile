FROM golang:alpine as build-env

ENV GO111MODULE=on
ENV GOLANG_CI_LINT_VERSION=v1.13.2

WORKDIR /app
ADD . /app

RUN cd /usr && \
    wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s ${GOLANG_CI_LINT_VERSION}

RUN cd /app/cmd/cartesian && go build cartesian.go


FROM alpine

RUN apk --no-cache update && \
    apk --no-cache add ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=build-env /app/cmd/cartesian/cartesian /app

EXPOSE 80

ENTRYPOINT ["/app/cartesian"]