# Build Stage
FROM junjuew/mkt-buildimage:1.11 AS build-stage

LABEL app="build-mkt"
LABEL REPO="https://github.com/junjuew/mkt"

ENV PROJPATH=/go/src/github.com/junjuew/mkt

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/junjuew/mkt
WORKDIR /go/src/github.com/junjuew/mkt

RUN make build-alpine

# Final Stage
FROM junjuew/mkt:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/junjuew/mkt"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/mkt/bin

WORKDIR /opt/mkt/bin

COPY --from=build-stage /go/src/github.com/junjuew/mkt/bin/mkt /opt/mkt/bin/
RUN chmod +x /opt/mkt/bin/mkt

# Create appuser
RUN adduser -D -g '' mkt
USER mkt

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/mkt/bin/mkt"]
