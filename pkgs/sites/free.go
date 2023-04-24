package sites

import (
	"fmt"
	"strings"
	"time"

	"github.com/moqsien/free/pkgs/query"
	"github.com/moqsien/free/pkgs/utils"
)

type PUrl struct {
	Url         string
	EnableProxy bool
}

type SomeFree struct {
	Urls   []PUrl
	Result []string
	d      *query.Downloader
}

func NewSomeFree() *SomeFree {
	r := &SomeFree{
		Result: []string{},
	}
	r.setUrls()
	return r
}

func (that *SomeFree) setUrls() {
	that.Urls = []PUrl{
		{"https://raw.fastgit.org/freefq/free/master/v2", false},
		{"https://sub.nicevpn.top/long", false},
		{fmt.Sprintf("https://clashnode.com/wp-content/uploads/%s.txt", time.Now().Format("2006/01/20060102")), false},
		{fmt.Sprintf("https://nodefree.org/dy/%s.txt", time.Now().Format("2006/01/20060102")), false},
		{"https://api.subcloud.xyz/sub?target=v2ray&url=https%3A%2F%2Fcdn.jsdelivr.net%2Fgh%2Fzyzmzyz%2Ffree-nodes%40master%2FClash.yml&insert=false", false},
		{"https://api.subcloud.xyz/sub?target=v2ray&url=https%3A%2F%2Fcdn.statically.io%2Fgh%2Fopenrunner%2Fclash-freenode%2Fmain%2Fclash.yaml&insert=false", false},
		{"https://freefq.neocities.org/free.txt", false},
		{"https://gitlab.com/api/v4/projects/36060645/repository/files/data%2Fv2ray%2FtvNUi5rjr.txt/raw?ref=main&private_token=glpat-iC4t7zq8nsV2xKYseBfU", false},
		{"https://raw.githubusercontent.com/mfuu/v2ray/master/v2ray", true},
		{"https://raw.githubusercontent.com/ermaozi/get_subscribe/main/subscribe/v2ray.txt", true},
		{"https://raw.githubusercontent.com/tbbatbb/Proxy/master/dist/v2ray.config.txt", true},
		{"https://raw.githubusercontent.com/vveg26/get_proxy/main/dist/v2ray.config.txt", true},
		{"https://freefq.neocities.org/free.txt", true},
		{"https://raw.githubusercontent.com/aiboboxx/v2rayfree/main/v2", true},
		{"https://raw.githubusercontent.com/ermaozi/get_subscribe/main/subscribe/v2ray.txt", true},
	}
}

func (that *SomeFree) getAndParse() {
	for _, _url := range that.Urls {
		fmt.Println("Handle: ", _url.Url)
		that.d = query.NewDownloader(_url.Url)
		if _url.EnableProxy {
			that.d.UseProxy()
		}
		resp := that.d.Get()
		if len(resp) > 0 {
			encryptStr := string(resp)
			htmlStr := utils.DecodeBase64(encryptStr)
			for _, s := range strings.Split(htmlStr, "\n") {
				s = strings.TrimSpace(s)
				if s != "" {
					that.Result = append(that.Result, s)
				}
			}
		}
	}
}

func (that *SomeFree) Run() []string {
	that.getAndParse()
	return that.Result
}
