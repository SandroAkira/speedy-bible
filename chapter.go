package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Chapter struct {
	Book       *Book
	Number     int
	NextNumber int
}

func (c Chapter) GetBody() string {
	apiKey := os.Getenv("esv_key")
	if len(apiKey) < 1 {
		apiKey = "IP"
	}
	cleanPassage := c.Book.CleanName() + "+" + strconv.Itoa(c.Number)
	options := make([]string, 0)
	options = append(options, "include-first-verse-numbers=false")
	options = append(options, "include-verse-numbers=false")
	options = append(options, "include-short-copyright=false")
	options = append(options, "include-headings=false")
	options = append(options, "include-passage-references=false")
	options = append(options, "include-footnotes=false")
	options = append(options, "include-subheadings=false")
	options = append(options, "line-length=0")
	options = append(options, "include-passage-horizontal-lines=false")
	options = append(options, "include-heading-horizontal-lines=false")
	options = append(options, "output-format=plain-text")
	url := fmt.Sprintf("http://www.esvapi.org/v2/rest/passageQuery?key=%s&passage=%s&%s", apiKey, cleanPassage, strings.Join(options, "&"))
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url)
		return "bad get"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "bad read"
	}
	return string(body)
}
