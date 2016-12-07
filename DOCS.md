Use this plugin for sending build status notifications via HipChat. You will
need to supply Drone with a HipChat authentication token. You can learn more
about authentication tokens [here](https://www.hipchat.com/docs/apiv2/auth). You
can override the default configuration with the following parameters:

* `url` - HipChat server URL, defaults to `https://api.hipchat.com`
* `auth_token` - HipChat V2 API token; use a room or user token with the `Send Notification` scope
* `room` - ID or URL encoded name of the room
* `from` - A label to be shown, defaults to `drone`
* `notify` - Whether this message should trigger a user notification (change the
  tab color, play a sound, notify mobile phones, etc). Each recipient's
  notification preferences are taken into account, defaults to false

The following secret values can be set to configure the plugin.

* **HIPCHAT_AUTH_TOKEN** - corresponds to **auth_token**

It is highly recommended to put the **HIPCHAT_AUTH_TOKEN** into secrets so it is
not exposed to users. This can be done using the
[drone-cli](http://readme.drone.io/0.5/reference/cli/overview/).

```bash
drone secret add --image=jmccann/drone-hipchat:0.5 \
    octocat/hello-world HIPCHAT_AUTH_TOKEN mytokenhere
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Example

The following is a sample configuration in your `.drone.yml` file:

```yaml
pipeline:
  hipchat:
    image: jmccann/drone-hipchat:0.5
    room: 1234567
    notify: true
```

### Custom Messages

In some cases you may want to customize the body of the HipChat message you can
use custom templates. For the use case we expose the following additional
parameters:

* `template` - A handlebars template to create a custom payload body. For more
  details take a look at the [docs](http://handlebarsjs.com/).

Example configuration that generate a custom message:

```yaml
pipeline:
  hipchat:
    image: jmccann/drone-hipchat:0.5
    room: 1234567
    from: drone
    notify: true
    template: >
      <strong>{{ uppercasefirst build.status }}</strong> <a href=\"{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}\">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} </br> - {{ build.message }}
```
