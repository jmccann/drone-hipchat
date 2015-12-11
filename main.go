package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"github.com/drone/drone-go/template"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone HipChat Plugin built at %s\n", buildDate)

	system := drone.System{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("system", &system)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Template) == 0 {
		vargs.Template = defaultTemplate
	}

	client := NewClient(
		vargs.Room.String(),
		vargs.Token)

	err := client.Send(&Message{
		From:   vargs.From,
		Notify: vargs.Notify,
		Color:  Color(&build),
		Message: BuildMessage(
			&system,
			&repo,
			&build,
			vargs.Template),
	})

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func BuildMessage(system *drone.System, repo *drone.Repo, build *drone.Build, tmpl string) string {
	payload := &drone.Payload{
		System: system,
		Repo:   repo,
		Build:  build,
	}

	msg, err := template.RenderTrim(
		tmpl,
		payload)

	if err != nil {
		return err.Error()
	}

	return msg
}

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
