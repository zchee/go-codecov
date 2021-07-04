// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package travis

import (
	"os"
)

// Service represents a Travis CI service.
//
// See
//
// http://docs.travis-ci.com/user/environment-variables/#Default-Environment-Variables
type Service struct{}

// Name implements service.Service.Name.
func (Service) Name() string {
	return "travis"
}

// Root implements service.Service.Root.
func (Service) Root() string {
	if env := os.Getenv("TRAVIS_BUILD_DIR"); env != "" {
		return env
	}

	cwd, _ := os.Getwd()
	return cwd
}

// Branch implements service.Service.Branch.
func (Service) Branch() string {
	return os.Getenv("TRAVIS_BRANCH")
}

// Build implements service.Service.Build.
func (Service) Build() string {
	return os.Getenv("TRAVIS_JOB_NUMBER")
}

// BuildURL implements service.Service.BuildURL.
func (Service) BuildURL() string { return "" }

// Job implements service.Service.Job.
func (Service) Job() string {
	return os.Getenv("TRAVIS_JOB_ID")
}

// Tag implements service.Service.Tag.
func (Service) Tag() string {
	return os.Getenv("TRAVIS_TAG")
}

// Commit implements service.Service.Commit.
func (Service) Commit() string {
	return os.Getenv("TRAVIS_COMMIT")
}

// PullRequest implements service.Service.PullRequest.
func (Service) PullRequest() string {
	return os.Getenv("TRAVIS_PULL_REQUEST")
}

// Slug implements service.Service.Slug.
func (Service) Slug() string {
	return os.Getenv("TRAVIS_REPO_SLUG")
}

func (Service) langVersion() (string, string) {
	const (
		langEnvPrefix = "TRAVIS_"
		langEnvSuffix = "_VERSION"
	)
	var langs = []string{"DART", "GO", "HEXE", "JDK", "JULIA", "NODE", "OTP", "XCODE", "PERL", "PHP", "PYTHON", "R", "RUBY", "RUST", "SCALA"}

	for _, lang := range langs {
		k := langEnvPrefix + lang + langEnvSuffix
		if v := os.Getenv(k); v != "" {
			return k, v
		}
	}

	return "", ""
}

// Envs implements service.Service.Envs.
func (t Service) Envs() (m map[string]string) {
	m = make(map[string]string)

	if osName := os.Getenv("TRAVIS_OS_NAME"); osName != "" {
		m["TRAVIS_OS_NAME"] = osName
	}

	if langEnv, langVersion := t.langVersion(); langVersion != "" {
		m[langEnv] = langVersion
	}

	return m
}
