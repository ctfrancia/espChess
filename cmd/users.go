package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create user")
}
