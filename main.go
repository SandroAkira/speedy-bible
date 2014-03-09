package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(BibleMixin())
	m.Get("/", BooksHandler)
	m.Get("/:book", ChaptersHandler)
	m.Get("/:book/:chapter", ReadHandler)
	m.Run()
}
