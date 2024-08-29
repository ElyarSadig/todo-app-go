package auth

import "github.com/nahojer/httprouter"

type Handler interface {
	Register() httprouter.Handler
	Login() httprouter.Handler
	Logout() httprouter.Handler
}
