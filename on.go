package actions

import (
	"encoding/json"
	"fmt"
)

type onEvent string

type OnEvents []onEvent

type OnEventConfig struct {
	CheckRun                 *OnCheckRun                 `json:"check_run,omitempty" yaml:"check_run,omitempty"`
	CheckSuite               *OnCheckSuite               `json:"check_suite,omitempty" yaml:"check_suite,omitempty"`
	IssueComment             *OnIssueComment             `json:"issue_comment,omitempty" yaml:"issue_comment,omitempty"`
	Issues                   *OnIssues                   `json:",omitempty" yaml:",omitempty"`
	Label                    *OnLabel                    `json:",omitempty" yaml:",omitempty"`
	Member                   *OnMember                   `json:",omitempty" yaml:",omitempty"`
	Milestone                *OnMilestone                `json:",omitempty" yaml:",omitempty"`
	Project                  *OnProject                  `json:",omitempty" yaml:",omitempty"`
	ProjectCard              *OnProjectCard              `json:"project_card,omitempty" yaml:"project_card,omitempty"`
	ProjectColumn            *OnProjectColumn            `json:"project_column,omitempty" yaml:"project_column,omitempty"`
	PullRequest              *OnPullRequest              `json:"pull_request,omitempty" yaml:"pull_request,omitempty"`
	PullRequestReview        *OnPullRequestReview        `json:"pull_request_review,omitempty" yaml:"pull_request_review,omitempty"`
	PullRequestReviewComment *OnPullRequestReviewComment `json:"pull_request_review_comment,omitempty" yaml:"pull_request_review_comment,omitempty"`
	PullRequestTarget        *OnPullRequestTarget        `json:"pull_request_target,omitempty" yaml:"pull_request_target,omitempty"`
	Push                     *OnPush                     `json:",omitempty" yaml:",omitempty"`
	RegistryPackage          *OnRegistryPackage          `json:"registry_package,omitempty" yaml:"registry_package,omitempty"`
	Release                  *OnRelease                  `json:",omitempty" yaml:",omitempty"`
	Schedule                 []*OnScheduleCron           `json:",omitempty" yaml:",omitempty"`
	WorkflowDispatch         *OnWorkflowDispatch         `json:"workflow_dispatch,omitempty" yaml:"workflow_dispatch,omitempty"`
	WorkflowRun              *OnWorkflowRun              `json:"workflow_run,omitempty" yaml:"workflow_run,omitempty"`
	Events                   map[string]interface{}      `json:",omitempty,inline" yaml:",omitempty,inline"`
}

// onCheckRunType
type onCheckRunType string

const (
	ON_CHECK_RUN_TYPE_CREATED          onCheckRunType = "created"
	ON_CHECK_RUN_TYPE_RE_REQUESTED     onCheckRunType = "rerequested"
	ON_CHECK_RUN_TYPE_COMPLETED        onCheckRunType = "completed"
	ON_CHECK_RUN_TYPE_REQUESTED_ACTION onCheckRunType = "requested_action"
)

type OnCheckRun struct {
	Types []onCheckRunType `json:",omitempty" yaml:",omitempty"`
}

// onCheckSuiteType
type onCheckSuiteType string

const (
	ON_CHECK_SUITE_TYPE_COMPLETED    onCheckSuiteType = "completed"
	ON_CHECK_SUITE_TYPE_REQUESTED    onCheckSuiteType = "requested"
	ON_CHECK_SUITE_TYPE_RE_REQUESTED onCheckSuiteType = "rerequested"
)

type OnCheckSuite struct {
	Types []onCheckSuiteType `json:",omitempty" yaml:",omitempty"`
}

// onIssueCommentType
type onIssueCommentType string

const (
	ON_ISSUE_COMMENT_TYPE_CREATED onIssueCommentType = "created"
	ON_ISSUE_COMMENT_TYPE_EDITED  onIssueCommentType = "edited"
	ON_ISSUE_COMMENT_TYPE_DELETED onIssueCommentType = "deleted"
)

type OnIssueComment struct {
	Types []onIssueCommentType `json:",omitempty" yaml:",omitempty"`
}

// onIssuesType
type onIssuesType string

