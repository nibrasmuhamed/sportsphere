package meta

var (
	Version   string
	BuildTime string
	CommitID  string
)

func GetVersion() string {
	if Version == "" {
		Version = "dev"
	}

	return Version
}

func GetBuildTime() string {
	if BuildTime == "" {
		BuildTime = "unknown"
	}

	return BuildTime
}

func GetCommitID() string {
	if CommitID == "" {
		CommitID = "unknown"
	}

	return CommitID
}
