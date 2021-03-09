package actions

import (
	"encoding/json"
	"fmt"
)

type Step struct {
	Id               string                 `json:",omitempty" yaml:",omitempty"`
	Name             string                 `json:",omitempty" yaml:",omitempty"`
	Uses             string                 `json:",omitempty" yaml:",omitempty"`
	Env              map[string]interface{} `json:",omitempty" yaml:",omitempty"`
	Run              string                 `json:",omitempty" yaml:",omitempty"`
	Shell            shell                  `json:",omitempty" yaml:",omitempty"`
	WorkingDirectory string                 `json:"working-directory,omitempty" yaml:"working-directory,omitempty"`
	With             map[string]interface{} `json:",omitempty" yaml:",omitempty"`
	TimeoutMinutes   float64                `json:"timeout-minutes,omitempty" yaml:"timeout-minutes,omitempty"`
	ContinueOnError  interface{}            `json:"continue-on-error,omitempty" yaml:"continue-on-error,omitempty"`
	If               string                 `json:",omitempty" yaml:",omitempty"`
}

func (s Step) Validate() error {
	for _, val := range s.With {
		switch val.(type) {
		case string, bool, float32, float64, int, int32, int64:
		default:
			return fmt.Errorf("invalid With value \"%s\", expected boolean, number or string", val)
		}
	}
	switch v := s.ContinueOnError.(type) {
	case bool, nil:
	case string:
		if !expressionSyntax.Match([]byte(s.ContinueOnError.(string))) {
			return fmt.Errorf("invalid continue-on-error value \"%s\", expected boolean or string matching %s", s.ContinueOnError, EXPRESSION_SYNTAX)
		}
	default:
		return fmt.Errorf("invalid continue-on-error value \"%s\", expected boolean or string matching %s", v, EXPRESSION_SYNTAX)
	}
	return nil
}

func (s Step) MarshalJSON() ([]byte, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

func (s Step) MarshalYAML() (interface{}, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}
