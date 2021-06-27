package main

import (
	"net/http"
)

import (
	"html/template"
	"io/ioutil"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}





func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := loadPage(title)
	t, err := template.ParseFiles("tamplates/index.html")
	t.Execute(w, p)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("tamplates/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
