package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"github.com/drone/drone-go/template"
)

// HipChat represents the settings needed to send a HipChat notification.
type HipChat struct {
	Notify   bool            `json:"notify"`
	From     string          `json:"from"`
	Room     drone.StringInt `json:"room_id_or_name"`
	Token    string          `json:"auth_token"`
	Template string          `json:"template"`
}

func main() {

	// plugin settings
	repo := drone.Repo{}
	build := drone.Build{}
	system := drone.System{}
	vargs := HipChat{}

	// set plugin parameters
	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("system", &system)
	plugin.Param("vargs", &vargs)

	// parse the parameters
	if err := plugin.Parse(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create the HipChat client
	client := NewClient(vargs.Room.String(), vargs.Token)

	// determine notification template
	if len(vargs.Template) == 0 {
		vargs.Template = "<strong>{{ uppercasefirst build.status }}</strong> <a href=\"{{ sys.link }}/{{ repo.owner }}/{{ repo.name }}/{{ build.number }}\">{{ repo.owner }}/{{ repo.name }}#{{ limit build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} </br> - {{ build.message }}"
	}

	// build the HipChat message
	msg := Message{
		From:    vargs.From,
		Notify:  vargs.Notify,
		Color:   Color(&build),
		Message: BuildMessage(&repo, &build, &system, vargs.Template),
	}

	// sends the HipChat message
	if err := client.Send(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildMessage takes a number of drone parameters and builds a message.
func BuildMessage(repo *drone.Repo, build *drone.Build, sys *drone.System, tmpl string) string {

	// data for custom template rendering, if we need it
	payload := &drone.Payload{
		Build:  build,
		Repo:   repo,
		System: sys,
	}

	// render template
	msg, err := template.RenderTrim(tmpl, payload)
	if err != nil {
		return err.Error()
	}

	return msg
}

// Color takes a *plugin.Build object and determines the appropriate
// notification/message color.
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

// FirstRuneToUpper takes a string and capitalizes the first letter.
func FirstRuneToUpper(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}
