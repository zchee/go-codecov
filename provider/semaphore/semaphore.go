// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package semaphore

import (
	"os"
)

// Service represents a Semaphore CI service.
//
// See
//
// https://semaphoreapp.com/docs/available-environment-variables.html
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "semaphore"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("BRANCH_NAME")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("SEMAPHORE_BUILD_NUMBER") + "." + os.Getenv("SEMAPHORE_CURRENT_THREAD")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("REVISION")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return os.Getenv("SEMAPHORE_REPO_SLUG")
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
