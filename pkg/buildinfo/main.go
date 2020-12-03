package buildinfo

import "time"

// GitSha of the build
var GitSha = "HEAD"

// Version contains the latest version tag
var Version = time.Now().String()
