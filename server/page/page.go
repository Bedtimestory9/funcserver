// Package page serve [layout] + [views] template
package page

import (
	"funcserver/server/misc"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	Login
	TemplateType
	templates *template.Template
}

type TemplateType struct {
	TemplateTypeName string
}

func NewPage() *Page {
	return &Page{
		templates: setupTemplates(),
	}
}

func setupTemplates() *template.Template {
	rootDir := misc.GetCmdDir()

	files := []string{rootDir + "/public/views/layout.html"}

	err := filepath.Walk(rootDir+"/public/views/content/", func(path string, d os.FileInfo, err error) error {
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
	// NOTE: This is to pass in template data
	p.TemplateTypeName = "home"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}
}

func (p *Page) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	p.TemplateTypeName = "login"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		log.Println(err)
	}

}

func (p *Page) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	p.TemplateTypeName = "signup"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *Page) InteractionPageHandler(w http.ResponseWriter, r *http.Request) {
	p.TemplateTypeName = "interaction"
	err := p.templates.ExecuteTemplate(w, "layout", &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
