package config

//
type ConfigConnector interface {
	List(parent string) []string
	Get(key string) (string, bool)
	Set(key string, value string) bool
}
