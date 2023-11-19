package main


import (
	"fmt"
	"net/http"
)

func (app *application) createTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create tournament")
}

func (app *application) showTournamentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "show tournament")
}