const (
	ON_ISSUES_TYPE_OPENED        onIssuesType = "opened"
	ON_ISSUES_TYPE_EDITED        onIssuesType = "edited"
	ON_ISSUES_TYPE_DELETED       onIssuesType = "deleted"
	ON_ISSUES_TYPE_TRANSFERRED   onIssuesType = "transferred"
	ON_ISSUES_TYPE_PINNED        onIssuesType = "pinned"
	ON_ISSUES_TYPE_UNPINNED      onIssuesType = "unpinned"
	ON_ISSUES_TYPE_CLOSED        onIssuesType = "closed"
	ON_ISSUES_TYPE_REOPENED      onIssuesType = "reopened"
	ON_ISSUES_TYPE_ASSIGNED      onIssuesType = "assigned"
	ON_ISSUES_TYPE_UNASSIGNED    onIssuesType = "unassigned"
	ON_ISSUES_TYPE_LABELED       onIssuesType = "labeled"
	ON_ISSUES_TYPE_UNLABELED     onIssuesType = "unlabeled"
	ON_ISSUES_TYPE_LOCKED        onIssuesType = "locked"
	ON_ISSUES_TYPE_UNLOCKED      onIssuesType = "unlocked"
	ON_ISSUES_TYPE_MILESTONED    onIssuesType = "milestoned"
	ON_ISSUES_TYPE_DE_MILESTONED onIssuesType = "demilestoned"
)

type OnIssues struct {
	Types []onIssuesType `json:",omitempty" yaml:",omitempty"`
}

// onLabelType
type onLabelType string

const (
	ON_LABEL_TYPE_CREATED onLabelType = "created"
	ON_LABEL_TYPE_EDITED  onLabelType = "edited"
	ON_LABEL_TYPE_DELETED onLabelType = "deleted"
)

type OnLabel struct {
	Types []onLabelType `json:",omitempty" yaml:",omitempty"`
}

// onMemberType
type onMemberType string

const (
	ON_MEMBER_TYPE_ADDED   onMemberType = "added"
	ON_MEMBER_TYPE_EDITED  onMemberType = "edited"
	ON_MEMBER_TYPE_DELETED onMemberType = "deleted"
)

type OnMember struct {
	Types []onMemberType `json:",omitempty" yaml:",omitempty"`
}

// onMilestoneType
type onMilestoneType string

const (
	ON_MILESTONE_TYPE_CREATED onMilestoneType = "created"
	ON_MILESTONE_TYPE_CLOSED  onMilestoneType = "closed"
	ON_MILESTONE_TYPE_OPENED  onMilestoneType = "opened"
	ON_MILESTONE_TYPE_EDITED  onMilestoneType = "edited"
	ON_MILESTONE_TYPE_DELETED onMilestoneType = "deleted"
)

type OnMilestone struct {
	Types []onMilestoneType `json:",omitempty" yaml:",omitempty"`
}

// onProjectType
type onProjectType string

const (
	ON_PROJECT_TYPE_CREATED  onProjectType = "created"
	ON_PROJECT_TYPE_UPDATED  onProjectType = "updated"
	ON_PROJECT_TYPE_CLOSED   onProjectType = "closed"
	ON_PROJECT_TYPE_REOPENED onProjectType = "reopened"
	ON_PROJECT_TYPE_EDITED   onProjectType = "edited"
	ON_PROJECT_TYPE_DELETED  onProjectType = "deleted"
)

type OnProject struct {
	Types []onProjectType `json:",omitempty" yaml:",omitempty"`
}

// onProjectCardType
type onProjectCardType string

const (
	ON_PROJECT_CARD_TYPE_CREATED   onProjectCardType = "created"
	ON_PROJECT_CARD_TYPE_MOVED     onProjectCardType = "moved"
	ON_PROJECT_CARD_TYPE_CONVERTED onProjectCardType = "converted"
	ON_PROJECT_CARD_TYPE_EDITED    onProjectCardType = "edited"
	ON_PROJECT_CARD_TYPE_DELETED   onProjectCardType = "deleted"
)

type OnProjectCard struct {
	Types []onProjectCardType `json:",omitempty" yaml:",omitempty"`
}

// onProjectColumnType
type onProjectColumnType string

const (
	ON_PROJECT_COLUMN_TYPE_CREATED onProjectColumnType = "created"
	ON_PROJECT_COLUMN_TYPE_UPDATED onProjectColumnType = "updated"
	ON_PROJECT_COLUMN_TYPE_MOVED   onProjectColumnType = "moved"
	ON_PROJECT_COLUMN_TYPE_DELETED onProjectColumnType = "deleted"
)

type OnProjectColumn struct {
	Types []onProjectColumnType `json:",omitempty" yaml:",omitempty"`
}

// onPullRequestType
type onPullRequestType string

