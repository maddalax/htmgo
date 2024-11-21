package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	fileServer := http.StripPrefix("/public", http.FileServer(http.Dir("./public")))
	router.Handle("/public/*", fileServer)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		RenderPage(request, writer, Index)
	})

	router.Get("/current-time", func(writer http.ResponseWriter, request *http.Request) {
		RenderPartial(request, writer, CurrentTime)
	})

	http.ListenAndServe(":3000", router)
}
