package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// HipChat represents the settings needed to send a HipChat notification.
type HipChat struct {
	Notify   bool            `json:"notify"`
	From     string          `json:"from"`
	Room     drone.StringInt `json:"room_id_or_name"`
	Token    string          `json:"auth_token"`
	Template Template        `json:"template"`
}

// Template represents template options for custom HipChat message
// notifications on success and failure.
type Template struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
	Default string `json:"default"`
}

type TemplateData struct {
	Repo   *drone.Repo
	Build  *drone.Build
	System *drone.System
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
		Message: BuildMessage(&repo, &build, &system, vargs.Template),
	}

	// sends the HipChat message
	if err := client.Send(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildMessage takes a number of drone parameters and builds a message.
func BuildMessage(repo *drone.Repo, build *drone.Build, sys *drone.System, t Template) string {

	// data for custom template rendering, if we need it
	data := struct {
		Repo     *drone.Repo
		Build    *drone.Build
		System   *drone.System
		BuildURL string
	}{repo, build, sys, BuildURL(repo, build, sys)}

	// since notification messages are first based
	// upon build status, we switch on that
	switch build.Status {
	case drone.StatusSuccess:
		return Render(t.Success, t.Default, &data)
	case drone.StatusFailure:
		return Render(t.Failure, t.Default, &data)
	default:
		return Render(t.Default, defaultTemplate, &data)
	}
}

// Render takes a string template and data interface to render the provided
// template to a string.
func Render(tmpl string, customDefault string, data interface{}) string {
	if len(tmpl) == 0 {
		if len(customDefault) > 0 {
			tmpl = customDefault
		} else {
			tmpl = defaultTemplate
		}
	}

	funcMap := template.FuncMap{
		"StringToShouting": StringToShouting,
		"FirstRuneToUpper": FirstRuneToUpper,
	}

	var buf bytes.Buffer
	t, err := template.New("_").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		fmt.Printf("Error parsing content template. %s\n", err)
		os.Exit(1)
	}
	if err := t.Execute(&buf, &data); err != nil {
		fmt.Printf("Error executing content template. %s\n", err)
		os.Exit(1)
	}
	return buf.String()
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

// StringToShoutingBold transforms a string to bold uppercase
func StringToShouting(s string) string {
	return strings.ToUpper(s)
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
