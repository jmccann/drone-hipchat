Use the HipChat plugin to send a message to your HipChat room when a build completes.
You will need to supply Drone with a HipChat authentication token.  
You can learn more here: https://www.hipchat.com/docs/apiv2/auth

The following parameters are used to configure the notification:

* **room_id_or_name** - The id or url encoded name of the room, valid length range: 1 - 100.
* **from** - A label to be shown in addition to the sender's name, valid length range: 0 - 25.
* **notify** - Whether this message should trigger a user notification (change the tab color, play a sound, notify mobile phones, etc). Each recipient's notification preferences are taken into account.
Defaults to false.

The following is a sample HipChat configuration for your .drone.yml file:

```yaml
notify:
  hipchat:
    from: drone
    notify: true
    room_id_or_name: 1234567
    auth_token: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```
