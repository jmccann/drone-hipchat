# drone-hipchat
Drone plugin for sending HipChat notifications


## Overview

This plugin is responsible for sending build notifications to your HipChat room:

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
        "commit": "9f2849d5",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
      "notify": true,
      "from": "drone",
      "room_id_or_name": "1234567",
      "auth_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
      "template": {
        "success": "{{ .repo.FullName }} successfully completed for {{ .build.Number }}",
        "failure": "{{ .repo.FullName }} failed for {{ .build.Number }}"
      }
    }
}
EOF
```

## Docker

Build the Docker container. Note that we need to use the `-netgo` tag so that
the binary is built without a CGO dependency:

```sh
CGO_ENABLED=0 go build -a -tags netgo
docker build --rm=true -t plugins/drone-hipchat .
```

Send a HipChat notification:

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
        "commit": "9f2849d5",
        "branch": "master",
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
    },
    "vargs": {
        "notify": true,
        "from": "drone",
        "room_id_or_name": "1234567",
        "auth_token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        "template": {
          "success": "{{ .repo.FullName }} successfully completed for {{ .build.Number }}",
          "failure": "{{ .repo.FullName }} failed for {{ .build.Number }}"
        }
    }
}
EOF
```
