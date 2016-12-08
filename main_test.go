package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone-go/drone"
)

func TestClient(t *testing.T) {
	hipchat := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "text/html")
	}))
	defer hipchat.Close()

	client := NewClient(
		hipchat.URL,
		"TheIronThrone",
		"xyz",
	)

	if err := client.Send(&Message{
		From:    "John Snow",
		Color:   "Green",
		Notify:  true,
		Message: "Testy test",
	}); err != nil {
		t.Error(err)
	}
}

func TestBuildTemplate(t *testing.T) {
	tests := []struct {
		system  *drone.System
		repo    *drone.Repo
		build   *drone.Build
		tmpl    string
		message string
	}{
		{
			system: &drone.System{Link: "https://beta.drone.io"},
			repo:   &drone.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
			build:  &drone.Build{Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete"},
			tmpl:   defaultTemplate,
			message: `<strong>Success</strong> <a href="https://beta.drone.io/drone-plugin/drone-hipchat/123">drone-plugin/drone-hipchat#</a> (master) by Self in 1s
 </br> - Complete`,
		},
		{
			system: &drone.System{Link: "https://beta.drone.io"},
			repo:   &drone.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
			build:  &drone.Build{Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete"},
			tmpl:   `{{ uppercasefirst build.status }} {{ system.link_url }} {{ repo.owner }} {{ repo.name }} {{ build.number }} {{ repo.owner }} {{ repo.name }} {{ truncate build.commit 8 }} {{ build.branch }} {{ build.author }} {{ duration build.started_at build.finished_at }} {{ build.message }}`,
			message: `Success https://beta.drone.io drone-plugin drone-hipchat 123 drone-plugin drone-hipchat  master Self 1s
 Complete`,
		},
		{
			system: &drone.System{Link: "https://beta.drone.io"},
			repo:   &drone.Repo{Owner: "drone-plugin", Name: "drone-hipchat"},
			build:  &drone.Build{Status: "Success", Number: 123, Author: "Self", Branch: "master", Started: 1234, Finished: 1235, Message: "Complete"},
			tmpl:   "{{ }",
			message: `Parse error on line 1:
Lexer error
Token: Error{"Unexpected character in expression: '}'"}`,
		},
	}

	for _, test := range tests {
		message := BuildTemplate(test.system, test.repo, test.build, test.tmpl)
		if test.message != message {
			t.Errorf("expected message:\n %s \n got message: \n %s \n", test.message, message)
		}
	}
}

func TestColor(t *testing.T) {
	tests := []struct {
		build *drone.Build
		color string
	}{
		{build: &drone.Build{Status: drone.StatusSkipped}, color: "yellow"},
		{build: &drone.Build{Status: drone.StatusPending}, color: "yellow"},
		{build: &drone.Build{Status: drone.StatusRunning}, color: "yellow"},
		{build: &drone.Build{Status: drone.StatusSuccess}, color: "green"},
		{build: &drone.Build{Status: drone.StatusFailure}, color: "red"},
		{build: &drone.Build{Status: drone.StatusKilled}, color: "red"},
		{build: &drone.Build{Status: drone.StatusError}, color: "red"},
		{build: &drone.Build{Status: "foobar"}, color: "yellow"},
	}

	for _, test := range tests {
		if Color(test.build) != test.color {
			t.Errorf("expected %s got %s", test.color, Color(test.build))
		}
	}
}
