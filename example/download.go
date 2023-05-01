package example

import "github.com/moqsien/free/pkgs/query"

func Test() {
	d := query.NewDownloader("https://gitee.com/moqsien/gvc/releases/download/v1/typst-darwin.zip")
	d.GetFileWithBar("/Users/moqsien/data/projects/go/src/free/typst.zip", true)
}
