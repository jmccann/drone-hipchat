package main

import (
	"testing"
	"strings"

	"github.com/drone/drone-go/drone"
)

var expected = []string{
	`href="https://beta.drone.io/drone-plugin/drone-hipchat/123"`,
	"Self",
	"1s",
	"Complete",
}

func TestMessage(t *testing.T) {
	system := drone.System{Link: "https://beta.drone.io"}
	repo := drone.Repo{Owner: "drone-plugin", Name: "drone-hipchat"}
	build := drone.Build{Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete"}
	message := BuildMessage(&system, &repo, &build, defaultTemplate)
	t.Log(message)

	for _,text := range expected {
		if !strings.Contains(message, text) {
			t.Error(text + " not present in message")
		}
	}

}
