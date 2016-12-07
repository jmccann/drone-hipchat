package main

import (
	"fmt"
	"os"
	"github.com/drone/drone/model"
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
		Repo   *model.Repo
		Build  *model.Build
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
			&p,
			p.Config.Template),
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	return nil
}

// BuildMessage renders the HipChat message from a template.
func BuildMessage(p *Plugin, tmpl string) string {
	msg, err := RenderTrim(tmpl, p)
	if err != nil {
		return err.Error()
	}

	return msg
}

// Color determins the notfication color based upon the current build status.
func Color(build *model.Build) string {
	switch build.Status {
	case model.StatusSuccess:
		return "green"
	case model.StatusFailure, model.StatusError, model.StatusKilled:
		return "red"
	default:
		return "yellow"
	}
}
