package bytesource

// A handler for an ordered set of keyed paths
type Paths struct {
	pathMap   map[string]string
	pathOrder []string
}

// Make sure that we have safe struct vars
func (paths *Paths) safe() {
	if paths.pathMap == nil {
		paths.pathMap = map[string]string{}
		paths.pathOrder = []string{}
	}
}

// Add a path
func (paths *Paths) Set(id string, path string) {
	paths.safe()

	// @TODO check if id already exists

	paths.pathMap[id] = path
	paths.pathOrder = append([]string{id}, paths.pathOrder...) // LIFO
}

// Retrieve a path
func (paths *Paths) Get(id string) (PathRoot, bool) {
	paths.safe()
	path, ok := paths.pathMap[id]
	return *NewPathRoot_FromStringPath(path), ok
}

// Retrieve the order of keys of the paths
func (paths *Paths) Order() []string {
	paths.safe()

	return paths.pathOrder
}
