package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/webservice-golang-api/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	temp.ExecuteTemplate(w, "Index", users)
}

func StoreView(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Create", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Panic("Method POST required")
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	models.StoreUser(name, email, phone)

	http.Redirect(w, r, "/", http.StatusFound)
}

func UpdateView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		log.Panic(err.Error())
	}

	userData := models.UpdateView(idConv)

	temp.ExecuteTemplate(w, "Edit", userData)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Panic("Error: PUT method required")
	}

	id := r.FormValue("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		log.Panic(err.Error())
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	models.Update(idConv, name, email, phone)

	http.Redirect(w, r, "/", http.StatusFound)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		log.Panic(err.Error())
	}

	models.Destroy(newId)

	http.Redirect(w, r, "/", http.StatusFound)
}
