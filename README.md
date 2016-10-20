# drone-hipchat

[![Build Status](http://beta.drone.io/api/badges/jmccann/drone-hipchat/status.svg)](http://beta.drone.io/jmccann/drone-hipchat)

Drone plugin to send build status notifications via HipChat. For the usage information and
a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t jmccann/drone-hipchat .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-hipchat' not found or does not exist.
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  jmccann/drone-hipchat --notify --from JohnDoe \
  --room MyHipchatRoom --auth_token xxxxxxxxxxx
```
