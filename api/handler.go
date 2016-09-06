package api

const HANDLER_CONFIG_IMPLEMENTATION = "config"
const HANDLER_ENCRYPTION_IMPLEMENTATION = "encryption"
const HANDLER_SECURITY_IMPLEMENTATION = "security"

type Handler interface {
	Init()
	Validate() bool

	Id() string

	Implements() []string

	Operations() *Operations	
}

