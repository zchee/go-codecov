// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package azurepipelines

import (
	"fmt"
	"os"
	"path"
)

// Service represents a Azure Pipelines service.
//
// See
//
// https://docs.microsoft.com/en-us/azure/devops/pipelines/build/variables?view=vsts
// https://docs.microsoft.com/en-us/azure/devops/pipelines/build/variables?view=azure-devops&viewFallbackFrom=vsts&tabs=yaml
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "azure_pipelines"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	// assume "refs/heads/{branch_name}"
	return path.Base(os.Getenv("BUILD_SOURCEBRANCH"))
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("BUILD_BUILDNUMBER")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string {
	return path.Join(
		os.Getenv("SYSTEM_TEAMFOUNDATIONSERVERURI"),
		os.Getenv("SYSTEM_TEAMPROJECT"),
		"_build",
		fmt.Sprintf("results?buildId=%s", os.Getenv("BUILD_BUILDID")),
	)
}

// Job implements service.Service.Job.
func (Service) Job() string {
	return os.Getenv("BUILD_BUILDID")
}

// Tag implements service.Service.Tag.
func (Service) Tag() string { return "" }

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	// TODO(zchee): parse below logic.
	// mc=
	// if [ -n "$pr" ] && [ "$pr" != false ];
	// then
	//   mc=$(git show --no-patch --format="%P" 2>/dev/null || echo "")
	//
	//   if [[ "$mc" =~ ^[a-z0-9]{40}[[:space:]][a-z0-9]{40}$ ]];
	//   then
	//     mc=$(echo "$mc" | cut -d' ' -f2)
	//     say "    Fixing merge commit SHA $commit -> $mc"
	//     commit=$mc
	//   fi
	// fi

	return os.Getenv("BUILD_SOURCEVERSION")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	if env := os.Getenv("SYSTEM_PULLREQUEST_PULLREQUESTID"); env != "" {
		return env
	}

	return os.Getenv("SYSTEM_PULLREQUEST_PULLREQUESTNUMBER")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string { return "" }

// Envs implements service.Service.Envs.
func (Service) Envs() map[string]string { return nil }
