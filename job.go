package actions

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	Container       *Container             `json:",omitempty" yaml:",omitempty"`
	ContinueOnError interface{}            `json:"continue-on-error,omitempty" yaml:"continue-on-error,omitempty"`
	Defaults        *Defaults              `json:",omitempty" yaml:",omitempty"`
	Env             map[string]interface{} `json:",omitempty" yaml:",omitempty"`
	Environment     *Environment           `json:",omitempty" yaml:",omitempty"`
	If              string                 `json:",omitempty" yaml:",omitempty"`
	Name            string                 `json:",omitempty" yaml:",omitempty"`
	Needs           needs                  `json:",omitempty" yaml:",omitempty"`
	Outputs         map[string]string      `json:",omitempty" yaml:",omitempty"`
	RunsOn          RunsOn                 `json:"runs-on,omitempty" yaml:"runs-on,omitempty"`
	Services        map[string]*Container  `json:",omitempty" yaml:",omitempty"`
	Steps           []*Step                `json:",omitempty" yaml:",omitempty"`
	Strategy        *Strategy              `json:",omitempty" yaml:",omitempty"`
	TimeoutMinutes  float64                `json:"timeout-minutes,omitempty" yaml:"timeout-minutes,omitempty"`
}

type needs interface{}

type RunsOn interface{}

type Environment struct {
	Name string `json:",omitempty" yaml:",omitempty"`
	Url  string `json:",omitempty" yaml:",omitempty"`
}

type Strategy struct {
	Matrix      *Matrix `json:",omitempty" yaml:",omitempty"`
	FailFast    bool    `json:"fail-fast,omitempty" yaml:"fail-fast,omitempty"`
	MaxParallel float64 `json:"max-parallel,omitempty" yaml:"max-parallel,omitempty"`
}

type Matrix struct {
	Include []interface{}          `json:",omitempty" yaml:",omitempty"`
	Exclude []interface{}          `json:",omitempty" yaml:",omitempty"`
	Custom  map[string]interface{} `json:",omitempty,inline" yaml:",omitempty,inline"`
}

const (
	RUNS_ON_ARM32          = "ARM32"
	RUNS_ON_LINUX          = "linux"
	RUNS_ON_MACOS          = "macos"
	RUNS_ON_MACOS_10_15    = "macos-10.15"
	RUNS_ON_MACOS_11_0     = "macos-11.0"
	RUNS_ON_MACOS_LATEST   = "macos-latest"
	RUNS_ON_SELF_HOSTED    = "self-hosted"
	RUNS_ON_UBUNTU_16_04   = "ubuntu-16.04"
	RUNS_ON_UBUNTU_18_04   = "ubuntu-18.04"
	RUNS_ON_UBUNTU_20_04   = "ubuntu-20.04"
	RUNS_ON_UBUNTU_LATEST  = "ubuntu-latest"
	RUNS_ON_WINDOWS        = "windows"
	RUNS_ON_WINDOWS_2016   = "windows-2016"
	RUNS_ON_WINDOWS_2019   = "windows-2019"
	RUNS_ON_WINDOWS_LATEST = "windows-latest"
	RUNS_ON_X64            = "x64"
	RUNS_ON_X86            = "x86"
)

func (j Job) Validate() error {
	switch v := j.ContinueOnError.(type) {
	case bool, nil:
	case string:
		if !expressionSyntax.Match([]byte(j.ContinueOnError.(string))) {
			return fmt.Errorf("invalid continue-on-error value \"%s\", expected boolean or string matching ^\\$\\{\\{.*\\}\\}$", j.ContinueOnError)
		}
	default:
		return fmt.Errorf("invalid continue-on-error value \"%s\", expected boolean or string matching ^\\$\\{\\{.*\\}\\}$", v)
	}
	return nil
}

func (j Job) MarshalJSON() ([]byte, error) {
	err := j.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(j)
}

func (j Job) MarshalYAML() (interface{}, error) {
	err := j.Validate()
	if err != nil {
		return nil, err
	}
	return j, nil
}

func NewNeeds(jobNames ...string) needs {
	for _, jobID := range jobNames {
		if !jobIDRegex.Match([]byte(jobID)) {
			return fmt.Errorf("invalid job_id \"%s\" for needs, expected id to match %s", jobID, JOB_ID_PATTERN)
		}
	}
	if len(jobNames) == 1 {
		return jobNames[0]
	}
	return jobNames
}

func NewRunsOn(machine ...string) RunsOn {
	if len(machine) == 1 {
		return machine[0]
	}
	return machine
}
