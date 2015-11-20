Use the HipChat plugin to send a message to your HipChat room when a build completes.

You will need to supply Drone with a HipChat authentication token. You can learn more about HipChat authentication tokens here: https://www.hipchat.com/docs/apiv2/auth

The following parameters are used to configure the notification:

* **from** - A label to be shown in addition to the sender's name, valid length range: 0 - 25.
* **notify** - Whether this message should trigger a user notification (change the tab color, play a sound, notify mobile phones, etc). Each recipient's notification preferences are taken into account.
Defaults to false.
* **room_id_or_name** - The id or url encoded name of the room, valid length range: 1 - 100.
* **auth_token** - Drone leverages the HipChat API and so it must pass an access token to authenticate correctly. If the token is not provided or invalid you will receive a 401 response.
* **template** - Optional, supply this object to specify custom message templates.

The following is a sample HipChat configuration for your .drone.yml file:

```yaml
notify:
  hipchat:
    from: drone
    notify: true
    room_id_or_name: 1234567
    auth_token: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,
    template:
      success: |
        {{ .Repo.FullName }} successfully completed for {{ .Build.Number }}
      failure: |
        {{ .Repo.FullName }} failed for {{ .Build.Number }}
```

####Available Template Variables
* **Repo** - [Drone Repo Object](https://github.com/drone/drone/blob/master/model/repo.go)
* **Build** - [Drone Build Object](https://github.com/drone/drone/blob/master/model/build.go)
* **System** - [Drone System Object](https://github.com/drone/drone/blob/master/model/system.go)
* **StringToShouting** - Takes in a string and outputs it in uppercase
* **FirstRuneToUpper** - Takes a string and uppercases the first rune
* **BuildURL** - URL to the drone build
