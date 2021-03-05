package actions

import (
	"encoding/json"
	"fmt"
)

type OnEvent string

type OnEvents []OnEvent

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

// OnCheckRunType
type OnCheckRunType string

const (
	ON_CHECK_RUN_TYPE_CREATED          OnCheckRunType = "created"
	ON_CHECK_RUN_TYPE_RE_REQUESTED     OnCheckRunType = "rerequested"
	ON_CHECK_RUN_TYPE_COMPLETED        OnCheckRunType = "completed"
	ON_CHECK_RUN_TYPE_REQUESTED_ACTION OnCheckRunType = "requested_action"
)

type OnCheckRun struct {
	Types []OnCheckRunType `json:",omitempty" yaml:",omitempty"`
}

// OnCheckSuiteType
type OnCheckSuiteType string

const (
	ON_CHECK_SUITE_TYPE_COMPLETED    OnCheckSuiteType = "completed"
	ON_CHECK_SUITE_TYPE_REQUESTED    OnCheckSuiteType = "requested"
	ON_CHECK_SUITE_TYPE_RE_REQUESTED OnCheckSuiteType = "rerequested"
)

type OnCheckSuite struct {
	Types []OnCheckSuiteType `json:",omitempty" yaml:",omitempty"`
}

// OnIssueCommentType
type OnIssueCommentType string

const (
	ON_ISSUE_COMMENT_TYPE_CREATED OnIssueCommentType = "created"
	ON_ISSUE_COMMENT_TYPE_EDITED  OnIssueCommentType = "edited"
	ON_ISSUE_COMMENT_TYPE_DELETED OnIssueCommentType = "deleted"
)

type OnIssueComment struct {
	Types []OnIssueCommentType `json:",omitempty" yaml:",omitempty"`
}

// OnIssuesType
type OnIssuesType string

const (
	ON_ISSUES_TYPE_OPENED        OnIssuesType = "opened"
	ON_ISSUES_TYPE_EDITED        OnIssuesType = "edited"
	ON_ISSUES_TYPE_DELETED       OnIssuesType = "deleted"
	ON_ISSUES_TYPE_TRANSFERRED   OnIssuesType = "transferred"
	ON_ISSUES_TYPE_PINNED        OnIssuesType = "pinned"
	ON_ISSUES_TYPE_UNPINNED      OnIssuesType = "unpinned"
	ON_ISSUES_TYPE_CLOSED        OnIssuesType = "closed"
	ON_ISSUES_TYPE_REOPENED      OnIssuesType = "reopened"
	ON_ISSUES_TYPE_ASSIGNED      OnIssuesType = "assigned"
	ON_ISSUES_TYPE_UNASSIGNED    OnIssuesType = "unassigned"
	ON_ISSUES_TYPE_LABELED       OnIssuesType = "labeled"
	ON_ISSUES_TYPE_UNLABELED     OnIssuesType = "unlabeled"
	ON_ISSUES_TYPE_LOCKED        OnIssuesType = "locked"
	ON_ISSUES_TYPE_UNLOCKED      OnIssuesType = "unlocked"
	ON_ISSUES_TYPE_MILESTONED    OnIssuesType = "milestoned"
	ON_ISSUES_TYPE_DE_MILESTONED OnIssuesType = "demilestoned"
)

type OnIssues struct {
	Types []OnIssuesType `json:",omitempty" yaml:",omitempty"`
}

// OnLabelType
type OnLabelType string

const (
	ON_LABEL_TYPE_CREATED OnLabelType = "created"
	ON_LABEL_TYPE_EDITED  OnLabelType = "edited"
	ON_LABEL_TYPE_DELETED OnLabelType = "deleted"
)

type OnLabel struct {
	Types []OnLabelType `json:",omitempty" yaml:",omitempty"`
}

// OnMemberType
type OnMemberType string

const (
	ON_MEMBER_TYPE_ADDED   OnMemberType = "added"
	ON_MEMBER_TYPE_EDITED  OnMemberType = "edited"
	ON_MEMBER_TYPE_DELETED OnMemberType = "deleted"
)

