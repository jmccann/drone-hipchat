package main

import (
	"fmt"
	"os"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/template"
)

type (
	Config struct {
		Url       string
		AuthToken string
		Room      string
		From      string
		Notify    bool
		Template  string
	}

	Plugin struct {
		Repo   *drone.Repo
		Build  *drone.Build
		Config Config
	}
)

func (p Plugin) Exec() error {
	client := NewClient(
		p.Config.Url,
		p.Config.Room,
		p.Config.AuthToken,
	)

	if err := client.Send(&Message{
		From:   p.Config.From,
		Notify: p.Config.Notify,
		Color:  Color(p.Build),
		Message: BuildMessage(
			p.Repo,
			p.Build,
			p.Config.Template),
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	return nil
}

// BuildMessage renders the HipChat message from a template.
func BuildMessage(repo *drone.Repo, build *drone.Build, tmpl string) string {

	payload := &drone.Payload{
		Repo:   repo,
		Build:  build,
	}

	msg, err := template.RenderTrim(tmpl, payload)
	if err != nil {
		return err.Error()
	}

	return msg
}

// Color determins the notfication color based upon the current build status.
func Color(build *drone.Build) string {
	switch build.Status {
	case drone.StatusSuccess:
		return "green"
	case drone.StatusFailure, drone.StatusError, drone.StatusKilled:
		return "red"
	default:
		return "yellow"
	}
}
