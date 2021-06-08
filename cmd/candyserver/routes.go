package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/candy", app.showCandy)
	mux.HandleFunc("/candy/create", app.storeCandy)

	fileServer := http.FileServer(http.Dir("./assets/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
