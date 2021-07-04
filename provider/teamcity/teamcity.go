// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package teamcity

import (
	"os"
)

// Service represents a TeamCity CI service.
//
// See
//
// https://confluence.jetbrains.com/plugins/servlet/mobile#content/view/74847298
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "teamcity"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string { return "" }

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("BUILD_NUMBER")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("BUILD_VCS_NUMBER")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string { return "" }

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
