// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package codefresh

import (
	"os"
)

// Service represents a CodeFresh CI service.
//
// See
//
// https://docs.codefresh.io/v1.0/docs/variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "codefresh"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("CF_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("CF_BUILD_ID")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string {
	return os.Getenv("CF_BUILD_URL")
}

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("CF_REVISION")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string { return "" }

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
