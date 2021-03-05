package actions

type OnEvent string

type OnEvents []OnEvent

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
