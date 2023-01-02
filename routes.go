package easygo

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (easyGo *EasyGo) routes() http.Handler {
	mux := chi.NewRouter()

	//Appends request id to requests
	mux.Use(middleware.RequestID)
	//Gets real IP of request
	mux.Use(middleware.RealIP)

	//Writes requests to console
	if easyGo.Debug {
		mux.Use(middleware.Logger)
	}
	//Recovers from panics
	mux.Use(middleware.Recoverer)

	//------Routes------
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "EasyGo")
	})
	return mux
}
