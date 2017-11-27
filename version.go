package version

import "fmt"

var _VER_ = "1.0"

var (
	// Version should be updated by hand at each release
	Version string

	//will be overwritten automatically by the build system
	GitCommit string
	GoVersion string
	BuildTime string
)

func PrintVersion() {
	fmt.Printf("Version: %6s \nGit commit: %6s \nGo version: %6s \nBuild time: %6s \n",
		Version, GitCommit, GoVersion, BuildTime)
}
