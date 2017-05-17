# Docker image for the Drone HipChat plugin
#
#     docker build -t jmccann/drone-hipchat .

FROM alpine:3.5

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-hipchat /bin/
ENTRYPOINT ["/bin/drone-hipchat"]
