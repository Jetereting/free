package sites

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/moqsien/free/pkgs/query"
)

var MFFQUrl string = "https://gitlab.com/mianfeifq/share/-/blob/master/README.md?format=json&viewer=rich"

type MianfeiFQ struct {
	Result []string
	Doc    *goquery.Document
	d      *query.Downloader
}

func NewMianfieFQ() *MianfeiFQ {
	return &MianfeiFQ{
		Result: []string{},
		d:      query.NewDownloader(MFFQUrl),
	}
}

func (that *MianfeiFQ) getDoc() {
	var (
		err error
	)
	result := that.d.Get()
	j := gjson.New(result)
	htmlStr := j.GetString("html")
	that.Doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(htmlStr)))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (that *MianfeiFQ) parse() {
	if that.Doc != nil {
		that.Doc.Find("code").Find("span").Each(func(i int, s *goquery.Selection) {
			if str := s.Text(); str != "" {
				that.Result = append(that.Result, str)
			}
		})
	}
}

func (that *MianfeiFQ) Run() []string {
	fmt.Println("Handle: ", MFFQUrl)
	that.getDoc()
	that.parse()
	return that.Result
}
