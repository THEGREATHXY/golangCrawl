package parse

import (
	"crawl/engine"
	"crawl/model"
	"regexp"
)
var bookNameRe = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var authorRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]+)[\d\D]*?<br>`)
var bookPageRe = regexp.MustCompile(`<span class="pl">页数:</span> [\d\D]*?([^<]+)[\d\D]*?<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span> [\d\D]*?([^<]+)[\d\D]*?<br>`)
//<span class="pl">定价:</span> 20.00元<br>
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " [\d\D]*?property="v:average"> ([^<]+) </strong>`)
var introduceRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p>[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetailSimple(contents []byte) engine.ParseResult {
	bookDetail := model.BookDetail{}
	bookDetail.BookName = ExtraString(contents,bookNameRe)
	bookDetail.Author = ExtraString(contents,authorRe)
	bookDetail.Publisher = ExtraString(contents,publisherRe)
	bookDetail.BookPages = ExtraString(contents,bookPageRe)
	bookDetail.Price = ExtraString(contents,priceRe)
	bookDetail.Score = ExtraString(contents,scoreRe)
	bookDetail.Introduce = ExtraString(contents,introduceRe)

	result := engine.ParseResult{
		Items: []interface{}{bookDetail},
	}
	return result
}

func ParseBookDetail(contents []byte, bookName string) engine.ParseResult {
	bookDetail := model.BookDetail{}
	bookDetail.BookName = bookName
	bookDetail.Author = ExtraString(contents,authorRe)
	bookDetail.Publisher = ExtraString(contents,publisherRe)
	bookDetail.BookPages = ExtraString(contents,bookPageRe)
	bookDetail.Price = ExtraString(contents,priceRe)
	bookDetail.Score = ExtraString(contents,scoreRe)
	bookDetail.Introduce = ExtraString(contents,introduceRe)

	result := engine.ParseResult{
		Items: []interface{}{bookDetail},
	}
	return result
}
func ExtraString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	str := ""
	for i := 1; i < len(match); i++ {
		str += string(match[i])
	}
	return str
}