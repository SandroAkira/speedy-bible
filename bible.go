package main

type Bible struct {
	Books   []Book
	BookMap map[string]*Book
}

func (b *Bible) addBook(name string, count int) {
	newBook := makeBook(name, count)
	b.Books = append(b.Books, newBook)
	b.BookMap[newBook.CleanName()] = &newBook
}

func (b *Bible) GetBook(s string) (book *Book) {
	book, ok := b.BookMap[s]
	if !ok {
		panic("no book found")
	}
	return
}

func makeBible() (bible Bible) {
	bible.Books = []Book{}
	bible.BookMap = make(map[string]*Book)
	bible.addBook("Matthew", 28)
	bible.addBook("Mark", 16)
	bible.addBook("Luke", 24)
	bible.addBook("John", 21)
	bible.addBook("Acts", 28)
	bible.addBook("Romans", 16)
	bible.addBook("1 Corinthians", 16)
	bible.addBook("2 Corinthians", 13)
	bible.addBook("Galatians", 6)
	bible.addBook("Ephesians", 6)
	bible.addBook("Philippians", 4)
	bible.addBook("Colossians", 4)
	bible.addBook("1 Thessalonians", 5)
	bible.addBook("2 Thessalonians", 3)
	bible.addBook("1 Timothy", 6)
	bible.addBook("2 Timothy", 4)
	bible.addBook("Titus", 3)
	bible.addBook("Philemon", 1)
	bible.addBook("Hebrews", 13)
	bible.addBook("James", 5)
	bible.addBook("1 Peter", 5)
	bible.addBook("2 Peter", 3)
	bible.addBook("1 John", 5)
	bible.addBook("2 John", 1)
	bible.addBook("3 John", 1)
	bible.addBook("Jude", 1)
	bible.addBook("Revelation", 22)
	return
}
