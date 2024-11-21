package main

import (
	"github.com/go-chi/chi/v5"
	"minimal-htmgo/pages"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.StripPrefix("/public", http.FileServer(http.Dir("./assets/dist")))
	router.Handle("/public/*", fileServer)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		RenderPage(request, writer, pages.Index)
	})

	router.Get("/current-time", func(writer http.ResponseWriter, request *http.Request) {
		RenderPartial(request, writer, pages.CurrentTime)
	})

	http.ListenAndServe(":3000", router)
}
