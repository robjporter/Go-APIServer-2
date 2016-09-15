package controllers

import (
	"net/http"

	"../../../render"
)

func CoreHome(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/core.v1/views/home.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CoreAbout(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/core.v1/views/about.html")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "About"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Robots(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/core.v1/public/robots.txt")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Error"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FavIcon(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/core.v1/public/images/favicon.ico")
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Error"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	tmp := req.Header["Accept"]
	if len(tmp) == 1 {
		if tmp[0] == "application/json" {
			render.RenderJSON(w, http.StatusNotFound, `{ "STATUS" : "NOT FOUND"}`)
		} else if tmp[0] == "application/xml" {
			render.RenderXML(w, http.StatusNotFound, "<xml><Status>NOT FOUND</Status></xml>")
		} else {
			//templates := render.GetBaseTemplates( c )
			//templates = append( templates, "apps/core.v1/views/notfound.html" )
			templates := []string{}
			templates = append(templates, "apps/core.v1/views/notfound.html")
			err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Error"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
