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
	Urls   []string
}

func NewFreeNode() *FreeNode {
	return &FreeNode{
		Result: []string{},
		host:   "https://www.freefq.com",
		Urls: []string{
			"https://freefq.com/v2ray/",
			"https://freefq.com/free-ssr/",
			"https://freefq.com/free-ss/",
			"https://www.freefq.com/free-xray/",
			"https://freefq.com/freeuser/",
			"https://freefq.com/free-trojan/",
		},
	}
}

func (that *FreeNode) getDoc() {
	that.d.UseProxy()
	that.d.SetTimeout(60 * time.Second)
	resp := that.d.Get()
	if len(resp) > 0 {
		if doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp)); err == nil {
			_url := doc.Find(".news_list").Find("table").Eq(1).Find("a").First().AttrOr("href", "")
			fmt.Println("  [x] ", _url)
			if !strings.Contains(_url, "http") {
				_url, _ = url.JoinPath(that.host, _url)
			}
			that.d = query.NewDownloader(_url)
			that.d.UseProxy()
			that.d.SetTimeout(60 * time.Second)
			res := that.d.Get()
			if doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(res)); err == nil {
				if _url = doc.Find("fieldset").Find("a").AttrOr("href", ""); _url != "" {
					that.d = query.NewDownloader(_url)
					fmt.Println("    [*] ", _url)
					that.d.SetTimeout(60 * time.Second)
					that.d.UseProxy()
					r := that.d.Get()
					if that.Doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(r)); err != nil {
						that.Doc = nil
					}
				}
			}
		}
	}
}

func (that *FreeNode) parse() {
	if that.Doc != nil {
		// that.Doc.Find("font").Each(func(i int, s *goquery.Selection) {
		// 	s.Find("p").Each(func(i int, ss *goquery.Selection) {
		// 		text := ss.Text()
		// 		for _, value := range strings.Split(text, "\n") {
		// 			value = strings.TrimSpace(value)
		// 			if strings.HasPrefix(value, "vmess://") {
		// 				that.Result = append(that.Result, value)
		// 			} else if strings.HasPrefix(value, "ss://") {
		// 				that.Result = append(that.Result, value)
		// 			} else if strings.HasPrefix(value, "ssr://") {
		// 				that.Result = append(that.Result, value)
		// 			} else if strings.HasPrefix(value, "vless://") {
		// 				that.Result = append(that.Result, value)
		// 			} else if strings.HasPrefix(value, "trojan://") {
		// 				that.Result = append(that.Result, value)
		// 			}
		// 		}
		// 	})
		// })
		that.Doc.Find("p").Each(func(i int, ss *goquery.Selection) {
			text := ss.Text()
			if !strings.Contains(text, "://") {
				return
			}
			for _, value := range strings.Split(text, "\n") {
				value = strings.TrimSpace(value)
				if strings.HasPrefix(value, "vmess://") {
					that.Result = append(that.Result, value)
				} else if strings.HasPrefix(value, "ss://") {
					that.Result = append(that.Result, value)
				} else if strings.HasPrefix(value, "ssr://") {
					that.Result = append(that.Result, value)
				} else if strings.HasPrefix(value, "vless://") {
					that.Result = append(that.Result, value)
				} else if strings.HasPrefix(value, "trojan://") {
					that.Result = append(that.Result, value)
				}
			}
		})
	}
}

func (that *FreeNode) Run() []string {
	for _, pUrl := range that.Urls {
		fmt.Println("Handle: ", pUrl)
		that.d = query.NewDownloader(pUrl)
		that.getDoc()
		that.parse()
	}
	// fmt.Println(that.Result)
	return that.Result
}
