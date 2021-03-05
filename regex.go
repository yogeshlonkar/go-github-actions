package actions

import (
	"regexp"
)

const (
	EXPRESSION_SYNTAX = `^\$\{\{.*\}\}$`
	JOB_ID_PATTERN    = `^[_a-zA-Z][a-zA-Z0-9_-]*`
	INPUT_ID_PATTERN  = `^[_a-zA-Z][a-zA-Z0-9_-]*$`
	CRON_PATTERN      = `^(((\d+,)+\d+|((\d+|\*)\/\d+|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)|(\d+-\d+)|\d+|\*|MON|TUE|WED|THU|FRI|SAT|SUN) ?){5,7}$`
	VOLUME_PATTERN    = `^[^:]+:[^:]+$`
)

var (
	expressionSyntax = regexp.MustCompile(EXPRESSION_SYNTAX)
	jobIDRegex       = regexp.MustCompile(JOB_ID_PATTERN)
	inputIDRegex     = regexp.MustCompile(INPUT_ID_PATTERN)
	cronRegex        = regexp.MustCompile(CRON_PATTERN)
	volumeRegex      = regexp.MustCompile(VOLUME_PATTERN)
)
