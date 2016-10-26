package bytesource

// Settings needed to make a file based bytesource API
type BytesourceFileSettings struct {
	ProjectDoesntExist bool
	ProjectRootPath    string
	UserHomePath       string
	ExecPath           string
	ConfigPaths        *Paths
}
