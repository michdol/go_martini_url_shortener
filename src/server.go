package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"url_shortener/db"
	"url_shortener/handlers"
	"url_shortener/serializers"
)


func main() {
	m := martini.New()

	// Middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	// Router
	r := martini.NewRouter()
	r.Get("/", handlers.HomeHandler)
	r.Get("/urls", handlers.AllUrlsHandler)
	r.Get("/urls/:pk", handlers.GetUrlHandler)
	r.Post("/add/", binding.Bind(serializers.Payload{}), handlers.AddUrlHandler)
	m.Action(r.Handle)

	// Global dependencies
	m.MapTo(db.AppDatabase, (*db.DB)(nil))

	m.RunOnAddr(":8000")
}
