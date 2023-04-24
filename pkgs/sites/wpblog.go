package sites

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/moqsien/free/pkgs/query"
)

var WPBlogUrl = "https://www.wenpblog.com/list/5/1.html"

type WPBlog struct {
	Result []string
	Doc    *goquery.Document
	d      *query.Downloader
}

func NewWPBlog() *WPBlog {
	return &WPBlog{
		Result: []string{},
		d:      query.NewDownloader(WPBlogUrl),
	}
}

func (that *WPBlog) getDoc() {
	resp := that.d.Get()
	if doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp)); err == nil {
		if _url := doc.Find(".list-pic-body").First().Find("h3").Find("a").First().AttrOr("href", ""); _url != "" {
			that.d = query.NewDownloader(_url)
			res := that.d.Get()
			if that.Doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(res)); err != nil {
				that.Doc = nil
			}
		}
	}
}

func (that *WPBlog) parse() {
	if that.Doc != nil {
		htmlStr := that.Doc.Find("blockquote").Text()
		for _, s := range strings.Split(htmlStr, "\n") {
			s = strings.TrimSpace(s)
			if s != "" {
				that.Result = append(that.Result, s)
			}
		}
	}
}

func (that *WPBlog) Run() []string {
	fmt.Println("Handle: ", WPBlogUrl)
	that.getDoc()
	that.parse()
	return that.Result
}
