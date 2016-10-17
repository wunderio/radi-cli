package config

/**
 * The config wrappers are utility structs used to connect to
 * specific configs as an abstract, so that single connections
 * can provide multiple operations, and can optimize, and
 * make config access safe.
 */
type ConfigWrapper interface {
	Get(key string) (ConfigScopedValues, error)
	Set(key string, values ConfigScopedValues) error
	List(parent string) ([]string, error)
}

type ConfigScopedValues struct {
	configMap map[string]ConfigScopedValue
	order     []string
}

type ConfigScopedValue []byte

// Save JIT initializer
func (values *ConfigScopedValues) safe() {
	if values.configMap == nil {
		values.configMap = map[string]ConfigScopedValue{}
		values.order = []string{}
	}
}

// Get a FileSource from the set
func (values *ConfigScopedValues) Get(key string) (ConfigScopedValue, bool) {
	values.safe()

	value, found := values.configMap[key]
	return value, found
}

// Add a FileSource to the set
func (values *ConfigScopedValues) Add(key string, source ConfigScopedValue) {
	values.safe()

	if _, found := values.configMap[key]; !found {
		values.order = append(values.order, key)
	}
	values.configMap[key] = source
}

// Get the key order for the set
func (values *ConfigScopedValues) Order() []string {
	values.safe()
	return values.order
}
