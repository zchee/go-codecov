// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package githubactions

import (
	"net/url"
	"os"
	"path"
	"strings"
)

// Service represents a GitHub Actions service.
//
// See
//
// https://github.com/features/actions
// https://help.github.com/en/articles/virtual-environments-for-github-actions#environment-variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "github-actions"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	if env := os.Getenv("GITHUB_HEAD_REF"); env != "" {
		return env
	}

	if env := os.Getenv("GITHUB_REF"); env != "" {
		ss := strings.SplitN(env, "/", 3)
		return ss[2]
	}

	return ""
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("GITHUB_RUN_ID")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string {
	u := url.URL{
		Scheme: "http",
		Host:   "github.com",
		Path:   "/",
	}
	u.Path = path.Join(u.Path, os.Getenv("GITHUB_REPOSITORY"), "actions", "runs", os.Getenv("GITHUB_RUN_ID"))

	return u.String()
}

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("GITHUB_SHA")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	if env := os.Getenv("GITHUB_HEAD_REF"); env != "" {
		ss := strings.Split(os.Getenv("GITHUB_REF"), "/")
		return ss[len(ss)-2]
	}

	return ""
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return os.Getenv("GITHUB_REPOSITORY")
}

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
