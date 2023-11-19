package main


import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create tournament")
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Fprintln(w, params)
	fmt.Fprintln(w, "show tournament")
}
