package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/moqsien/free/pkgs/query"
	"github.com/moqsien/free/pkgs/utils"
)

type CFMem struct {
	Result []string
	Doc    *goquery.Document
	d      *query.Downloader
}

func NewCFMem() *CFMem {
	return &CFMem{
		Result: []string{},
		d:      query.NewDownloader("https://www.cfmem.com/search/label/free"),
	}
}

func (that *CFMem) getDoc() {
	resp := that.d.Get()
	// fmt.Println(string(resp))
	if len(resp) > 0 {
		if doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp)); err == nil {
			if _url := doc.Find("article").First().Find("a").First().AttrOr("href", ""); _url != "" {
				that.d = query.NewDownloader(_url)
				res := that.d.Get()
				if that.Doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(res)); err != nil {
					that.Doc = nil
				}
			}
		}
	}
}

func (that *CFMem) parse() {
	if that.Doc != nil {
		var encryptStr string
		that.Doc.Find("pre").Find("span").Find("span").Each(func(i int, s *goquery.Selection) {
			if str := s.Text(); str != "" {
				encryptStr += str
			}
		})
		htmlStr := utils.DecodeBase64(encryptStr)
		for _, s := range strings.Split(htmlStr, "\n") {
			s = strings.TrimSpace(s)
			if s != "" {
				that.Result = append(that.Result, s)
			}
		}
	}
}

func (that *CFMem) Run() []string {
	that.getDoc()
	that.parse()
	return that.Result
}
