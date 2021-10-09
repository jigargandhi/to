package version

var (
	version = "development"
	commit  = "edge"
)

func Version() string {
	return version
}

func Commit() string {
	return commit
}
