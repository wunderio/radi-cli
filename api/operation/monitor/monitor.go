package monitor

/**
 * Monitor operations are meant to provide ways for the CLI/service
 * to either retrieve information about the app state, or to connect
 * to an external monitoring service.
 *
 * Monitoring operations that return information will typically do so
 * by either offering to output state information (to an io.Writer,)
 * which should be done by offering an io.Writer variable in the operation
 * Configuration.
 * Connecting operations will need to provide a means to
 * connect in a custom manner, or by using a common messaging service
 * to connect.
 *
 * @NOTE this is currently the least though out set of operations
 */
 