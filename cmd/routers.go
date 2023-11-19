package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	v1 = "/v1"
	v2 = "/v2"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// v1 healthcheck
	hcr := fmt.Sprintf("%s/healthcheck", v1)
	router.HandlerFunc(http.MethodGet, hcr, app.healthcheckHandler)

	ct := fmt.Sprintf("%s/tournament", v1)
	router.HandlerFunc(http.MethodPost, ct, app.createTournamentHandler)

	st := fmt.Sprintf("%s/tournament/:id", v1)
	router.HandlerFunc(http.MethodGet, st, app.showTournamentHandler)
	return router
}
