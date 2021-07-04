// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package circleci

import (
	"os"
	"path"
)

// Service represents a CircleCI CI service.
//
// See
//
// https://circleci.com/docs/environment-variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "circleci"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("CIRCLE_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("CIRCLE_BUILD_NUM") + "." + os.Getenv("CIRCLE_NODE_INDEX")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string {
	return os.Getenv("CIRCLE_BUILD_NUM") + "." + os.Getenv("CIRCLE_NODE_INDEX")
}

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("CIRCLE_SHA1")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	return os.Getenv("CIRCLE_PR_NUMBER")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return path.Join(os.Getenv("CIRCLE_PROJECT_USERNAME"), os.Getenv("CIRCLE_PROJECT_REPONAME"))
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