type OnMember struct {
	Types []OnMemberType `json:",omitempty" yaml:",omitempty"`
}

// OnMilestoneType
type OnMilestoneType string

const (
	ON_MILESTONE_TYPE_CREATED OnMilestoneType = "created"
	ON_MILESTONE_TYPE_CLOSED  OnMilestoneType = "closed"
	ON_MILESTONE_TYPE_OPENED  OnMilestoneType = "opened"
	ON_MILESTONE_TYPE_EDITED  OnMilestoneType = "edited"
	ON_MILESTONE_TYPE_DELETED OnMilestoneType = "deleted"
)

type OnMilestone struct {
	Types []OnMilestoneType `json:",omitempty" yaml:",omitempty"`
}

// OnProjectType
type OnProjectType string

const (
	ON_PROJECT_TYPE_CREATED  OnProjectType = "created"
	ON_PROJECT_TYPE_UPDATED  OnProjectType = "updated"
	ON_PROJECT_TYPE_CLOSED   OnProjectType = "closed"
	ON_PROJECT_TYPE_REOPENED OnProjectType = "reopened"
	ON_PROJECT_TYPE_EDITED   OnProjectType = "edited"
	ON_PROJECT_TYPE_DELETED  OnProjectType = "deleted"
)

type OnProject struct {
	Types []OnProjectType `json:",omitempty" yaml:",omitempty"`
}

// OnProjectCardType
type OnProjectCardType string

const (
	ON_PROJECT_CARD_TYPE_CREATED   OnProjectCardType = "created"
	ON_PROJECT_CARD_TYPE_MOVED     OnProjectCardType = "moved"
	ON_PROJECT_CARD_TYPE_CONVERTED OnProjectCardType = "converted"
	ON_PROJECT_CARD_TYPE_EDITED    OnProjectCardType = "edited"
	ON_PROJECT_CARD_TYPE_DELETED   OnProjectCardType = "deleted"
)

type OnProjectCard struct {
	Types []OnProjectCardType `json:",omitempty" yaml:",omitempty"`
}

// OnProjectColumnType
type OnProjectColumnType string

const (
	ON_PROJECT_COLUMN_TYPE_CREATED OnProjectColumnType = "created"
	ON_PROJECT_COLUMN_TYPE_UPDATED OnProjectColumnType = "updated"
	ON_PROJECT_COLUMN_TYPE_MOVED   OnProjectColumnType = "moved"
	ON_PROJECT_COLUMN_TYPE_DELETED OnProjectColumnType = "deleted"
)

type OnProjectColumn struct {
	Types []OnProjectColumnType `json:",omitempty" yaml:",omitempty"`
}

// OnPullRequestType
type OnPullRequestType string

const (
	ON_PULL_REQUEST_TYPE_ASSIGNED               OnPullRequestType = "assigned"
	ON_PULL_REQUEST_TYPE_UNASSIGNED             OnPullRequestType = "unassigned"
	ON_PULL_REQUEST_TYPE_LABELED                OnPullRequestType = "labeled"
	ON_PULL_REQUEST_TYPE_UNLABELED              OnPullRequestType = "unlabeled"
	ON_PULL_REQUEST_TYPE_OPENED                 OnPullRequestType = "opened"
	ON_PULL_REQUEST_TYPE_EDITED                 OnPullRequestType = "edited"
	ON_PULL_REQUEST_TYPE_CLOSED                 OnPullRequestType = "closed"
	ON_PULL_REQUEST_TYPE_REOPENED               OnPullRequestType = "reopened"
	ON_PULL_REQUEST_TYPE_SYNCHRONIZE            OnPullRequestType = "synchronize"
	ON_PULL_REQUEST_TYPE_READY_FOR_REVIEW       OnPullRequestType = "ready_for_review"
	ON_PULL_REQUEST_TYPE_LOCKED                 OnPullRequestType = "locked"
	ON_PULL_REQUEST_TYPE_UNLOCKED               OnPullRequestType = "unlocked"
	ON_PULL_REQUEST_TYPE_REVIEW_REQUESTED       OnPullRequestType = "review_requested"
	ON_PULL_REQUEST_TYPE_REVIEW_REQUEST_REMOVED OnPullRequestType = "review_request_removed"
)