const (
	ON_PULL_REQUEST_TYPE_ASSIGNED               onPullRequestType = "assigned"
	ON_PULL_REQUEST_TYPE_UNASSIGNED             onPullRequestType = "unassigned"
	ON_PULL_REQUEST_TYPE_LABELED                onPullRequestType = "labeled"
	ON_PULL_REQUEST_TYPE_UNLABELED              onPullRequestType = "unlabeled"
	ON_PULL_REQUEST_TYPE_OPENED                 onPullRequestType = "opened"
	ON_PULL_REQUEST_TYPE_EDITED                 onPullRequestType = "edited"
	ON_PULL_REQUEST_TYPE_CLOSED                 onPullRequestType = "closed"
	ON_PULL_REQUEST_TYPE_REOPENED               onPullRequestType = "reopened"
	ON_PULL_REQUEST_TYPE_SYNCHRONIZE            onPullRequestType = "synchronize"
	ON_PULL_REQUEST_TYPE_READY_FOR_REVIEW       onPullRequestType = "ready_for_review"
	ON_PULL_REQUEST_TYPE_LOCKED                 onPullRequestType = "locked"
	ON_PULL_REQUEST_TYPE_UNLOCKED               onPullRequestType = "unlocked"
	ON_PULL_REQUEST_TYPE_REVIEW_REQUESTED       onPullRequestType = "review_requested"
	ON_PULL_REQUEST_TYPE_REVIEW_REQUEST_REMOVED onPullRequestType = "review_request_removed"
)

type OnPullRequest struct {
	Types          []onPullRequestType `json:",omitempty" yaml:",omitempty"`
	Branches       []string            `json:",omitempty" yaml:",omitempty"`
	BranchesIgnore []string            `json:"branches-ignore,omitempty" yaml:"branches-ignore,omitempty"`
	Tags           []string            `json:",omitempty" yaml:",omitempty"`
	TagsIgnore     []string            `json:"tags-ignore,omitempty" yaml:"tags-ignore,omitempty"`
	Paths          []string            `json:",omitempty" yaml:",omitempty"`
	PathsIgnore    []string            `json:"paths-ignore,omitempty" yaml:"paths-ignore,omitempty"`
}

// onPullRequestReviewType
type onPullRequestReviewType string

const (
	ON_PULL_REQUEST_REVIEW_TYPE_SUBMITTED onPullRequestReviewType = "submitted"
	ON_PULL_REQUEST_REVIEW_TYPE_EDITED    onPullRequestReviewType = "edited"
	ON_PULL_REQUEST_REVIEW_TYPE_DISMISSED onPullRequestReviewType = "dismissed"
)

type OnPullRequestReview struct {
	Types []onPullRequestReviewType `json:",omitempty" yaml:",omitempty"`
}

// onPullRequestReviewCommentType
type onPullRequestReviewCommentType string

const (
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_CREATED onPullRequestReviewCommentType = "created"
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_EDITED  onPullRequestReviewCommentType = "edited"
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_DELETED onPullRequestReviewCommentType = "deleted"
)

type OnPullRequestReviewComment struct {
	Types []onPullRequestReviewCommentType `json:",omitempty" yaml:",omitempty"`
}

// onPullRequestTargetType
type onPullRequestTargetType string

const (
	ON_PULL_REQUEST_TARGET_TYPE_ASSIGNED               onPullRequestTargetType = "assigned"
	ON_PULL_REQUEST_TARGET_TYPE_UNASSIGNED             onPullRequestTargetType = "unassigned"
	ON_PULL_REQUEST_TARGET_TYPE_LABELED                onPullRequestTargetType = "labeled"
	ON_PULL_REQUEST_TARGET_TYPE_UNLABELED              onPullRequestTargetType = "unlabeled"
	ON_PULL_REQUEST_TARGET_TYPE_OPENED                 onPullRequestTargetType = "opened"
	ON_PULL_REQUEST_TARGET_TYPE_EDITED                 onPullRequestTargetType = "edited"
	ON_PULL_REQUEST_TARGET_TYPE_CLOSED                 onPullRequestTargetType = "closed"
	ON_PULL_REQUEST_TARGET_TYPE_REOPENED               onPullRequestTargetType = "reopened"
	ON_PULL_REQUEST_TARGET_TYPE_SYNCHRONIZE            onPullRequestTargetType = "synchronize"
	ON_PULL_REQUEST_TARGET_TYPE_READY_FOR_REVIEW       onPullRequestTargetType = "ready_for_review"
	ON_PULL_REQUEST_TARGET_TYPE_LOCKED                 onPullRequestTargetType = "locked"
	ON_PULL_REQUEST_TARGET_TYPE_UNLOCKED               onPullRequestTargetType = "unlocked"
	ON_PULL_REQUEST_TARGET_TYPE_REVIEW_REQUESTED       onPullRequestTargetType = "review_requested"
	ON_PULL_REQUEST_TARGET_TYPE_REVIEW_REQUEST_REMOVED onPullRequestTargetType = "review_request_removed"
)

