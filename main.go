package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-plugin-go/plugin"
)

type HipChat struct {
	Notify bool   `json:"notify"`
	From   string `json:"from"`
	Room   string `json:"room_id_or_name"`
	Token  string `json:"auth_token"`
}

func main() {
	repo := plugin.Repo{}
	build := plugin.Build{}
	system := plugin.System{}
	vargs := HipChat{}

	plugin.Param("build", &build)
	plugin.Param("repo", &repo)
	plugin.Param("system", &system)
	plugin.Param("vargs", &vargs)

	// parse the parameters
	if err := plugin.Parse(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create the HipChat client
	client := NewClient(vargs.Room, vargs.Token)

	// generate the HipChat message
	msg := Message{
		From:    vargs.From,
		Notify:  vargs.Notify,
		Color:   Color(&build),
		Message: BuildMessage(&repo, &build, &system),
	}

	// sends the message
	if err := client.Send(&msg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func BuildMessage(repo *plugin.Repo, build *plugin.Build, sys *plugin.System) string {
	return fmt.Sprintf("*%s* <%s|%s/%s#%s> (%s) by %s",
		build.Status,
		fmt.Sprintf("%s/%s/%v", sys.Link, repo.FullName, build.Number),
		repo.Owner,
		repo.Name,
		build.Commit[:8],
		build.Branch,
		build.Author,
	)
}

func Color(build *plugin.Build) string {
	switch build.Status {
	case plugin.StateSuccess:
		return "green"
	case plugin.StateFailure, plugin.StateError, plugin.StateKilled:
		return "red"
	default:
		return "yellow"
	}
}
