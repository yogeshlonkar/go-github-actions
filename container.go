package actions

import (
	"encoding/json"
	"fmt"
)

type Container struct {
	Credentials *Credentials           `json:",omitempty" yaml:",omitempty"`
	Env         map[string]interface{} `json:",omitempty" yaml:",omitempty"`
	Image       string                 `json:",omitempty" yaml:",omitempty"`
	Options     string                 `json:",omitempty" yaml:",omitempty"`
	Ports       []interface{}          `json:",omitempty" yaml:",omitempty"`
	Volumes     []string               `json:",omitempty" yaml:",omitempty"`
}

type Credentials struct {
	Password string `json:",omitempty" yaml:",omitempty"`
	Username string `json:",omitempty" yaml:",omitempty"`
}

func (c Container) Validate() error {
	for _, volume := range c.Volumes {
		if !volumeRegex.Match([]byte(volume)) {
			return fmt.Errorf("invalid volume value \"%s\", expected string matching %s", volume, VOLUME_PATTERN)
		}
	}
	return nil
}

func (c Container) MarshalJSON() ([]byte, error) {
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(c)
}

func (c Container) MarshalYAML() (interface{}, error) {
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}
