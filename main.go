package main

import (
	"fmt"
	//go get github.com/huichen/sego
	"github.com/huichen/sego"
	"io"
	"net/http"
	"log"
	"os"
	"encoding/json"
)

var segmenter sego.Segmenter

type response struct {
	Status   int
	Msg      string
	Response string
}

// hello world, the web server
func handler(w http.ResponseWriter, req *http.Request) {
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	content, found := req.Form["content"]
	_, deep_search_found := req.Form["deep_search"]

	//fmt.Println(time.Now(), " found:", found)

	res := response{1, "", ""}

	if found {
		if len(content[0]) > 10000 {
			res.Status = 0
			res.Msg = "超过最大处理内容长度"
		} else {
			var text []byte = []byte(content[0])
			segments := segmenter.Segment(text)
			if deep_search_found {
				res.Response = sego.SegmentsToString(segments, true)
			} else {
				res.Response = sego.SegmentsToString(segments, false)
			}
		}
	} else {
		res.Status = 0
		res.Msg = "请传入正确的参数"
	}

	rjson, errs := json.Marshal(res)

	if errs != nil {
		fmt.Println(errs.Error())
	}

	io.WriteString(w, string(rjson))
	rjson = nil
}

func main() {
	var port string;
	if len(os.Args) >= 2 {
		port = os.Args[1]
	} else {
		println("port not found")
		return
	}
	segmenter.LoadDictionary("dictionary.txt")
	http.HandleFunc("/segment", handler)
	err := http.ListenAndServe("127.0.0.1:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