type OnPullRequestTarget struct {
	Types          []onPullRequestTargetType `json:",omitempty" yaml:",omitempty"`
	Branches       []string                  `json:",omitempty" yaml:",omitempty"`
	BranchesIgnore []string                  `json:"branches-ignore,omitempty" yaml:"branches-ignore,omitempty"`
	Tags           []string                  `json:",omitempty" yaml:",omitempty"`
	TagsIgnore     []string                  `json:"tags-ignore,omitempty" yaml:"tags-ignore,omitempty"`
	Paths          []string                  `json:",omitempty" yaml:",omitempty"`
	PathsIgnore    []string                  `json:"paths-ignore,omitempty" yaml:"paths-ignore,omitempty"`
}

// OnPush
type OnPush struct {
	Branches       []string `json:",omitempty" yaml:",omitempty"`
	BranchesIgnore []string `json:"branches-ignore,omitempty" yaml:"branches-ignore,omitempty"`
	Tags           []string `json:",omitempty" yaml:",omitempty"`
	TagsIgnore     []string `json:"tags-ignore,omitempty" yaml:"tags-ignore,omitempty"`
	Paths          []string `json:",omitempty" yaml:",omitempty"`
	PathsIgnore    []string `json:"paths-ignore,omitempty" yaml:"paths-ignore,omitempty"`
}

// onRegistryPackageType
type onRegistryPackageType string

const (
	ON_REGISTRY_PACKAGE_TYPE_PUBLISHED onRegistryPackageType = "published"
	ON_REGISTRY_PACKAGE_TYPE_UPDATED   onRegistryPackageType = "updated"
)

type OnRegistryPackage struct {
	Types []onRegistryPackageType `json:",omitempty" yaml:",omitempty"`
}

// onReleaseType
type onReleaseType string

const (
	ON_RELEASE_TYPE_PUBLISHED   onReleaseType = "published"
	ON_RELEASE_TYPE_UNPUBLISHED onReleaseType = "unpublished"
	ON_RELEASE_TYPE_CREATED     onReleaseType = "created"
	ON_RELEASE_TYPE_EDITED      onReleaseType = "edited"
	ON_RELEASE_TYPE_DELETED     onReleaseType = "deleted"
	ON_RELEASE_TYPE_PRERELEASED onReleaseType = "prereleased"
	ON_RELEASE_TYPE_RELEASED    onReleaseType = "released"
)

type OnRelease struct {
	Types []onReleaseType
}

// OnWorkflowDispatch
type OnWorkflowDispatch struct {
	// InputId should match ^[_a-zA-Z][a-zA-Z0-9_-]*$
	Inputs map[string]*OnWorkflowDispatchInput `json:",omitempty" yaml:",omitempty"`
}

type OnWorkflowDispatchInput struct {
	Description        string `json:",omitempty" yaml:",omitempty"`
	DeprecationMessage string `json:"deprecationMessage,omitempty" yaml:"deprecationMessage,omitempty"`
	Required           bool
	Default            string `json:",omitempty" yaml:",omitempty"`
}

// onWorkflowRunType
type onWorkflowRunType string

const (
	ON_WORKFLOW_RUN_TYPE_REQUESTED onWorkflowRunType = "requested"
	ON_WORKFLOW_RUN_TYPE_COMPLETED onWorkflowRunType = "completed"
)

type OnWorkflowRun struct {
	Types          []onWorkflowRunType `json:",omitempty" yaml:",omitempty"`
	Workflows      []string            `json:",omitempty" yaml:",omitempty"`
	Branches       []string            `json:",omitempty" yaml:",omitempty"`
	BranchesIgnore []string            `json:"branches-ignore,omitempty" yaml:"branches-ignore,omitempty"`
}

// OnScheduleCron
type OnScheduleCron struct {
	// should match ^(((\d+,)+\d+|((\d+|\*)\/\d+|JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)|(\d+-\d+)|\d+|\*|MON|TUE|WED|THU|FRI|SAT|SUN) ?){5,7}$
	Cron string `json:",omitempty" yaml:",omitempty"`
}

