package command

// A Connector for running commands
type CommandConnector interface {
	List() []string
	Get(key string) (Command, bool)
	Set(key string, com Command) bool
}
