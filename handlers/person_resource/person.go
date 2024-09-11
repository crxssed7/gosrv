package person_resource

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"gosrv/database"
	"gosrv/models"
	"gosrv/templates"

	"github.com/gorilla/mux"
)

type _resource struct {
	Person models.Person
}

type _resources struct {
	People []models.Person
}

func index(w http.ResponseWriter, r *http.Request) {
	var people []models.Person
	database.DATABASE.Find(&people)

	data := _resources{People: people}

	err := render(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func show(w http.ResponseWriter, r *http.Request) {
	person, personErr := setPerson(r)
	if personErr != nil {
		notFound(w)
		return
	}

	data := _resource{Person: person}

	err := render(w, "show", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func new(w http.ResponseWriter, r *http.Request) {
	err := render(w, "new", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func edit(w http.ResponseWriter, r *http.Request) {
	person, personErr := setPerson(r)
	if personErr != nil {
		notFound(w)
		return
	}

	data := _resource{Person: person}

	err := render(w, "edit", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	// TODO: Form error handling
	r.ParseForm()
	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	person := models.Person{
		Name: name,
		Age:  age,
	}

	if err := database.DATABASE.Create(&person).Error; err != nil {
		http.Error(w, "Could not create person", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/people", http.StatusFound)
}

func update(w http.ResponseWriter, r *http.Request) {
	person, personErr := setPerson(r)
	if personErr != nil {
		notFound(w)
		return
	}

	// Parse form data
	r.ParseForm()
	person.Name = r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	person.Age = age

	if err := database.DATABASE.Save(&person).Error; err != nil {
		http.Error(w, "Could not update person", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/people", http.StatusFound)
}

func destroy(w http.ResponseWriter, r *http.Request) {
	person, personErr := setPerson(r)
	if personErr != nil {
		notFound(w)
		return
	}

	if err := database.DATABASE.Delete(&person).Error; err != nil {
		http.Error(w, "Could not delete person", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/people", http.StatusFound)
}

// Either show, update, or destroy a resource
func resource(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		show(w, r)
	case http.MethodPost:
		actualMethod := r.FormValue("_method")
		if actualMethod == "" || actualMethod == http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		r.Method = actualMethod
		resource(w, r)
	case http.MethodPut:
		update(w, r)
	case http.MethodDelete:
		destroy(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func setPerson(r *http.Request) (models.Person, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	if err := database.DATABASE.First(&person, id).Error; err != nil {
		return person, err
	}

	return person, nil
}

func notFound(w http.ResponseWriter) {
	http.Error(w, "Person not found", http.StatusNotFound)
}

func render(w http.ResponseWriter, name string, data any) error {
	templateName := fmt.Sprintf("people/%s.html", name)
	tmpl := template.Must(template.ParseFS(templates.TEMPLATE_FILES, "layouts/application.html", templateName))
	return tmpl.Execute(w, data)
}

func Register(router *mux.Router) {
	router.HandleFunc("/people/new", new).Methods("GET")
	router.HandleFunc("/people", index).Methods("GET")
	router.HandleFunc("/people/{id}", resource)
	router.HandleFunc("/people/{id}/edit", edit).Methods("GET")
	router.HandleFunc("/people", create).Methods("POST")
}
