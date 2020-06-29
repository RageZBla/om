package download_clients

type PivnetFileArtifact struct {
	name          string
	sha256        string
	slug          string
	releaseID     int
	productFileID int
}

func (f PivnetFileArtifact) Name() string {
	return f.name
}

func (f PivnetFileArtifact) SHA256() string {
	return f.sha256
}

type stowFileArtifact struct {
	name   string
	sha256 string
}

func (f stowFileArtifact) Name() string {
	return f.name
}

func (f stowFileArtifact) SHA256() string {
	return f.sha256
}
