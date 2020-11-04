package function

import (
	"database/sql"
	"html/template"
	"net/http"
)

var db *sql.DB
var err error

func RouteIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func RouteSubmitPost(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:nadipw@tcp(127.0.0.1:3306)/northwind")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	
	if r.Method == "POST" {

		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var firstname = r.FormValue("firstname")
		var lastname = r.Form.Get("lastname")

		var data = map[string]string{"firstname": firstname, "lastname": lastname}
		
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//tugas insertkan ke database ke table employees

		return
	}

	http.Error(w, "", http.StatusBadRequest)
}