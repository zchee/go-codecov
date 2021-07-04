// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package gitlab

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Service represents a GitLab CI service.
//
// See
//
// https://docs.gitlab.com/ce/ci/variables/
// https://docs.gitlab.com/ee/ci/variables/predefined_variables.html
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "gitlab"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	projectDir := os.Getenv("CI_PROJECT_DIR")

	switch {
	case runtime.GOOS == "windows":
		return projectDir

	case projectDir[0] == '/': // handle os.getenv("CI_PROJECT_DIR", "").startswith("/") in codecov-python
		return projectDir

	default:
		return filepath.Join(os.Getenv("HOME"), projectDir)
	}
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return getenvs("CI_COMMIT_REF_NAME", "CI_BUILD_REF_NAME")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return getenvs("CI_JOB_ID", "CI_BUILD_ID")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return getenvs("CI_COMMIT_SHA", "CI_BUILD_REF")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	env := getenvs("CI_REPOSITORY_URL", "CI_BUILD_REPO")
	ss := strings.SplitN(env, "/", 3)
	if len(ss) > 2 {
		return ""
	}
	return ss[0] + ss[1][:4]
}

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
