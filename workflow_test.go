package actions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
)

func TestWorkflow(t *testing.T) {
	workflow := &Workflow{
		Name: "Test Workflow",
		Defaults: &Defaults{
			Run: &Run{
				Shell:            "bash",
				WorkingDirectory: "/home",
			},
		},
		Env: map[string]interface{}{
			"ENV_1": 213,
			"ENV_2": "ABC",
			"ENV_3": false,
		},
		On: OnEventConfig{
			CheckSuite: &OnCheckSuite{
				Types: []OnCheckSuiteType{
					ON_CHECK_SUITE_TYPE_COMPLETED,
					ON_CHECK_SUITE_TYPE_REQUESTED,
				},
			},
			Schedule: []*OnScheduleCron{
				{
					Cron: "*/15 * * * *",
				},
			},
			CheckRun: &OnCheckRun{
				Types: []OnCheckRunType{
					ON_CHECK_RUN_TYPE_COMPLETED,
				},
			},
			Push: &OnPush{},
			WorkflowDispatch: &OnWorkflowDispatch{
				Inputs: map[string]*OnWorkflowDispatchInput{
					"some-thing": {
						Description: "something something",
						Required:    false,
					},
				},
			},
			Events: map[string]interface{}{
				string(ON_EVENT_CREATE): "",
			},
		},
		Jobs: map[string]*Job{
			"first-job": {
				Steps: []*Step{
					{
						Name: "Echo Hello",
						Run:  "echo \"Hello There!\"",
					},
					{
						Run: "echo \"Build something!\"",
					},
				},
			},
		},
	}
	out, err := yaml.Marshal(workflow)
	if err != nil {
		t.Error(err)
		return
	}
	// t.Error("\n" + strings.ReplaceAll(strings.ReplaceAll(string(out), "{}", ""), "\"\"", ""))
	expected := `defaults:
  run:
    shell: bash
    working-directory: /home
env:
  ENV_1: 213
  ENV_2: ABC
  ENV_3: false
jobs:
  first-job:
    steps:
    - name: Echo Hello
      run: echo "Hello There!"
    - run: echo "Build something!"
name: Test Workflow
"on":
  check_run:
    types:
    - completed
  check_suite:
    types:
    - completed
    - requested
  push: {}
  schedule:
  - cron: '*/15 * * * *'
  workflow_dispatch:
    inputs:
      some-thing:
        description: something something
        required: false
  create: ""
`
	got := string(out)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("Workflow yaml mismatch (-want +got):\n%s", diff)
	}

}
