package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moqsien/free/pkgs/sites"
)

func PathIsExist(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

type Result struct {
	VmessList []string `json:"vmess"`
	SSRList   []string `json:"ssr"`
	VlessList []string `json:"vless"`
	SSList    []string `json:"ss"`
	Trojan    []string `json:"trojan"`
}

type Runner struct {
	result *Result
	sites  []sites.Site
}

func NewRunner() *Runner {
	return &Runner{
		result: &Result{
			VmessList: []string{},
			SSRList:   []string{},
			VlessList: []string{},
			SSList:    []string{},
			Trojan:    []string{},
		},
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
			v = strings.ReplaceAll(v, " ", "")
			if strings.HasPrefix(v, "vmess") {
				that.result.VmessList = append(that.result.VmessList, v)
			} else if strings.HasPrefix(v, "vless") {
				that.result.VlessList = append(that.result.VlessList, v)
			} else if strings.HasPrefix(v, "ssr") {
				that.result.SSRList = append(that.result.SSRList, v)
			} else if strings.HasPrefix(v, "ss") {
				that.result.SSList = append(that.result.SSList, v)
			} else if strings.HasPrefix(v, "trojan") {
				that.result.Trojan = append(that.result.Trojan, v)
			} else {
				fmt.Println("Do not support: ", v)
			}
		}
	}
}

func (that *Runner) save() {
	fpath := os.Getenv("FREE_VPN_DIR")
	if ok, _ := PathIsExist(fpath); !ok {
		os.MkdirAll(fpath, 0666)
	}
	if result, err := json.MarshalIndent(that.result, "", "   "); err == nil {
		fpath = filepath.Join(fpath, "free_vpn.json")
		os.WriteFile(fpath, result, os.ModePerm)
	}
}

func (that *Runner) Run() {
	that.getVpns()
	that.save()
}