type OnPullRequest struct {
	Types          []OnPullRequestType `json:",omitempty" yaml:",omitempty"`
	Branches       []string            `json:",omitempty" yaml:",omitempty"`
	BranchesIgnore []string            `json:"branches-ignore,omitempty" yaml:"branches-ignore,omitempty"`
	Tags           []string            `json:",omitempty" yaml:",omitempty"`
	TagsIgnore     []string            `json:"tags-ignore,omitempty" yaml:"tags-ignore,omitempty"`
	Paths          []string            `json:",omitempty" yaml:",omitempty"`
	PathsIgnore    []string            `json:"paths-ignore,omitempty" yaml:"paths-ignore,omitempty"`
}

// OnPullRequestReviewType
type OnPullRequestReviewType string

const (
	ON_PULL_REQUEST_REVIEW_TYPE_SUBMITTED OnPullRequestReviewType = "submitted"
	ON_PULL_REQUEST_REVIEW_TYPE_EDITED    OnPullRequestReviewType = "edited"
	ON_PULL_REQUEST_REVIEW_TYPE_DISMISSED OnPullRequestReviewType = "dismissed"
)

type OnPullRequestReview struct {
	Types []OnPullRequestReviewType `json:",omitempty" yaml:",omitempty"`
}

// OnPullRequestReviewCommentType
type OnPullRequestReviewCommentType string

const (
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_CREATED OnPullRequestReviewCommentType = "created"
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_EDITED  OnPullRequestReviewCommentType = "edited"
	ON_PULL_REQUEST_REVIEW_COMMENT_TYPE_DELETED OnPullRequestReviewCommentType = "deleted"
)

type OnPullRequestReviewComment struct {
	Types []OnPullRequestReviewCommentType `json:",omitempty" yaml:",omitempty"`
}

// OnPullRequestTargetType
type OnPullRequestTargetType string

const (
	ON_PULL_REQUEST_TARGET_TYPE_ASSIGNED               OnPullRequestTargetType = "assigned"
	ON_PULL_REQUEST_TARGET_TYPE_UNASSIGNED             OnPullRequestTargetType = "unassigned"
	ON_PULL_REQUEST_TARGET_TYPE_LABELED                OnPullRequestTargetType = "labeled"
	ON_PULL_REQUEST_TARGET_TYPE_UNLABELED              OnPullRequestTargetType = "unlabeled"
	ON_PULL_REQUEST_TARGET_TYPE_OPENED                 OnPullRequestTargetType = "opened"
	ON_PULL_REQUEST_TARGET_TYPE_EDITED                 OnPullRequestTargetType = "edited"
	ON_PULL_REQUEST_TARGET_TYPE_CLOSED                 OnPullRequestTargetType = "closed"
	ON_PULL_REQUEST_TARGET_TYPE_REOPENED               OnPullRequestTargetType = "reopened"
	ON_PULL_REQUEST_TARGET_TYPE_SYNCHRONIZE            OnPullRequestTargetType = "synchronize"
	ON_PULL_REQUEST_TARGET_TYPE_READY_FOR_REVIEW       OnPullRequestTargetType = "ready_for_review"
	ON_PULL_REQUEST_TARGET_TYPE_LOCKED                 OnPullRequestTargetType = "locked"
	ON_PULL_REQUEST_TARGET_TYPE_UNLOCKED               OnPullRequestTargetType = "unlocked"
	ON_PULL_REQUEST_TARGET_TYPE_REVIEW_REQUESTED       OnPullRequestTargetType = "review_requested"
	ON_PULL_REQUEST_TARGET_TYPE_REVIEW_REQUEST_REMOVED OnPullRequestTargetType = "review_request_removed"
)

