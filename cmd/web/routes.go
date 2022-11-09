package main

import "net/http"

func (app *App) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/farm", app.showFarm)
	mux.HandleFunc("/calc", app.showCalcByBrand)
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", app.home)
	//mux.HandleFunc("/snippet", app.showSnippet)
	//mux.HandleFunc("/snippet/create", app.createSnippet)
	//
	//fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//
	//return mux
	return mux
}
