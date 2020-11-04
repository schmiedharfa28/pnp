package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	st "pnp/HtmlPage/common"

	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB
var err error

func Index(w http.ResponseWriter, r *http.Request) {

	var customers []st.Customer


	sql := `SELECT
				CustomerID,
				IFNULL(CompanyName,''),
				IFNULL(ContactName,'') ContactName,
				IFNULL(ContactTitle,'') ContactTitle,
				IFNULL(Address,'') Address,
				IFNULL(City,'') City,
				IFNULL(Country,'') Country,
				IFNULL(Phone,'') Phone ,
				IFNULL(PostalCode,'') PostalCode
			FROM customers ORDER BY CustomerID`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var customer st.Customer
		err := result.Scan(&customer.CustomerID, &customer.CompanyName, &customer.ContactName,
			&customer.ContactTitle, &customer.Address, &customer.City, &customer.Country,
			&customer.Phone, &customer.PostalCode)

		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	t, err := template.ParseFiles("index.html")
	t.Execute(w, customers)

	if err != nil {
		panic(err.Error())
	}

}

func main() {
							   //<user>:<passwprd>@tcp<IP address>/<Password>
	db, err = sql.Open("mysql", "root:nadipw@tcp(127.0.0.1:3306)/northwind")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	log.Println("Server started on: http://localhost:8081")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8081", nil)
	

}



