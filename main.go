package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Product struct {
	Id       string
	Name     string
	Price    int
	Quantity int
	Category Category
}
type Category struct {
	NameOfCategory string
}

func main() {

	category := Category{NameOfCategory: "Овощи"}

	product := Product{Id: "10", Name: "Морковь", Price: 100, Quantity: 10, Category: category}

	s, _ := json.Marshal(product)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		time.Sleep(time.Second * 1)
		w.Write([]byte(s))
	})
	r.Post("/add", func(w http.ResponseWriter, r *http.Request) {

		var p Product
		json.NewDecoder(r.Body).Decode(&p)
		fmt.Println(p)
		w.Write([]byte("Added"))
	})
	r.Put("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		var p Product
		p.Id = chi.URLParam(r, "userId")
		w.Write([]byte("Update user: " + p.Id))
	})

	http.ListenAndServe(":8080", r)
}
