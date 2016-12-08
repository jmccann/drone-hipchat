package main

import (
	"testing"

  "github.com/drone/drone/model"
)

var tests = []struct {
	Plugin *Plugin
	Input   string
	Output  string
}{
	{
		&Plugin{Build: &model.Build{
			Commit: "0a266f42a9aef9db97a005ab46f6c53890339a9c"},
		},
		"{{ truncate build.commit 8 }}",
		"0a266f42",
	},
	{
		&Plugin{Build: &model.Build{Number: 1}},
		"build #{{build.number}}",
		"build #1",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusSuccess}},
		"{{uppercase build.status}}",
		"SUCCESS",
	},
	{
		&Plugin{Build: &model.Build{Author: "Octocat"}},
		"{{lowercase build.author}}",
		"octocat",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusSuccess}},
		"{{uppercasefirst build.status}}",
		"Success",
	},
	{
		&Plugin{Build: &model.Build{
			Started:  1448127131,
			Finished: 1448127505},
		},
		"{{ duration build.started build.finished }}",
		"6m14s",
	},
	{
		&Plugin{Build: &model.Build{Finished: 1448127505}},
		`finished at {{ datetime build.finished "3:04PM" "UTC" }}`,
		"finished at 5:38PM",
	},
	// verify the success if / else block works
	{
		&Plugin{Build: &model.Build{Status: model.StatusSuccess}},
		"{{#success build.status}}SUCCESS{{/success}}",
		"SUCCESS",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusFailure}},
		"{{#success build.status}}SUCCESS{{/success}}",
		"",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusFailure}},
		"{{#success build.status}}SUCCESS{{else}}NOT SUCCESS{{/success}}",
		"NOT SUCCESS",
	},
	// verify the failure if / else block works
	{
		&Plugin{Build: &model.Build{Status: model.StatusFailure}},
		"{{#failure build.status}}FAILURE{{/failure}}",
		"FAILURE",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusSuccess}},
		"{{#failure build.status}}FAILURE{{/failure}}",
		"",
	},
	{
		&Plugin{Build: &model.Build{Status: model.StatusSuccess}},
		"{{#failure build.status}}FAILURE{{else}}NOT FAILURE{{/failure}}",
		"NOT FAILURE",
	},
	{
		&Plugin{Build: &model.Build{Author: "url&unsafe=author!"}},
		"{{#urlencode}}google https://www.google.co.jp/ {{{build.author}}}{{/urlencode}}",
		"google+https%3A%2F%2Fwww.google.co.jp%2F+url%26unsafe%3Dauthor%21",
	},
}

func TestTemplate(t *testing.T) {

	for _, test := range tests {
		got, err := RenderTrim(test.Input, test.Plugin)
		if err != nil {
			t.Errorf("Failed rendering template %q, got error %s.", test.Input, err)
		}
		if got != test.Output {
			t.Errorf("Wanted rendered template %q, got %q", test.Output, got)
		}
	}
}
