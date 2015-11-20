package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
	"bytes"
	"unicode"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// HipChat represents the settings needed to send a HipChat notification.
type HipChat struct {
	Notify bool                `json:"notify"`
	From   string              `json:"from"`
	Room   drone.StringInt     `json:"room_id_or_name"`
	Token  string              `json:"auth_token"`
	Template map[string]string `json:"template"`
}

const defaultTemplate = "{{.statusFirstRuneUpper}} <a href=\"{{.buildURL}}\">{{.repo.FullName}}#{{.build.Commit}}</a> ({{.build.Branch}}) by {{.build.Author}}"

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

	// build the HipChat message
	msg := Message{
		From:    vargs.From,
		Notify:  vargs.Notify,
		Color:   Color(&build),
		Message: BuildMessage(&repo, &build, &system, &vargs),
	}

	// sends the HipChat message
	if err := client.Send(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildMessage takes a number of drone parameters and builds a message.
func BuildMessage(repo *drone.Repo, build *drone.Build, sys *drone.System, vargs *HipChat) string {
	var buf bytes.Buffer
	message, err := template.New("_").Parse(GetTemplate(build, vargs))
	if err != nil {
		fmt.Printf("Error parsing content template. %s\n", err)
		os.Exit(1)
	}
	message.Execute(&buf, map[string]interface{}{
		"repo": repo,
		"build": build,
		"system": sys,
		"statusShoutingBold": StringToShoutingBold(build.Status),
		"statusFirstRuneUpper": FirstRuneToUpper(build.Status),
		"buildURL": BuildURL(repo, build, sys),
		"buildDuration": BuildDuration(build),
	})
	return buf.String()

}

func GetTemplate(build *drone.Build, vargs *HipChat) string {
	if tmpl, ok := vargs.Template[NormalizeStatus(build)]; ok {
		return tmpl
	} else {
		return defaultTemplate
	}
}

// NormalizeStatus output is success or failure
func NormalizeStatus(build *drone.Build) string {
	if build.Status == drone.StatusSuccess {
		return "success"
	} else {
		return "failure"
	}
}

// Color takes a *plugin.Build object and determines the appropriate
// notification/message color.
func Color(build *drone.Build) string {
	switch build.Status{
	case drone.StatusSuccess:
		return "green"
	case drone.StatusFailure, drone.StatusError, drone.StatusKilled:
		return "red"
	default:
		return "yellow"
	}
}

// StringToShoutingBold transforms a string to bold uppercase
func StringToShoutingBold(s string) string {
	return "<strong>" + strings.ToUpper(s) + "</strong>"
}

// FirstRuneToUpper takes a string and capitalizes the first letter.
func FirstRuneToUpper(s string) string {
	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])
	s = string(a)
	return s
}

func BuildURL(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return sys.Link + "/" + repo.FullName + "/" + strconv.Itoa(build.Number)
}

// BuildToDuration takes a *drone.Build and converts it to a duration
func BuildDuration(build *drone.Build) string {
	durationSeconds := build.Finished - build.Started
	return fmt.Sprintf("%s", time.Duration(durationSeconds) * time.Second)
}