const (
	ON_EVENT_CHECK_RUN                   onEvent = "check_run"
	ON_EVENT_CHECK_SUITE                 onEvent = "check_suite"
	ON_EVENT_CREATE                      onEvent = "create"
	ON_EVENT_DELETE                      onEvent = "delete"
	ON_EVENT_DEPLOYMENT                  onEvent = "deployment"
	ON_EVENT_DEPLOYMENT_STATUS           onEvent = "deployment_status"
	ON_EVENT_FORK                        onEvent = "fork"
	ON_EVENT_GOLLUM                      onEvent = "gollum"
	ON_EVENT_ISSUES                      onEvent = "issues"
	ON_EVENT_ISSUE_COMMENT               onEvent = "issue_comment"
	ON_EVENT_LABEL                       onEvent = "label"
	ON_EVENT_MEMBER                      onEvent = "member"
	ON_EVENT_MILESTONE                   onEvent = "milestone"
	ON_EVENT_PAGE_BUILD                  onEvent = "page_build"
	ON_EVENT_PROJECT                     onEvent = "project"
	ON_EVENT_PROJECT_CARD                onEvent = "project_card"
	ON_EVENT_PROJECT_COLUMN              onEvent = "project_column"
	ON_EVENT_PUBLIC                      onEvent = "public"
	ON_EVENT_PULL_REQUEST                onEvent = "pull_request"
	ON_EVENT_PULL_REQUEST_REVIEW         onEvent = "pull_request_review"
	ON_EVENT_PULL_REQUEST_REVIEW_COMMENT onEvent = "pull_request_review_comment"
	ON_EVENT_PULL_REQUEST_TARGET         onEvent = "pull_request_target"
	ON_EVENT_PUSH                        onEvent = "push"
	ON_EVENT_REGISTRY_PACKAGE            onEvent = "registry_package"
	ON_EVENT_RELEASE                     onEvent = "release"
	ON_EVENT_REPOSITORY_DISPATCH         onEvent = "repository_dispatch"
	ON_EVENT_STATUS                      onEvent = "status"
	ON_EVENT_WATCH                       onEvent = "watch"
	ON_EVENT_WORKFLOW_DISPATCH           onEvent = "workflow_dispatch"
	ON_EVENT_WORKFLOW_RUN                onEvent = "workflow_run"
)

func (o OnEventConfig) Validate() error {
	for event := range o.Events {
		v := onEvent(event)
		switch v {
		case ON_EVENT_CHECK_RUN, ON_EVENT_CHECK_SUITE, ON_EVENT_CREATE, ON_EVENT_DELETE, ON_EVENT_DEPLOYMENT, ON_EVENT_DEPLOYMENT_STATUS, ON_EVENT_FORK, ON_EVENT_GOLLUM, ON_EVENT_ISSUES, ON_EVENT_ISSUE_COMMENT, ON_EVENT_LABEL, ON_EVENT_MEMBER, ON_EVENT_MILESTONE, ON_EVENT_PAGE_BUILD, ON_EVENT_PROJECT, ON_EVENT_PROJECT_CARD, ON_EVENT_PROJECT_COLUMN, ON_EVENT_PUBLIC, ON_EVENT_PULL_REQUEST, ON_EVENT_PULL_REQUEST_REVIEW, ON_EVENT_PULL_REQUEST_REVIEW_COMMENT, ON_EVENT_PULL_REQUEST_TARGET, ON_EVENT_PUSH, ON_EVENT_REGISTRY_PACKAGE, ON_EVENT_RELEASE, ON_EVENT_REPOSITORY_DISPATCH, ON_EVENT_STATUS, ON_EVENT_WATCH, ON_EVENT_WORKFLOW_DISPATCH, ON_EVENT_WORKFLOW_RUN:
		default:
			return fmt.Errorf("unexpected event \"%s\" in OnEventConfig", v)
		}
	}
	return nil
}

func (o OnEventConfig) MarshalJSON() ([]byte, error) {
	err := o.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(o)
}

func (o OnEventConfig) MarshalYAML() (interface{}, error) {
	err := o.Validate()
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s OnScheduleCron) Validate() error {
	if !cronRegex.Match([]byte(s.Cron)) {
		return fmt.Errorf("invalid cron value \"%s\", expected value to match %s", s.Cron, CRON_PATTERN)
	}
	return nil
}

func (s OnScheduleCron) MarshalJSON() ([]byte, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

func (s OnScheduleCron) MarshalYAML() (interface{}, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (w OnWorkflowDispatch) Validate() error {
	for inputID := range w.Inputs {
		if !inputIDRegex.Match([]byte(inputID)) {
			return fmt.Errorf("invalid workflow input id \"%s\", expected value to match ^[_a-zA-Z][a-zA-Z0-9_-]*$", inputID)
		}
	}
	return nil
}

func (w OnWorkflowDispatch) MarshalJSON() ([]byte, error) {
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return json.Marshal(w)
}

func (w OnWorkflowDispatch) MarshalYAML() (interface{}, error) {
	err := w.Validate()
	if err != nil {
		return nil, err
	}
	return w, nil
}
