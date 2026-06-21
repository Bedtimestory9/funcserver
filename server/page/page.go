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
	PageType  string
	templates *template.Template
}

const ProjectRoot = "/Users/lawrence/Projects/Learn WebDev/functional-server/"

func NewPage() *Page {
	return &Page{
		templates: setupTemplates(),
	}
}

func setupTemplates() *template.Template {
	files := []string{ProjectRoot + "/public/views/layout.html"}

	err := filepath.Walk(ProjectRoot+"/public/views/content/", func(path string, d os.FileInfo, err error) error {
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
}

func (p *Page) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "home"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}
}

func (p *Page) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "login"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}
}

func (p *Page) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "signup"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) InteractionPageHandler(w http.ResponseWriter, r *http.Request) {
	p.PageType = "interaction"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
