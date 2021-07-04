// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package jenkins

import (
	"os"
)

// Service represents a Jenkins CI service.
//
// See
//
// https://wiki.jenkins-ci.org/display/JENKINS/Building+a+software+project
// https://wiki.jenkins-ci.org/display/JENKINS/GitHub+pull+request+builder+plugin#GitHubpullrequestbuilderplugin-EnvironmentVariables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "jenkins"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	if env := os.Getenv("WORKSPACE"); env != "" {
		return env
	}
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return getenvs("ghprbSourceBranch", "GIT_BRANCH", "BRANCH_NAME")
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
	return getenvs("ghprbActualCommit", "GIT_COMMIT")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	return getenvs("ghprbPullId", "CHANGE_ID")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string { return "" }

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }

func getenvs(envs ...string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}

	return ""
}
