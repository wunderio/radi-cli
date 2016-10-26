package document

/**
 * Document operations wrapper
 */

type DocumentWrapper interface {
	Get(key string) (string, error)
	Set(key string, doc string) error
	List(parent string) ([]string, error)
}
