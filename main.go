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
		"templates/components/menu_item.html",
		"templates/pages/home.html",
		"templates/components/footer.html",
	),
)

var menuPageTmpl = template.Must(
	template.ParseFiles(
		"templates/base_layout.html",
		"templates/components/navbar.html",
		"templates/components/menu_item.html",
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

type MenuItem struct {
	Name        string
	Description string
	Price       float64
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

		specials := []MenuItem{
			{Name: "Fettuccine Alfredo", Description: "Creamy pasta with Parmesan cheese and garlic sauce.", Price: 13.99},
			{Name: "Hawaiian Pizza", Description: "Pizza with ham, pineapple, and mozzarella cheese.", Price: 11.99},
			{Name: "Caesar Salad", Description: "Crisp romaine lettuce with Caesar dressing, croutons, and parmesan cheese.", Price: 7.49},
			{Name: "Bacon Cheeseburger", Description: "Deluxe burger with bacon, cheese, lettuce, and tomato.", Price: 10.99},
			{Name: "Vegan Pad Thai", Description: "Stir-fried rice noodles with tofu and a tangy sauce.", Price: 12.49},
			{Name: "Tiramisu", Description: "Classic Italian dessert with coffee-soaked ladyfingers and mascarpone cream.", Price: 6.99},
		}

		homePageTmpl.Execute(w, map[string]interface{}{
			"Boosted":    r.Header.Get("Hx-Boosted") == "true",
			"ActivePage": r.RequestURI,
			"NavItems":   navItems,
			"Specials":   specials,
		})
	})

	router.Get("/menu", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		menuItems := []MenuItem{
			{Name: "Spaghetti Carbonara", Description: "Delicious pasta with eggs, cheese, and pancetta.", Price: 12.99},
			{Name: "Margherita Pizza", Description: "Classic pizza with tomato sauce, mozzarella, and basil.", Price: 10.49},
			{Name: "Grilled Chicken Salad", Description: "Fresh salad with grilled chicken, mixed greens, and vinaigrette.", Price: 9.99},
			{Name: "Cheeseburger", Description: "Juicy beef patty with cheese, lettuce, and tomato.", Price: 8.99},
			{Name: "Vegetarian Stir-Fry", Description: "Stir-fried vegetables with tofu in a savory sauce.", Price: 11.49},
			{Name: "Chocolate Brownie Sundae", Description: "Warm chocolate brownie with ice cream and whipped cream.", Price: 6.99},
		}

		menuPageTmpl.Execute(w, map[string]interface{}{
			"Boosted":    r.Header.Get("Hx-Boosted") == "true",
			"ActivePage": r.RequestURI,
			"NavItems":   navItems,
			"MenuItems":  menuItems,
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