type OnPullRequestTarget struct {
	Types          []OnPullRequestTargetType `json:",omitempty" yaml:",omitempty"`
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

// OnRegistryPackageType
type OnRegistryPackageType string

const (
	ON_REGISTRY_PACKAGE_TYPE_PUBLISHED OnRegistryPackageType = "published"
	ON_REGISTRY_PACKAGE_TYPE_UPDATED   OnRegistryPackageType = "updated"
)

type OnRegistryPackage struct {
	Types []OnRegistryPackageType `json:",omitempty" yaml:",omitempty"`
}

// OnReleaseType
type OnReleaseType string

const (
	ON_RELEASE_TYPE_PUBLISHED   OnReleaseType = "published"
	ON_RELEASE_TYPE_UNPUBLISHED OnReleaseType = "unpublished"
	ON_RELEASE_TYPE_CREATED     OnReleaseType = "created"
	ON_RELEASE_TYPE_EDITED      OnReleaseType = "edited"
	ON_RELEASE_TYPE_DELETED     OnReleaseType = "deleted"
	ON_RELEASE_TYPE_PRERELEASED OnReleaseType = "prereleased"
	ON_RELEASE_TYPE_RELEASED    OnReleaseType = "released"
)

type OnRelease struct {
	Types []OnReleaseType
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

// OnWorkflowRunType
type OnWorkflowRunType string

const (
	ON_WORKFLOW_RUN_TYPE_REQUESTED OnWorkflowRunType = "requested"
	ON_WORKFLOW_RUN_TYPE_COMPLETED OnWorkflowRunType = "completed"
)

type OnWorkflowRun struct {
	Types          []OnWorkflowRunType `json:",omitempty" yaml:",omitempty"`
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
	ON_EVENT_CHECK_RUN                   OnEvent = "check_run"
	ON_EVENT_CHECK_SUITE                 OnEvent = "check_suite"
	ON_EVENT_CREATE                      OnEvent = "create"
	ON_EVENT_DELETE                      OnEvent = "delete"
	ON_EVENT_DEPLOYMENT                  OnEvent = "deployment"
	ON_EVENT_DEPLOYMENT_STATUS           OnEvent = "deployment_status"
	ON_EVENT_FORK                        OnEvent = "fork"
	ON_EVENT_GOLLUM                      OnEvent = "gollum"
	ON_EVENT_ISSUES                      OnEvent = "issues"
	ON_EVENT_ISSUE_COMMENT               OnEvent = "issue_comment"
	ON_EVENT_LABEL                       OnEvent = "label"
	ON_EVENT_MEMBER                      OnEvent = "member"
	ON_EVENT_MILESTONE                   OnEvent = "milestone"
	ON_EVENT_PAGE_BUILD                  OnEvent = "page_build"
	ON_EVENT_PROJECT                     OnEvent = "project"
	ON_EVENT_PROJECT_CARD                OnEvent = "project_card"
	ON_EVENT_PROJECT_COLUMN              OnEvent = "project_column"
	ON_EVENT_PUBLIC                      OnEvent = "public"
	ON_EVENT_PULL_REQUEST                OnEvent = "pull_request"
	ON_EVENT_PULL_REQUEST_REVIEW         OnEvent = "pull_request_review"
	ON_EVENT_PULL_REQUEST_REVIEW_COMMENT OnEvent = "pull_request_review_comment"
	ON_EVENT_PULL_REQUEST_TARGET         OnEvent = "pull_request_target"
	ON_EVENT_PUSH                        OnEvent = "push"
	ON_EVENT_REGISTRY_PACKAGE            OnEvent = "registry_package"
	ON_EVENT_RELEASE                     OnEvent = "release"
	ON_EVENT_REPOSITORY_DISPATCH         OnEvent = "repository_dispatch"
	ON_EVENT_STATUS                      OnEvent = "status"
	ON_EVENT_WATCH                       OnEvent = "watch"
	ON_EVENT_WORKFLOW_DISPATCH           OnEvent = "workflow_dispatch"
	ON_EVENT_WORKFLOW_RUN                OnEvent = "workflow_run"
)

func (o OnEventConfig) Validate() error {
	for event := range o.Events {
		v := OnEvent(event)
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
