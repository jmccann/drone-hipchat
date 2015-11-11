package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

// HipChat represents the settings needed to send a HipChat notification.
type HipChat struct {
	Notify bool            `json:"notify"`
	From   string          `json:"from"`
	Room   drone.StringInt `json:"room_id_or_name"`
	Token  string          `json:"auth_token"`
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

	// build the HipChat message
	msg := Message{
		From:    vargs.From,
		Notify:  vargs.Notify,
		Color:   Color(&build),
		Message: BuildMessage(&repo, &build, &system),
	}

	// sends the HipChat message
	if err := client.Send(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildMessage takes a number of drone parameters and builds a message.
func BuildMessage(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	return fmt.Sprintf("%s %s (%s) by %s",
		FirstRuneToUpper(build.Status),
		BuildLink(repo, build, sys),
		build.Branch,
		build.Author,
	)
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

// BuildLink builds the link to a build.
func BuildLink(repo *drone.Repo, build *drone.Build, sys *drone.System) string {
	repoName := repo.Owner + "/" + repo.Name
	url := sys.Link + "/" + repoName + "/" + strconv.Itoa(build.Number)
	return fmt.Sprintf("<a href=\"%s\">%s#%s</a>", url, repoName, build.Commit[:8])
}
