# drone-hipchat

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-hipchat/status.svg)](http://beta.drone.io/drone-plugins/drone-hipchat)
[![](https://badge.imagelayers.io/plugins/drone-hipchat:latest.svg)](https://imagelayers.io/?images=plugins/drone-hipchat:latest 'Get your own badge on imagelayers.io')

Drone plugin for sending build status notifications via HipChat

## Usage

```sh
./drone-hipchat <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
        "auth_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "room_id_or_name": "1234567",
        "notify": true
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-hipchat .
```

### Example

```sh
docker run -i plugins/drone-hipchat <<EOF
{
    "repo" : {
        "owner": "foo",
        "name": "bar",
        "full_name": "foo/bar"
    },
    "system": {
        "link_url": "http://drone.mycompany.com"
    },
    "build" : {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "commit": "64908ed2414b771554fda6508dd56a0c43766831",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
        "auth_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "room_id_or_name": "1234567",
        "notify": true
    }
}
EOF
```
