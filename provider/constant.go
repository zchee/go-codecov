// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package provider

// ServiceName represents a CI service name.
type ServiceName string

// list of ServiceName.
const (
	// Appveyor name of Appveyor.
	Appveyor = ServiceName("appveyor")

	// AwsCodeBuild name of AWS CodeBuild.
	AwsCodeBuild = ServiceName("codebuild")

	// AzurePipelines name of Azure Pipelines.
	AzurePipelines = ServiceName("azure_pipelines")

	// Buildkite name of Buildkite.
	Buildkite = ServiceName("buildkite")

	// CircleCI name of CircleCI.
	CircleCI = ServiceName("circleci")

	// CodeShip name of CodeShip.
	CodeShip = ServiceName("codeship")

	// Drone name of Drone.
	Drone = ServiceName("drone")

	// GitHubActions name of Github Action.
	GitHubActions = ServiceName("github-actions")

	// GitLab name of GitLab.
	GitLab = ServiceName("gitlab")

	// Jenkins name of Jenkins.
	Jenkins = ServiceName("jenkins")

	// Magnum name of Magnum.
	Magnum = ServiceName("magnum")

	// Semaphore name of Semaphore.
	Semaphore = ServiceName("semaphore")

	// Shippable name of Shippable.
	Shippable = ServiceName("shippable")

	// TeamCity name of TeamCity.
	TeamCity = ServiceName("teamcity")

	// Travis name of Travis.
	Travis = ServiceName("travis")

	// Wercker name of Wercker.
	Wercker = ServiceName("wercker")
)

// list of DetectEnv.
const (
	// DetectEnvCI is the common environment variable for detects CI service.
	//
	// It parses whether the have "CI" environment variable.
	DetectEnvCI = "CI"

	// DetectEnvCIName is the common environment variable for detects CI service.
	//
	// It parses CI name.
	DetectEnvCIName = "CI_NAME"

	// DetectEnvCIServerName is the common environment variable for detects CI service.
	//
	// It parses CI service name.
	DetectEnvCIServerName = "CI_SERVER_NAME"

	// DetectEnvTravis is the environment variable for detects whether the Appveyor service.
	DetectEnvAppveyor = "APPVEYOR"

	// DetectEnvAwsCodeBuild is the environment variable for detects whether the AWS CodeBuild service.
	DetectEnvAwsCodeBuild = "CODEBUILD_CI"

	// DetectEnvAzurePipelines is the environment variable for detects whether the Azure Pipelines service.
	DetectEnvAzurePipelines = "SYSTEM_TEAMFOUNDATIONSERVERURI"

	// DetectEnvBuildkite is the environment variable for detects whether the Buildkite service.
	DetectEnvBuildkite = "BUILDKITE"

	// DetectEnvTravis is the environment variable for detects whether the CircleCI service.
	DetectEnvCircleCI = "CIRCLECI"

	// DetectEnvCodeFreshBuildURL is the environment variable for detects whether the CodeFresh service.
	DetectEnvCodeFreshBuildURL = "CF_BUILD_URL"

	// DetectEnvCodeFreshBuildID is the environment variable for detects whether the CodeFresh service.
	DetectEnvCodeFreshBuildID = "CF_BUILD_ID"

	// DetectEnvTravis is the environment variable for detects whether the Drone service.
	DetectEnvDrone = "DRONE"

	// DetectEnvGitHubActions is the environment variable for detects whether the GitHub Action service.
	DetectEnvGitHubActions = "GITHUB_ACTION"

	// DetectEnvGitLab is the environment variable for detects whether the GitLab Action service.
	DetectEnvGitLab = "GITLAB_CI"

	// DetectEnvJenkins is the environment variable for detects whether the Jenkins service.
	DetectEnvJenkins = "JENKINS_URL"

	// DetectEnvTravis is the environment variable for detects whether the Magnum service.
	DetectEnvMagnum = "MAGNUM"

	// DetectEnvTravis is the environment variable for detects whether the Semaphore service.
	DetectEnvSemaphore = "SEMAPHORE"

	// DetectEnvTravis is the environment variable for detects whether the Shippable service.
	DetectEnvShippable = "SHIPPABLE"

	// DetectEnvTravis is the environment variable for detects whether the Teamcity service.
	DetectEnvTeamcity = "TEAMCITY_VERSION"

	// DetectEnvTravis is the environment variable for detects whether the Travis service.
	DetectEnvTravis = "TRAVIS"

	// DetectEnvTravis is the environment variable for detects whether the Wercker service.
	DetectEnvWercker = "WERCKER_GIT_BRANCH"
)
