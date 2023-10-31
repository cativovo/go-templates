package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var homePageTmpl = template.Must(
	template.ParseFiles(
		"templates/base_layout.html",
		"templates/components/navbar.html",
		"templates/pages/home.html",
		"templates/components/footer.html",
	),
)

var menuPageTmpl = template.Must(
	template.ParseFiles(
		"templates/base_layout.html",
		"templates/components/navbar.html",
		"templates/pages/menu.html",
		"templates/components/footer.html",
	),
)

var contactPageTmpl = template.Must(
	template.ParseFiles(
		"templates/base_layout.html",
		"templates/components/navbar.html",
		"templates/pages/contact.html",
		"templates/components/footer.html",
	),
)

type NavItem struct {
	Uri   string
	Label string
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Compress(5, "text/html", "text/css"))

	navItems := []NavItem{
		{
			Uri:   "/",
			Label: "Home",
		},
		{
			Uri:   "/menu",
			Label: "Menu",
		},
		{
			Uri:   "/contact",
			Label: "Contact",
		},
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		homePageTmpl.Execute(w, map[string]interface{}{
			"Boosted":    r.Header.Get("Hx-Boosted") == "true",
			"ActivePage": r.RequestURI,
			"NavItems":   navItems,
		})
	})

	router.Get("/menu", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		menuPageTmpl.Execute(w, map[string]interface{}{
			"Boosted":    r.Header.Get("Hx-Boosted") == "true",
			"ActivePage": r.RequestURI,
			"NavItems":   navItems,
		})
	})

	router.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		contactPageTmpl.Execute(w, map[string]interface{}{
			"Boosted":    r.Header.Get("Hx-Boosted") == "true",
			"ActivePage": r.RequestURI,
			"NavItems":   navItems,
		})
	})

	http.ListenAndServe("127.0.0.1:4000", router)
}
