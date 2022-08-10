package main

import (
	"fmt"
	"runtime/debug"
)

// version contains the version string set by -ldflags.
var version string

// Version returns a version string based on how the binary was compiled.
// For binaries compiled with "make", the version set by -ldflags is returned.
// For binaries compiled with "go install", the version and commit hash from
// the embedded build information is returned if available.
func Version() string {
	if info, ok := debug.ReadBuildInfo(); ok && version == "" {
		version = info.Main.Version
		for _, s := range info.Settings {
			if s.Key == "vcs.revision" {
				version += "-" + s.Value
				break
			}
		}
	}
	return version
}

func main() {
	fmt.Println(Version())
}
