package main

import (
	"encoding/json"
	"example/internal/storage/person"
	"example/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	personStorage := person.NewStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/person", func(w http.ResponseWriter, r *http.Request) {
		personStorage.AddUser(
			models.Person{FirstName: "Andrey"},
		)
		// получать json с описанием пользователя
		// из json получать структуру и добавлять в storage
	})
	r.Get("/person", func(w http.ResponseWriter, r *http.Request) {
		p, _ := personStorage.GetUser("12")
		// нужно получать id из параметров запроса
		//
		body, _ := json.Marshal(p)

		w.Write(body)

	})
	http.ListenAndServe(":3000", r)

}
