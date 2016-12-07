# Docker image for the Drone HipChat plugin
#
#     cd $GOPATH/src/github.com/jmccann/drone-hipchat
#     make deps build
#     docker build --rm=true -t jmccann/drone-hipchat .

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-hipchat /bin/
ENTRYPOINT ["/bin/drone-hipchat"]
