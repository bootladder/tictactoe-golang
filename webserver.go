package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/browser"
)

func NOT___main() {

	router := httprouter.New()
	router.ServeFiles("/static/*filepath",
		http.Dir("public"))
	router.ServeFiles("/images/*filepath",
		http.Dir("public/images"))

	router.POST("/api/", createPostHandler(""))

	browser.OpenURL("http://localhost:9000/static/index.html")
	http.ListenAndServe(":9000", router)
}

func postNoteHandler(w http.ResponseWriter, r *http.Request) {

	writePostResponse(w)
}
func writePostResponse(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte{0, 0, 0})
}
func createPostHandler(msg string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		postNoteHandler(w, r)
	}
}
