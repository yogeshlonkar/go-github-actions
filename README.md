go-github-actions
====

Go types to create Github actions/ workflows that can be marshaled/ unmashaled in JSON and YAML. Types are created with reference to [github-workflow schema](https://json.schemastore.org/github-workflow).

# Usage
```shell
go get github.com/yogeshlonkar/go-github-actions
```

```go
import (
    actions "github.com/yogeshlonkar/go-github-actions"
)

var workflow1 = actions.Workflow{
	On: actions.ON_EVENT_PUSH,
	Jobs: map[string]*actions.Job{
		"first-job": {
			RunsOn: actions.RUNS_ON_UBUNTU_LATEST,
			Steps: []*actions.Step{
				{
					Name: "Echo Hello",
					Run:  "echo \"Hello There!\"",
				},
			},
		},
	},
}
```

