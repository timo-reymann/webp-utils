package buildinfo

import "time"

// GitSha of the build
var GitSha = "unknown"

// Version contains the latest version tag
var Version = time.Now().UTC().Format("2006-01-02T15:04:05")
