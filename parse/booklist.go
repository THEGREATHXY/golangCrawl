package parse

import (
	"crawl/engine"
	"regexp"
)

const BookListRe = `<a href="([^"]+)" title="([^"]+)"`
func ParseBookList(contents []byte) engine.ParseResult {
	//fmt.Println(string(contents))
	re := regexp.MustCompile(BookListRe)
	matches := re.FindAllSubmatch(contents,-1)
	//matches := re.FindString(string(contents))

	result := engine.ParseResult{}
	//fmt.Println(matches)
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: ParseBookDetail,
		})
	}
	return result
}
