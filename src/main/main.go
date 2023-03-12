package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	id      int
	model   string
	company string
	price   int
}

var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := database.Query("select * from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	// for _, p := range products {
	// 	fmt.Println(p.id, p.model, p.company, p.price)
	// }
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, products)

}

// функция добавления данных

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = database.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)",
			model, company, price)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/create.html")
	}
}

func main() {

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)
	// http.HandleFunc("/create", CreateHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

