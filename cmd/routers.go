package main

import (
	"github.com/julienschmidt/httprouter"
)

func (a application) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("GET", "/hello", a.helloHandler)
	return router
}
