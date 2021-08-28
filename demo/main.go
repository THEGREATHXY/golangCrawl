package main

import (
	"fmt"
	"regexp"
)
func main() {
	str := "<a href=\"https://book.douban.com/subject/4913064/\" title=\"活着\" onclick=\"moreurl(this,{i:'0',query:'',subject_id:'4913064',from:'book_subject_search'})\">\n\n    活着"
	const BookListRe = `<a href="([^"]+)" title="([^"]+)" onclick="moreurl`
	re := regexp.MustCompile(BookListRe)

	result := re.FindString(str)
	fmt.Println(result)
}
