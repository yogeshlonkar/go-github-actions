go-github-actions
====

Go types to create Github actions/ workflows that can be marshaled/ unmashaled in JSON and YAML

# Usage
```shell
go get github.com/yogeshlonkar/go-github-actions
```

```go
import (
    "github.com/yogeshlonkar/go-github-actions"
)

var workflow = actions.Workflow{
    Jobs: map[string]*actions.Job{
        "first-job": {
            RunsOn: actions.RUNS_ON_UBUNTU_LATEST,
            Steps: []*actions.Step{
                {
                    Name: "Echo Hello",
                    Run: "echo \"Hello There!\"",
                },
            },
        },
    },
}
```

