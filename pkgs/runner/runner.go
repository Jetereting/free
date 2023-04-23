package runner

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/container/gmap"
	"github.com/moqsien/free/pkgs/sites"
	"github.com/moqsien/free/pkgs/utils"
)

type VList struct {
	Total int      `json:"total"`
	List  []string `json:"list"`
}

type Result struct {
	VmessList  *VList `json:"vmess"`
	SSRList    *VList `json:"ssr"`
	VlessList  *VList `json:"vless"`
	SSList     *VList `json:"ss"`
	Trojan     *VList `json:"trojan"`
	Other      *VList `json:"other"`
	UpdateTime string `json:"update_time"`
}

type Runner struct {
	result *Result
	vmess  *gmap.StrAnyMap
	ssr    *gmap.StrAnyMap
	ss     *gmap.StrAnyMap
	trojan *gmap.StrAnyMap
	vless  *gmap.StrAnyMap
	other  *gmap.StrAnyMap
	sites  []sites.Site
}

func NewRunner() *Runner {
	return &Runner{
		result: &Result{
			VmessList: &VList{},
			VlessList: &VList{},
			SSRList:   &VList{},
			SSList:    &VList{},
			Trojan:    &VList{},
			Other:     &VList{},
		},
		vmess:  gmap.NewStrAnyMap(true),
		ssr:    gmap.NewStrAnyMap(true),
		ss:     gmap.NewStrAnyMap(true),
		trojan: gmap.NewStrAnyMap(true),
		vless:  gmap.NewStrAnyMap(true),
		other:  gmap.NewStrAnyMap(true),
	}
}

func (that *Runner) RegisterSite(site sites.Site) {
	if site != nil {
		that.sites = append(that.sites, site)
	}
}

func (that *Runner) getVpns() {
	for _, site := range that.sites {
		vpnList := site.Run()
		for _, v := range vpnList {
			v = strings.TrimSpace(v)
			if strings.HasPrefix(v, "vmess") {
				that.vmess.Set(v, struct{}{})
			} else if strings.HasPrefix(v, "vless") {
				that.vless.Set(v, struct{}{})
			} else if strings.HasPrefix(v, "ssr") {
				that.ssr.Set(v, struct{}{})
			} else if strings.HasPrefix(v, "ss") {
				that.ss.Set(v, struct{}{})
			} else if strings.HasPrefix(v, "trojan") {
				that.trojan.Set(v, struct{}{})
			} else {
				that.other.Set(v, struct{}{})
			}
		}
	}
	that.result.VmessList.List = that.vmess.Keys()
	that.result.VmessList.Total = that.vmess.Size()
	that.result.VlessList.List = that.vless.Keys()
	that.result.VlessList.Total = that.vless.Size()
	that.result.SSRList.List = that.ssr.Keys()
	that.result.SSRList.Total = that.ssr.Size()
	that.result.SSList.List = that.ss.Keys()
	that.result.SSList.Total = that.ss.Size()
	that.result.Trojan.List = that.trojan.Keys()
	that.result.Trojan.Total = that.trojan.Size()
	that.result.Other.List = that.other.Keys()
	that.result.Other.Total = that.other.Size()
	that.result.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
}

func (that *Runner) save(fpath string) {
	if ok, _ := utils.PathIsExist(fpath); !ok {
		os.MkdirAll(fpath, 0666)
	}
	if result, err := json.MarshalIndent(that.result, "", "   "); err == nil {
		fpath = filepath.Join(fpath, "free.txt")
		res := base64.StdEncoding.EncodeToString(result)
		os.WriteFile(fpath, []byte(res), os.ModePerm)
	}
}

func (that *Runner) Run(storage_dir string) {
	that.getVpns()
	that.save(storage_dir)
}
