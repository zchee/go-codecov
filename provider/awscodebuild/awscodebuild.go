// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package awscodebuild

import (
	"os"
	"path"
	"strings"
)

// Service represents a AWS Codebuild CI service.
//
// See
//
// https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-env-vars.html
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "codebuild"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	// assume "refs/heads/{branch_name}"
	return path.Base(os.Getenv("CODEBUILD_WEBHOOK_HEAD_REF"))
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("CODEBUILD_BUILD_ID")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string {
	return os.Getenv("CODEBUILD_BUILD_ID")
}

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("CODEBUILD_RESOLVED_SOURCE_VERSION")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	const prPrefix = "pr"

	version := os.Getenv("CODEBUILD_SOURCE_VERSION")
	if strings.HasPrefix(version, prPrefix) {
		return "false"
	}

	return strings.TrimPrefix(version, prPrefix)
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	// assume "https://github.com/zchee/go-codecov.git"
	repoURL := strings.TrimSuffix(os.Getenv("CODEBUILD_SOURCE_REPO_URL"), ".git")

	var (
		user string
		repo string
	)

	if idx := strings.LastIndex(repoURL, "/"); idx > -1 {
		repo = repoURL[:idx]

		repoURL = repoURL[idx:]
		if idx2 := strings.LastIndex(repoURL, "/"); idx2 > -1 {
			user = repoURL[:idx2]
		}
	}

	if user == "" || repo == "" {
		return ""
	}

	return path.Join(user, repo)
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
