// all the handlers for the api are defined here
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/arnabnandikgp/greenlight/internal/data"
)

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "meowmeow",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

  err = app.writeJSON(w, http.StatusOK, envelope{"movie":movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server went into some problem with your req", http.StatusInternalServerError)
	}
}
