package build

import "fmt"

func ShowBuildParams(buildVersion string, buildDate string, buildCommit string) {
	na := "N/A"
	if buildVersion == "" {
		buildVersion = na
	}
	if buildDate == "" {
		buildDate = na
	}
	if buildCommit == "" {
		buildCommit = na
	}
	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)
}
