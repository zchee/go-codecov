// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package wercher

import (
	"os"
	"path"
)

// Service represents a Wercker CI service.
//
// See
//
// http://devcenter.wercker.com/articles/steps/variables.html
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "wercker"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("WERCKER_GIT_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("WERCKER_MAIN_PIPELINE_STARTED")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("WERCKER_GIT_COMMIT")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return path.Join(os.Getenv("WERCKER_GIT_OWNER"), os.Getenv("WERCKER_GIT_REPOSITORY"))
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
