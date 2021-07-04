// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package provider

import (
	"os"
	"strconv"

	"github.com/zchee/go-codecov/provider/appveyor"
	"github.com/zchee/go-codecov/provider/awscodebuild"
	"github.com/zchee/go-codecov/provider/azurepipelines"
	"github.com/zchee/go-codecov/provider/buildkite"
	"github.com/zchee/go-codecov/provider/circleci"
	"github.com/zchee/go-codecov/provider/codefresh"
	"github.com/zchee/go-codecov/provider/codeship"
	"github.com/zchee/go-codecov/provider/drone"
	"github.com/zchee/go-codecov/provider/githubactions"
	"github.com/zchee/go-codecov/provider/gitlab"
	"github.com/zchee/go-codecov/provider/jenkins"
	"github.com/zchee/go-codecov/provider/magnum"
	"github.com/zchee/go-codecov/provider/semaphore"
	"github.com/zchee/go-codecov/provider/shippable"
	"github.com/zchee/go-codecov/provider/teamcity"
	"github.com/zchee/go-codecov/provider/travis"
	"github.com/zchee/go-codecov/provider/wercher"
)

// Service represents a CI service.
type Service interface {
	// Name is service name.
	Name() string

	// Root is the current build root.
	Root() string

	// Branch is the name of the Git branch currently being built.
	Branch() string

	// Build is the number of the current build.
	Build() string

	// BuildURL is the url of the current build.
	BuildURL() string

	// Job is the number of the current job.
	Job() string

	// Tag is the current build is for a git tag.
	Tag() string

	// Commit is the hash of the last commit of the current build.
	Commit() string

	// PullRequest is the number of pull request.
	PullRequest() string

	// Slug is the owner/repo slug.
	Slug() string

	// Envs is the additional environment variables.
	Envs() map[string]string
}

// make sure each service implements the Service interface.
var (
	_ Service = (*appveyor.Service)(nil)
	_ Service = (*awscodebuild.Service)(nil)
	_ Service = (*azurepipelines.Service)(nil)
	_ Service = (*buildkite.Service)(nil)
	_ Service = (*circleci.Service)(nil)
	_ Service = (*codefresh.Service)(nil)
	_ Service = (*codeship.Service)(nil)
	_ Service = (*drone.Service)(nil)
	_ Service = (*githubactions.Service)(nil)
	_ Service = (*gitlab.Service)(nil)
	_ Service = (*jenkins.Service)(nil)
	_ Service = (*magnum.Service)(nil)
	_ Service = (*semaphore.Service)(nil)
	_ Service = (*shippable.Service)(nil)
	_ Service = (*teamcity.Service)(nil)
	_ Service = (*travis.Service)(nil)
	_ Service = (*wercher.Service)(nil)
)

func parseStringEnv(env string) bool {
	return os.Getenv(env) != ""
}

func parseBoolEnv(env string) bool {
	b, err := strconv.ParseBool(os.Getenv(env))
	return err == nil && b
}

// Detect detects the service from the parse environment variables.
func Detect() Service {
	switch {
	case parseStringEnv(DetectEnvJenkins):
		return &jenkins.Service{}

	case parseBoolEnv(DetectEnvCI):
		switch {
		case parseBoolEnv(DetectEnvTravis):
			return &travis.Service{}

		case os.Getenv(DetectEnvCIName) == string(CodeShip):
			return &codeship.Service{}

		case parseBoolEnv(DetectEnvBuildkite):
			return &buildkite.Service{}

		case parseBoolEnv(DetectEnvCircleCI):
			return &circleci.Service{}

		case parseBoolEnv(DetectEnvSemaphore):
			return &semaphore.Service{}

		case parseBoolEnv(DetectEnvAppveyor):
			return &appveyor.Service{}

		case parseStringEnv(DetectEnvWercker):
			return &wercher.Service{}

		case parseBoolEnv(DetectEnvMagnum):
			return &magnum.Service{}
		}

	case parseBoolEnv(os.Getenv(DetectEnvAwsCodeBuild)):
		return &awscodebuild.Service{}

	case parseStringEnv(DetectEnvCodeFreshBuildURL) && parseStringEnv(DetectEnvCodeFreshBuildID):
		return &codefresh.Service{}

	case os.Getenv(DetectEnvCI) == string(Drone) && parseBoolEnv(DetectEnvDrone):
		return &drone.Service{}

	case parseStringEnv(DetectEnvAzurePipelines):
		return &azurepipelines.Service{}

	case parseStringEnv(DetectEnvTeamcity):
		return &teamcity.Service{}

	case parseBoolEnv(DetectEnvShippable):
		return &shippable.Service{}

	case parseStringEnv(DetectEnvGitHubActions):
		return &githubactions.Service{}

	case parseStringEnv(DetectEnvGitLab):
		return &gitlab.Service{}
	}

	panic("unkonwn service")
}
