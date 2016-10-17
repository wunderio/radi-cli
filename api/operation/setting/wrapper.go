package setting

/**
 * An easy to use wrapper around the settings operations to provide
 * a more functional approach to retrieving settings
 */

type SettingWrapper interface {
	Get(key string) (string, error)
	Set(key, value string) error
	List(parent string) ([]string, error)
}
