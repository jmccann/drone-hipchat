Use the HipChat plugin to send a message to your HipChat room when a build completes.
You will need to supply Drone with an Incoming Webhook URL. You can create a new
Webhook URL here: https://my.slack.com/services/new/incoming-webhook

The following parameters are used to configuration the notification:

* **webhook_url** - json payloads are sent here
* **room** - messages sent to the above webhook are posted here
* **username** - choose the username this integration will post as

The following is a sample HipChat configuration in your .drone.yml file:

```yaml
notify:
  hipchat:
    webhook_url: https://hooks.slack.com/services/...
    room: dev
    username: drone
```
