package main

import (
	"os"
	"time"

	"github.com/moqsien/free/pkgs/runner"
	"github.com/moqsien/free/pkgs/sites"
)

/*
https://clashnode.com/wp-content/uploads/%s.txt
https://nodefree.org/dy/%s.txt
https://gitlab.com/mianfeifq/share/-/raw/master/data2023036.txt
https://raw.fastgit.org/freefq/free/master/v2
https://raw.githubusercontent.com/mfuu/v2ray/master/v2ray
https://sub.nicevpn.top/long
https://raw.githubusercontent.com/ermaozi/get_subscribe/main/subscribe/v2ray.txt
https://raw.githubusercontent.com/tbbatbb/Proxy/master/dist/v2ray.config.txt
https://raw.githubusercontent.com/vveg26/get_proxy/main/dist/v2ray.config.txt
https://freefq.neocities.org/free.txt
https://ghproxy.com/https://raw.githubusercontent.com/kxswa/k/k/base64
https://raw.githubusercontent.com/aiboboxx/v2rayfree/main/v2
https://gitlab.com/api/v4/projects/36060645/repository/files/data%2Fv2ray%2FtvNUi5rjr.txt/raw?ref=main&private_token=glpat-iC4t7zq8nsV2xKYseBfU


https://freefq.com/v2ray/
https://cdn.statically.io/gh/openrunner/clash-freenode/main/clash.yaml
https://cdn.jsdelivr.net/gh/zyzmzyz/free-nodes@master/Clash.yml
https://raw.githubusercontent.com/aiboboxx/v2rayfree/main/README.md
https://jichangtuijian.com/%E5%85%8D%E8%B4%B9ssr%E5%92%8Cv2ray%E6%9C%BA%E5%9C%BA.html
https://jichangtuijian.com/%E6%9C%BA%E5%9C%BA%E8%AE%A2%E9%98%85%E9%93%BE%E6%8E%A5%E8%BD%AC%E6%8D%A2%E6%95%99%E7%A8%8B.html
4 https://www.wenpblog.com/list/5/1.html
3 https://www.cfmem.com/2023/04/20230418-202350-v2rayclash-vpn.html
*/

func init() {
	var cstZone = time.FixedZone("CST", 8*3600)
	time.Local = cstZone
	// os.Setenv("FREE_PROXY", "socks5://localhost:1089")
	os.Setenv("FREE_PROXY", "socks5://localhost:2019")
}

func main() {
	rn := runner.NewRunner()
	rn.RegisterSite(sites.NewMianfieFQ())
	rn.RegisterSite(sites.NewCFMem())
	rn.RegisterSite(sites.NewWPBlog())
	rn.RegisterSite(sites.NewSomeFree())
	rn.RegisterSite(sites.NewFreeNode())
	storage_dir := os.Getenv("FREE_VPN_DIR")
	rn.Run(storage_dir)
	// example.Test()
}
