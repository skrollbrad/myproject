package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Id       int
	Name     string
	Price    int
	Quantity int
	Category Category
}

type Category struct {
	NameOfCategory string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var mySignKey = []byte("clique")

var user = User{
	Username: "1",
	Password: "1",
}

const (
	connStr = "user=clique password=password dbname=postgres sslmode=disable"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("user:", u)
	checkLogin(u)
}
func checkLogin(u User) string {
	if user.Username != u.Username && user.Password != u.Password {
		fmt.Println("Error")
		err := "error"
		return err
	}
	validToken, err := GenerateJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println(err)
	}
	return validToken

}
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodES256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 1000).Unix()
	//claims["user"] = "ss"
	claims["authorized"] = true
	tokenString, err := token.SignedString(mySignKey)

	if err != nil {
		log.Fatal(err)
	}
	return tokenString, nil

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `delete from Product where id = $1`
	_, err = db.Exec(sqlStatement, 800)
	if err != nil {
		panic(err)
	}
	fmt.Println("IS DELETED")

}
func addProduct(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO Product (name, id, price,quantity) 
    VALUES ($1, $2, $3,$4)`
	_, err = db.Exec(sqlStatement, "Smith", 800, 200, 0001)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}

func getUsdFundsShares(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from Product")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity) //&p.Category

		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		w.Write([]byte(fmt.Sprintf("%v,%v,%v,%v\n", p.Id, p.Name, p.Price, p.Quantity))) //p.Category

	}

}
func main() {

	r := chi.NewRouter()
	r.Get("/", getUsdFundsShares)
	r.Post("/ad", addProduct)
	r.Delete("/dl", deleteProduct)
	r.Post("/login", login)
	http.ListenAndServe(":3030", r)
}
