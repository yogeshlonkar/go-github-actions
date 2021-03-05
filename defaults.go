package actions

type Defaults struct {
	Run *Run `json:",omitempty" yaml:",omitempty"`
}

type shell string

type Run struct {
	Shell            shell  `json:",omitempty" yaml:",omitempty"`
	WorkingDirectory string `json:"working-directory,omitempty" yaml:"working-directory,omitempty"`
}

const (
	SHELL_BASH       shell = "bash"
	SHELL_PWSH       shell = "pwsh"
	SHELL_PYTHON     shell = "python"
	SHELL_SH         shell = "sh"
	SHELL_CMD        shell = "cmd"
	SHELL_POWERSHELL shell = "powershell"
)

func CustomShell(custom string) shell {
	return shell(custom)
}
