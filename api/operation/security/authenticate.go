package security

/**
 * Operations for authenticating access to the API
 * which can come in a few forms.  This files holds
 * the Base operations, on which handlers should 
 * build.
 */
 
// A token based authentication
type OperationAuthenticateByToken struct {

}
// Id for OperationAuthenticateByToken
func (operation *OperationAuthenticateByToken) Id() string {
	return "security.authenticate.bytoken"
}
func (operation *OperationAuthenticateByToken) Label() string {
	return "Authenticate user via a Passed token"
}
