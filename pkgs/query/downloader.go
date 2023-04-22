package query

import (
	"bytes"
	"io"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Downloader struct {
	url      string
	colletor *colly.Collector
}

func NewDownloader(url string) *Downloader {
	return &Downloader{
		url: url,
		colletor: colly.NewCollector(func(c *colly.Collector) {
			extensions.RandomUserAgent(c)
		}),
	}
}

func (that *Downloader) Get() (body []byte) {
	that.colletor.OnResponse(func(r *colly.Response) {
		body = r.Body
	})
	that.colletor.Visit(that.url)
	return
}

func (that *Downloader) File(fPath string) {
	that.colletor.OnResponse(func(r *colly.Response) {
		if f, err := os.Create(fPath); err == nil {
			if len(r.Body) > 0 {
				io.Copy(f, bytes.NewBuffer(r.Body))
			}
			f.Close()
		}
	})
}

func (that *Downloader) UseProxy(proxy ...string) {
	var p string
	if len(proxy) > 0 && proxy[0] != "" {
		p = proxy[0]
	} else {
		p = os.Getenv("FREE_PROXY")
	}
	if p != "" {
		that.colletor.SetProxy(p)
	}
}
