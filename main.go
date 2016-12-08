package main

import (
	"os"

	"github.com/drone/drone/model"
	"github.com/joho/godotenv"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

<<<<<<< HEAD
var (
	buildCommit     string
	defaultTemplate = `<strong>{{ uppercasefirst build.status }}</strong> <a href="{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} </br> - {{ build.message }}`
	defaultTitleTemplate = `by {{ build.author }} in {{ duration build.started_at build.finished_at }}`
	defaultDescTemplate = `<a href="{{ build.link_url }}">{{ truncate build.commit 8 }}</a> - <i>{{ build.message }}</i>`
	defaultActivityTemplate = `<a href="{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}"><strong>{{ build.status }}</strong> {{ repo.name }} ({{ build.branch }})</a>`
	defaultIcon = "http://readme.drone.io/logos/downstream.svg"
)
=======
var revision string // build number set at compile-time
>>>>>>> upstream/master

func main() {
	app := cli.NewApp()
	app.Name = "hipchat plugin"
	app.Usage = "hipchat plugin"
	app.Action = run
	app.Version = revision
	app.Flags = []cli.Flag{

		//
		// plugin args
		//

		cli.StringFlag{
			Name: "url",
			Usage: "HipChat server URL",
			Value: "https://api.hipchat.com",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name: "auth_token",
			Usage: "HipChat server URL",
			EnvVar: "HIPCHAT_AUTH_TOKEN,PLUGIN_AUTH_TOKEN",
		},
		cli.StringFlag{
			Name: "room",
			Usage: "ID or URL encoded name of the room",
			EnvVar: "PLUGIN_ROOM",
		},
		cli.StringFlag{
			Name: "from",
			Usage: "A label to be shown",
			Value: "drone",
			EnvVar: "PLUGIN_FROM",
		},
		cli.BoolFlag{
			Name: "notify",
			Usage: "Whether this message should trigger a user notification (change the tab color, play a sound, notify mobile phones, etc). Each recipient's notification preferences are taken into account",
			EnvVar: "PLUGIN_NOTIFY",
		},
		cli.StringFlag{
			Name: "template",
			Usage: "A handlebars template to create a custom payload body.",
			Value: "<strong>{{ uppercasefirst build.status }}</strong> <a href=\"{{ build.link }}\">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.created build.finished }} </br> - {{ build.message }}",
			EnvVar: "PLUGIN_TEMPLATE",
		},

		//
		// repo args
		//

		cli.StringFlag{
			Name:   "repo.fullname",
			Usage:  "repository full name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "repository link",
			EnvVar: "DRONE_REPO_LINK",
		},
		cli.StringFlag{
			Name:   "repo.avatar",
			Usage:  "repository avatar",
			EnvVar: "DRONE_REPO_AVATAR",
		},
		cli.StringFlag{
			Name:   "repo.branch",
			Usage:  "repository default branch",
			EnvVar: "DRONE_REPO_BRANCH",
		},
		cli.BoolFlag{
			Name:   "repo.private",
			Usage:  "repository is private",
			EnvVar: "DRONE_REPO_PRIVATE",
		},
		cli.BoolFlag{
			Name:   "repo.trusted",
			Usage:  "repository is trusted",
			EnvVar: "DRONE_REPO_TRUSTED",
		},

		//
		// commit args
		//

		cli.StringFlag{
			Name:   "remote.url",
			Usage:  "git remote url",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.link",
			Usage:  "git commit link",
			EnvVar: "DRONE_COMMIT_LINK",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},

		//
		// build args
		//

		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.IntFlag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.IntFlag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "build.deploy",
			Usage:  "build deployment target",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.BoolFlag{
			Name:   "yaml.verified",
			Usage:  "build yaml is verified",
			EnvVar: "DRONE_YAML_VERIFIED",
		},
		cli.BoolFlag{
			Name:   "yaml.signed",
			Usage:  "build yaml is signed",
			EnvVar: "DRONE_YAML_SIGNED",
		},

		cli.StringFlag{
			Name:	"env-file",
			Usage: "source env file",
		},
	}

<<<<<<< HEAD
	message := &Message{
		From:    vargs.From,
		Notify:  vargs.Notify,
		Color:   Color(&build),
		Message: BuildTemplate(
			&system,
			&repo,
			&build,
			vargs.Template,
		),
	}

	if vargs.UseCard {
		message.Card = BuildCard(
			&system,
			&repo,
			&build,
			&vargs,
		)
	}

	client := NewClient(
		vargs.URL,
		vargs.Room.String(),
		vargs.Token,
	)

	if err := client.Send(message); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}

// BuildCard creates the HipChat card
func BuildCard(system *drone.System, repo *drone.Repo, build *drone.Build, vargs *Params) *Card {
	if len(vargs.TitleTemplate) == 0 {
		vargs.TitleTemplate = defaultTitleTemplate
	}

	if len(vargs.Icon) == 0 {
		vargs.Icon = defaultIcon
	}

	if len(vargs.DescTemplate) == 0 {
		vargs.DescTemplate = defaultDescTemplate
	}

	if len(vargs.ActivityTemplate) == 0 {
		vargs.ActivityTemplate = defaultActivityTemplate
	}

	card := &Card{
		ID:     build.Commit,
		Style:  "application",
		Format: "medium",
		Title: BuildTemplate(
			system,
			repo,
			build,
			vargs.TitleTemplate,
		),
		URL: BuildTemplate(
			system,
			repo,
			build,
			"{{ system.link_url }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}",
		),
		Activity: Activity{
			Icon: vargs.Icon,
			HTML: BuildTemplate(
				system,
				repo,
				build,
				vargs.ActivityTemplate,
			),
		},
	}

	if len(build.Avatar) > 0 {
		card.Icon = &build.Avatar
	}

	if len(vargs.DescTemplate) > 0 {
		card.Description = &Description{
			Format: "html",
			Value: BuildTemplate(
				system,
				repo,
				build,
				vargs.DescTemplate,
			),
		}
	}

	return card
}
// BuildMessage renders the HipChat message from a template.
func BuildTemplate(system *drone.System, repo *drone.Repo, build *drone.Build, tmpl string) string {
=======
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	logrus.WithFields(logrus.Fields{
		"Revision": revision,
	}).Info("Drone Hipchat Plugin Version")
>>>>>>> upstream/master

	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
    Repo: &model.Repo{
			Owner:     c.String("repo.owner"),
			Name:      c.String("repo.name"),
			Link:      c.String("repo.link"),
			Avatar:    c.String("repo.avatar"),
			Branch:    c.String("repo.branch"),
			IsPrivate: c.Bool("repo.private"),
			IsTrusted: c.Bool("repo.trusted"),
		},

		Build: &model.Build{
			Number:    c.Int("build.number"),
			Event:     c.String("build.event"),
			Status:    c.String("build.status"),
			Enqueued:  0,
			Created:   int64(c.Int("build.created")),
			Started:   int64(c.Int("build.started")),
			Finished:  int64(c.Int("build.finished")),
			Deploy:    c.String("build.deploy"),
			Commit:    c.String("commit.sha"),
			Branch:    c.String("commit.branch"),
			Ref:       c.String("commit.sha"),
			Refspec:   "",
			Remote:    c.String("remote.url"),
			Title:     "",
			Message:   c.String("commit.message"),
			Timestamp: 0,
			Author:    c.String("commit.author.name"),
			Avatar:    c.String("commit.author.avatar"),
			Email:     c.String("commit.author.email"),
			Link:      c.String("build.link"),
		},

		Config: Config{
			Url:       c.String("url"),
			AuthToken: c.String("auth_token"),
			Room:      c.String("room"),
			From:      c.String("from"),
			Notify:    c.Bool("notify"),
			Template:  c.String("template"),
		},
	}

	return plugin.Exec()
}
