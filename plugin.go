package main

import (
	"net/url"
	"github.com/drone/drone/model"
	"github.com/tbruyelle/hipchat-go/hipchat"
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
	c := hipchat.NewClient(p.Config.AuthToken)
	url, err := url.Parse(p.Config.Url)

	if err != nil {
		return err
	}

	c.BaseURL = url
	notifRq := &hipchat.NotificationRequest{
		Color: Color(p.Build),
		Message: BuildMessage(&p, p.Config.Template),
	}
	_, err = c.Room.Notification(p.Config.Room, notifRq)

	return err
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
func Color(build *model.Build) hipchat.Color {
	switch build.Status {
	case model.StatusSuccess:
		return hipchat.ColorGreen
	case model.StatusFailure, model.StatusError, model.StatusKilled:
		return hipchat.ColorRed
	default:
		return hipchat.ColorYellow
	}
}
