package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title       string
	Content     string
	MenuOptions []MenuItem
}

type MenuItem struct {
	Title    string
	URL      string
	SubMenus []MenuItem
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := fillSiteTemplate()
	data.Title = "Home"
	data.Content = "Welcome home!"
	renderTemplate(w, "template.html", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := fillSiteTemplate()
	data.Title = "About"
	data.Content = "This is the about page."
	renderTemplate(w, "template.html", data)
}

func fillSiteTemplate() PageData {
	return PageData{
		Title:   "404",
		Content: "The page you've requested cannot be found. Did you really just up this URL?",
		MenuOptions: []MenuItem{
			{Title: "Home", URL: "/"},
			{Title: "About", URL: "/about"},
			{Title: "Services", URL: "/foo", SubMenus: []MenuItem{
				{Title: "Service 1", URL: "/services/service1"},
				{Title: "Service 2", URL: "/services/service2"},
			}},
		},
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		data := fillSiteTemplate()
		renderTemplate(w, "template.html", data)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
