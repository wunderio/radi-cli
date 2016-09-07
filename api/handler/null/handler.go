package null

import (
	"github.com/james-nesbitt/wundertools-go/api/operation"
)

// NullHandler Constructor, doesn't do much preprocessing really
func NewNullHandler() *NullHandler {
	nullHandler := NullHandler{}
	return &nullHandler
}

// NullHandler is a handler implementation that provides many core operations, but does very little (but is safe to use)
type NullHandler struct {}
// [Handler.]Id returns a string ID for the handler
func (handler *NullHandler) Id() string {
	return "null"
}
// [Handler.]Init tells the NullHandler to process it's configurations
func (handler *NullHandler) Init() {}
// [Handler.]Validate always returns true, as this handler never fails
func (handler *NullHandler) Validate() bool {
	return true
}
// [Handler.]Operations returns an Operations list of a number of different Null operations
func (handler *NullHandler) Operations() *operation.Operations {
	operations := operation.Operations{}
	return &operations
}
