package command

type CommandWrapper interface {
	Get(key string) (Command, error)
	List(parent string) ([]string, error)
}
