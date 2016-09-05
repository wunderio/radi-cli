package version

var (
	// VERSION should be updated by hand at each release
	VERSION = "0.2.0-dev"

	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"

	// BUILDTIME will be overwritten automatically by the build system
	BUILDTIME = ""

	// SHOWWARNING might be overwritten by the build system to not show the warning
	SHOWWARNING = "true"
)