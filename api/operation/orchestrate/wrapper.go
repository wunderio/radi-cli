package orchestrate

/**
 * A wrapper for the orchestrate operations, to make them easier
 * to run, without having to handle operations and properties.
 *
 * @NOTE this is heavily tied to docker syntax at the moment,
 * only in order to get a manageable interface, not because it
 * is any good.  We should consider dumping this if it makes sense.
 */

// A wrapper on orchestrate operations to make them easier to use
type OrchestrateWrapper interface {
	Up() error
	Down() error
}
