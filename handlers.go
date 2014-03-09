package main

import (
	"log"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func BibleMixin() martini.Handler {
	bible := makeBible()

	return func(c martini.Context) {
		c.Map(&bible)
		c.Next()
	}
}

func BooksHandler(bible *Bible, r render.Render, p martini.Params) {
	r.HTML(200, "books", bible)
}

func ChaptersHandler(bible *Bible, r render.Render, p martini.Params) {
	book := bible.BookMap[p["book"]]
	r.HTML(200, "chapters", book)
}

type ReadingContext struct {
	Chapter Chapter
	Body    string
}

func ReadHandler(bible *Bible, r render.Render, p martini.Params, l *log.Logger) {
	book := bible.GetBook(p["book"])
	chapter := book.GetChapter(p["chapter"])
	body := chapter.GetBody()
	rc := ReadingContext{*chapter, body}
	r.HTML(200, "read", rc)
}
