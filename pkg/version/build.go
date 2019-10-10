package version

var (
	// commitFromGit is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	commitSHA string
	// versionFromGit is a constant representing the version tag that
	// generated this build. It should be set during build via -ldflags.
	tag string
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	date string

	version string
)

// Info holds build information
type Info struct {
	Version   string `json:"Version"`
	BuildDate string `json:"BuildDate"`
	GitCommit string `json:"GitCommit"`
	GitTag    string `json:"GitTag"`
}

// Get creates and initialized Info object
func Get() Info {
	return Info{
		Version:   version,
		BuildDate: date,
		GitCommit: commitSHA,
		GitTag:    tag,
	}
}
