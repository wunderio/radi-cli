package operation

// A set of Configurations
type Configurations struct {
	configurationMap map[string]Configuration
	configurationOrder []string
}
func (configurations *Configurations) AddConfiguration(configuration Configuration) {
	/**
	 * @TODO check if it already exists (by key)
	 */
	configurations.configurationMap[configuration.Id()] = configuration
	configurations.configurationOrder = append( configurations.configurationOrder, configuration.Id() )
}
func (configurations *Configurations) Configuration(id string) (Configuration, bool) {
	configuration, ok := configurations.configurationMap[id]
	return configuration, ok
}
func (configurations *Configurations) ConfigurationOrder() []string {
	return configurations.configurationOrder
}

// A single Configuration
type Configuration interface {
	Id()  string
	Label() string
	Description() string

	Value() interface{}
}


type BaseConfiguration struct {
	id string
	label string
	description string
}
func (config *BaseConfiguration) Id() string {
	return config.id
}
func (config *BaseConfiguration) Label() string {
	return config.label
}
func (config *BaseConfiguration) Description() string {
	return config.description
}


type StringConfiguration struct {
	value string
}
func (config *StringConfiguration) Value() interface{} {
	return interface{}(&config.value)
}

type BytesArrayConfiguration struct {
	value []byte
}
func (config *BytesArrayConfiguration) Value() interface{} {
	return interface{}(&config.value)
}


type BooleanConfiguration struct {
	value bool
}
func (config *BooleanConfiguration) Value() interface{} {
	return interface{}(&config.value)
}
