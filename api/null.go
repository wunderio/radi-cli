package api

func MakeNullAPI() API {
	nAPI := BaseAPI{}

	nAPI.AddHandler( Handler(&NullHandler{}) )

	return API(&nAPI)
}

type NullHandler struct {

}
func (handler *NullHandler) Id() string {
	return "null"
}
func (handler *NullHandler) Init() {
	
}
func (handler *NullHandler) Validate() bool {
	return true
}

func (handler *NullHandler) Implements() []string {
	return []string{
		HANDLER_CONFIG_IMPLEMENTATION,
		HANDLER_ENCRYPTION_IMPLEMENTATION,
		HANDLER_SECURITY_IMPLEMENTATION,
	}
}

func (handler *NullHandler) Operations() *Operations {
	operations := Operations{}

	return &operations
}