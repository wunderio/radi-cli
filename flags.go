package main

/**
 * Parse command flags to configure the command
 */
func parseGlobalFlags(flags []string) (commandName string, globalFlags map[string]string, commandFlags []string) {
	commandName = ""

	globalFlags = map[string]string{} // start of with no flags

	global := true // start of assuming everything is a global arg
	for index := 1; index < len(flags); index++ {
		arg := flags[index]

		switch arg {
		case "-v":
			fallthrough
		case "--info":
			globalFlags["verbosity"] = "info"
		case "-vv":
			fallthrough
		case "--verbose":
			globalFlags["verbosity"] = "verbose"
		case "-vvv":
			fallthrough
		case "--debug":
			globalFlags["verbosity"] = "debug"
		case "-vvvv":
			fallthrough
		case "--staaap":
			globalFlags["verbosity"] = "staaap"


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
