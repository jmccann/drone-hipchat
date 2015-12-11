# Docker image for the Drone HipChat plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-hipchat
#     make deps build
#     docker build --rm=true -t plugins/drone-hipchat .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-hipchat /bin/
ENTRYPOINT ["/bin/drone-hipchat"]
