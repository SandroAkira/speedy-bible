package main

import (
	"log"
	"strconv"
	"strings"
)

type Book struct {
	Name       string
	Chapters   []Chapter
	ChapterMap map[string]*Chapter
}

func (b *Book) Url() string {
	return b.CleanName()
}

func (b *Book) CleanName() string {
	return strings.Replace(b.Name, " ", "+", -1)
}

func (b *Book) ShortName() string {
	if len(b.Name) < 10 {
		return b.Name
	} else {
		return b.Name[0:9]
	}
}

func (b *Book) GetChapter(s string) (c *Chapter) {
	c, ok := b.ChapterMap[s]
	if !ok {
		log.Panicf("no such chapter %s in map %v", s, b.ChapterMap)
	}
	return
}

func (b *Book) addChapter(ci, total int) {
	nextCi := 1
	if ci < total {
		nextCi = ci + 1
	}
	newChapter := Chapter{b, ci, nextCi}
	b.Chapters = append(b.Chapters, newChapter)
	cs := strconv.Itoa(ci)
	b.ChapterMap[cs] = &newChapter
}

func makeBook(name string, count int) (book Book) {
	book.Name = name
	book.Chapters = []Chapter{}
	book.ChapterMap = make(map[string]*Chapter)
	for i := 1; i <= count; i++ {
		book.addChapter(i, count)
	}
	return
}
