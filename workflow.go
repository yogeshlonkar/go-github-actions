package actions

import (
	"encoding/json"
	"fmt"
)

type Workflow struct {
	Defaults *Defaults              `json:",omitempty" yaml:",omitempty"`
	Env      map[string]interface{} `json:",omitempty" yaml:",omitempty"`
	Jobs     map[string]*Job        `json:",omitempty" yaml:",omitempty"`
	Name     string                 `json:",omitempty" yaml:",omitempty"`
	On       interface{}            `json:",omitempty" yaml:",omitempty"`
}

func (w Workflow) Validate() error {
	for jobID := range w.Jobs {
		if !jobIDRegex.Match([]byte(jobID)) {
			return fmt.Errorf("invalid job_id \"%s\", expected id to match %s", jobID, JOB_ID_PATTERN)
		}
	}
	switch v := w.On.(type) {
	case onEvent, OnEvents, OnEventConfig:
	default:
		return fmt.Errorf("unexpected value \"%s\" for \"On\" expected OnEvent, OnEvents or OnEventConfig", v)
	}
	return nil
}

func (w Workflow) MarshalJSON() ([]byte, error) {
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(w)
}

func (w Workflow) MarshalYAML() (interface{}, error) {
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return w, nil
}
