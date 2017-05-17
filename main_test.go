package main

import (
	"testing"

	"github.com/drone/drone/model"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

func TestBuildMessage(t *testing.T) {
	tests := []struct {
		plugin  *Plugin
		tmpl    string
		message string
	}{
		{
			plugin: &Plugin {
				Repo:   &model.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
				Build:  &model.Build{Commit: "1234567890", Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete", Link: "https://beta.drone.io/drone-plugin/drone-hipchat/123"},
			},
			tmpl:   "<strong>{{ uppercasefirst build.Status }}</strong> <a href=\"{{ build.Link }}\">{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}</a> ({{ build.branch }}) by {{ build.author }} in {{ duration build.Started build.Finished }} </br> - {{ build.message }}",
			message: `<strong>Success</strong> <a href="https://beta.drone.io/drone-plugin/drone-hipchat/123">drone-plugin/drone-hipchat#12345678</a> (master) by Self in 1s
 </br> - Complete`,
		},
		{
			plugin: &Plugin {
				Repo:   &model.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
				Build:  &model.Build{Commit: "1234567890", Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete", Link: "https://beta.drone.io/drone-plugin/drone-hipchat/123"},
			},
			tmpl:   `{{ uppercasefirst build.status }} {{ build.Link }} {{ repo.owner }} {{ repo.name }} {{ build.number }} {{ repo.owner }} {{ repo.name }} {{ truncate build.commit 8 }} {{ build.branch }} {{ build.author }} {{ duration build.Started build.Finished }} {{ build.message }}`,
			message: `Success https://beta.drone.io/drone-plugin/drone-hipchat/123 drone-plugin drone-hipchat 123 drone-plugin drone-hipchat 12345678 master Self 1s
 Complete`,
		},
		{
			plugin: &Plugin {
				Repo:   &model.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
				Build:  &model.Build{Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete", Link: "https://beta.drone.io/drone-plugin/drone-hipchat/123"},
			},
			tmpl:   "{{ }",
			message: `Parse error on line 1:
Lexer error
Token: Error{"Unexpected character in expression: '}'"}`,
		},
	}

	for _, test := range tests {
		message := BuildMessage(test.plugin, test.tmpl)
		if test.message != message {
			t.Errorf("expected message:\n %s \n got message: \n %s \n", test.message, message)
		}
	}
}

func TestColor(t *testing.T) {
	tests := []struct {
		build *model.Build
		color hipchat.Color
	}{
		{build: &model.Build{Status: model.StatusSkipped}, color: hipchat.ColorYellow},
		{build: &model.Build{Status: model.StatusPending}, color: hipchat.ColorYellow},
		{build: &model.Build{Status: model.StatusRunning}, color: hipchat.ColorYellow},
		{build: &model.Build{Status: model.StatusSuccess}, color: hipchat.ColorGreen},
		{build: &model.Build{Status: model.StatusFailure}, color: hipchat.ColorRed},
		{build: &model.Build{Status: model.StatusKilled}, color: hipchat.ColorRed},
		{build: &model.Build{Status: model.StatusError}, color: hipchat.ColorRed},
		{build: &model.Build{Status: "foobar"}, color: hipchat.ColorYellow},
	}

	for _, test := range tests {
		if Color(test.build) != test.color {
			t.Errorf("expected %s got %s", test.color, Color(test.build))
		}
	}
}
