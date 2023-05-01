package query

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/moqsien/free/pkgs/utils"
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
	that.colletor.Visit(that.url)
}

func (that *Downloader) GetFileWithBar(fpath string, force ...bool) {
	fc := false
	if len(force) > 0 {
		fc = force[0]
	}
	if ok, _ := utils.PathIsExist(fpath); ok && fc {
		os.RemoveAll(fpath)
	}

	// var bar *progressbar.ProgressBar

	// that.colletor.OnResponseHeaders(func(r *colly.Response) {
	// 	clen := r.Headers.Get("content-length")
	// 	clenInt, _ := strconv.Atoi(clen)
	// 	bar = progressbar.DefaultBytes(
	// 		int64(clenInt),
	// 		"downloading",
	// 	)
	// })

	that.colletor.OnScraped(func(r *colly.Response) {
		fmt.Println(string(r.Body))
		if f, err := os.Create(fpath); err == nil {
			io.Copy(f, bytes.NewBuffer(r.Body))
			f.Close()
		} else {
			fmt.Println(err)
		}
	})

	that.colletor.Visit(that.url)
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

func (that *Downloader) SetTimeout(timeout time.Duration) {
	that.colletor.SetRequestTimeout(timeout)
}

func (that *Downloader) SetHeader(key, value string) {
	that.colletor.OnRequest(func(r *colly.Request) {
		r.Headers.Set(key, value)
	})
}
