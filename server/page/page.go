// Package page serve [layout] + [views] template
package page

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	PageType string
}

func NewPage() *Page {
	return &Page{}
}

var templ = func() *template.Template {
	files := []string{"../../public/views/layout.html"}

	err := filepath.Walk("../../public/views/content/", func(path string, d os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	t := template.Must(template.ParseFiles(files...))
	return t
}()

func (p *Page) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "home"
	err := templ.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}
}

func (p *Page) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "login"
	err := templ.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}
}

func (p *Page) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "signup"
	err := templ.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) InteractionPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "interaction"
	err := templ.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
