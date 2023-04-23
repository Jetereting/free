package sites

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/moqsien/free/pkgs/query"
)

type FreeNode struct {
	Result []string
	Doc    *goquery.Document
	d      *query.Downloader
	host   string
}

func NewFreeNode() *FreeNode {
	return &FreeNode{
		Result: []string{},
		d:      query.NewDownloader("https://www.freefq.com/free-xray/"),
		host:   "https://www.freefq.com",
	}
}

func (that *FreeNode) getDoc() {
	that.d.UseProxy()
	that.d.SetTimeout(30 * time.Second)
	resp := that.d.Get()
	if len(resp) > 0 {
		if doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp)); err == nil {
			_url := doc.Find(".news_list").Find("table").Eq(1).Find("a").First().AttrOr("href", "")
			if !strings.Contains(_url, "http") {
				_url, _ = url.JoinPath(that.host, _url)
			}
			fmt.Println("1. ", _url)
			that.d = query.NewDownloader(_url)
			that.d.SetHeader("referer", "https://www.freefq.com/free-xray/")
			that.d.SetTimeout(30 * time.Second)
			res := that.d.Get()
			fmt.Println(string(res))
			if doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(res)); err == nil {
				if _url = doc.Find("fieldset").Find("a").AttrOr("href", ""); _url != "" {
					that.d = query.NewDownloader(_url)
					fmt.Println(_url)
					that.d.SetTimeout(120 * time.Second)
					r := that.d.Get()
					if that.Doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(r)); err != nil {
						that.Doc = nil
					}
				}
			}
		}
	}
}

func (that *FreeNode) Run() []string {
	that.getDoc()
	return that.Result
}
