package api

type API interface {
	AddHandler(handler Handler) bool
	HandlerById(id string) (Handler, bool)
	HandlerByImplementation(implementation string) (Handler, bool)
}

type BaseAPI struct {
	handlers map[string]Handler
	implements map[string]string
}

func (base *BaseAPI) AddHandler(handler Handler) bool {
	base.handlers[handler.Id()] = handler
	for _, implements := range handler.Implements() {
		base.implements[implements] = handler.Id()
	}
	return true
}
func (base *BaseAPI) HandlerById(id string) (Handler, bool) {
	handler, ok := base.handlers[id]
	return handler, ok
}
func (base *BaseAPI) HandlerByImplementation(implementation string) (Handler, bool) {
	if id, ok := base.implements[implementation]; ok {
		return base.HandlerById(id)
	}	
	return nil, false
}