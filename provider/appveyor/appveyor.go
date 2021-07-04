// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package appveyor

import (
	"os"
	"path"
)

// Service represents a AppVeyor CI service.
//
// See
//
// http://www.appveyor.com/docs/environment-variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "appveyor"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("APPVEYOR_REPO_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("APPVEYOR_JOB_ID")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string {
	return path.Join(os.Getenv("APPVEYOR_ACCOUNT_NAME"), os.Getenv("APPVEYOR_PROJECT_SLUG"), os.Getenv("APPVEYOR_BUILD_VERSION"))
}

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("APPVEYOR_REPO_COMMIT")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	return os.Getenv("APPVEYOR_PULL_REQUEST_NUMBER")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return os.Getenv("APPVEYOR_REPO_NAME")
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
