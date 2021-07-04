// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package drone

import (
	"os"
)

// Service represents a Drone CI service.
//
// See
//
// http://docs.drone.io/env.html
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "drone.io"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	if env := os.Getenv("DRONE_BUILD_DIR"); env != "" {
		return env
	}
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("DRONE_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("DRONE_BUILD_NUMBER")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string {
	return os.Getenv("DRONE_BUILD_LINK")
}

// Job implements service.Service.Job.
func (Service) Job() string { return "" }

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string { return "" }

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string { return "" }

// Slug implements service.Service.Slug.
func (Service) Slug() string { return "" }

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
