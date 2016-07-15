package main

/**
 * @TODO This will likely soon be replaced with github.com/urfave/cli
 */

/**
 * Parse command flags to configure the command
 */
func parseGlobalFlags(flags []string) (commandName string, globalFlags map[string]string, commandFlags []string) {
	commandName = ""

	globalFlags = map[string]string{} // start of with no flags

	global := true // start of assuming everything is a global arg
	for index := 1; index < len(flags); index++ {
		arg := flags[index]

		/**
		 * Verbosity flags here should match the 
		 * github.com/sirupsen/logrus string equivalent
		 * of the levels.  we don't actually set any
		 * log level here to keep it abstract, and to
		 * allow later overrides
		 */

		switch arg {
		case "-q":
			fallthrough
		case "--quiet":
			globalFlags["verbosity"] = "error"
		case "-qq":
			fallthrough
		case "--very-quiet":
			globalFlags["verbosity"] = "fatal"
		case "-v":
			fallthrough
		case "--debug":
			globalFlags["verbosity"] = "debug"


		default:

			commandName = arg
			index++

			global = false

		}

		// all remaining flags are local
		if !global {
			commandFlags = flags[index:]
			break
		}
	}

	// return is handles via named arguments
	return
}
