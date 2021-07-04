// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package shippable

import (
	"os"
)

// Service represents a Shippable CI service.
//
// See
//
// http://docs.shippable.com/en/latest/config.html#common-environment-variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "shippable"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("BUILD_NUMBER")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string {
	return os.Getenv("BUILD_URL")
}

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("COMMIT")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	return os.Getenv("PULL_REQUEST")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return os.Getenv("REPO_NAME")
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
