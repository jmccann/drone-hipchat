package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drone/drone/model"
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
		// repo    *model.Repo
		// build   *model.Build
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
<<<<<<< HEAD
		message := BuildTemplate(test.system, test.repo, test.build, test.tmpl)
=======
		message := BuildMessage(test.plugin, test.tmpl)
>>>>>>> upstream/master
		if test.message != message {
			t.Errorf("expected message:\n %s \n got message: \n %s \n", test.message, message)
		}
	}
}

func TestColor(t *testing.T) {
	tests := []struct {
		build *model.Build
		color string
	}{
		{build: &model.Build{Status: model.StatusSkipped}, color: "yellow"},
		{build: &model.Build{Status: model.StatusPending}, color: "yellow"},
		{build: &model.Build{Status: model.StatusRunning}, color: "yellow"},
		{build: &model.Build{Status: model.StatusSuccess}, color: "green"},
		{build: &model.Build{Status: model.StatusFailure}, color: "red"},
		{build: &model.Build{Status: model.StatusKilled}, color: "red"},
		{build: &model.Build{Status: model.StatusError}, color: "red"},
		{build: &model.Build{Status: "foobar"}, color: "yellow"},
	}

	for _, test := range tests {
		if Color(test.build) != test.color {
			t.Errorf("expected %s got %s", test.color, Color(test.build))
		}
	}
}
