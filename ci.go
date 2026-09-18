package hxhelpers

import (
	"os"
	"sync"
)

var IsRunningInCI = sync.OnceValue(
	func() bool {
		for _, key := range [...]string{
			"CI",
			"GITHUB_ACTIONS",
			"GITLAB_CI",
			"TRAVIS",
			"CIRCLECI",
			"JENKINS_URL",
			"BUILDKITE",
			"TEAMCITY_VERSION",
		} {
			if len(os.Getenv(key)) > 0 {
				return true
			}
		}

		return false
	},
)
