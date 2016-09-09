package operation

/**
 * Configurations are abstract Operations settings which
 * are meant to be set outside of the Operation, before it
 * is executed.  Operations of a single key, from different
 * Handlers should probably try to all offer the same
 * Configuration key/types, even if the work differently.
 *
 * This file provides a configuration collection sturct,
 * and the intercface for a single Configuration along with
 * some standard data type configuration base structs.
 *
 * A Configuration consumer should either recognized the
 * Operation by it's keys, and then handle it's Configuraitons
 * as "knowns", or it should iterate through the Configurations
 * and use some user-interface to allow interaction.
 *
 * A Configuration is typically used, by overwriting it's
 * Value() reference with a different value.
 */

// A single Configuration
type Configuration interface {
	// ID returns string unique configuration Identifier
	Id() string
	// Label returns a short user readable label for the configuration
	Label() string
	// Description provides a longer multi-line string description of what the Configuration does
	Description() string

	// Value allows the retrieval and setting of unknown Typed values for the Configuration.
	Get() interface{}
	Set(interface{}) bool
}

// A set of Configurations
type Configurations struct {
	configurationMap   map[string]Configuration
	configurationOrder []string
}

func (configurations *Configurations) Add(configuration Configuration) {
	if configurations.configurationMap == nil {
		configurations.configurationMap = map[string]Configuration{}
		configurations.configurationOrder = []string{}
	}
	/**
	 * @TODO check if it already exists (by key)
	 */
	configurations.configurationMap[configuration.Id()] = configuration
	configurations.configurationOrder = append(configurations.configurationOrder, configuration.Id())
}
func (configurations *Configurations) Get(id string) (Configuration, bool) {
	configuration, ok := configurations.configurationMap[id]
	return configuration, ok
}
func (configurations *Configurations) Order() []string {
	return configurations.configurationOrder
}
