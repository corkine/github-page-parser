package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

const root = "."

var href = regexp.MustCompile(`(?:href|src)="(/.*?)"`)

func parsePage(host string, src string) []byte {
	log.Printf("http request for %s%s start\n", host, src)
	get, err := http.Get(host + src)
	if err != nil {
		log.Printf("http request %s%s err in get\n", host, src)
		return nil
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		log.Printf("http request %s%s err in read\n", host, src)
		return nil
	}
	return all
}

func findPage(data []byte) map[string]bool {
	in := string(data)
	matched := href.FindAllStringSubmatch(in, -1)
	var result = make(map[string]bool)
	isResource := func(s string) bool {
		return strings.HasSuffix(strings.ToUpper(s), ".CSS") ||
			strings.HasSuffix(strings.ToUpper(s), ".JS") ||
			strings.HasSuffix(strings.ToUpper(s), ".XML") ||
			strings.HasSuffix(strings.ToUpper(s), ".ICO")
	}
	for _, item := range matched {
		if len(item) <= 1 {
			log.Fatalf("Can't handle match: %v", item)
		}
		find := item[1]
		if find != "/" {
			if isResource(find) {
				log.Printf("Find Resource: %s", find)
				result[find] = false
			} else {
				log.Printf("Find Link: %s", find)
				result[find] = false
			}
		}
	}
	return result
}

func writePage2File(root string, src string, in []byte) {
	paths := path.Dir(src)
	log.Printf("Handleing src %s at path %s\n", src, paths)
	//for .css resource, fetch dir and try to create it,
	//if dir exist, do nothing
	err := os.MkdirAll(paths, os.ModePerm)
	if err != nil {
		log.Printf("Mkdir for %s failed %s\n", src, err)
		return
	}
	var fullFileName string
	//for non resource html page, new file name is path/index.html
	if path.Ext(src) == "" {
		fullFileName = path.Join(root, src, "index.html")
		log.Printf("Write page to %s\n", fullFileName)
	} else {
		//for .css, .js etc resource, just write to src
		fullFileName = path.Join(root, src)
		log.Printf("Write resource to %s\n", src)
	}
	err = ioutil.WriteFile(fullFileName, in, fs.ModePerm)
	if err != nil {
		log.Printf("Write %s error: %s\n", src, err)
		return
	}
}

func handlePage(host string, path string) map[string]bool {
	pageData := parsePage(host, path)
	if pageData == nil {
		log.Fatalf("http fetch for %s failure, exit now", path)
	}
	writePage2File(root, path, pageData)
	pageRes := findPage(pageData)
	return pageRes
}

func main() {
	var host = "https://blog.mazhangjing.com"
	var resource = handlePage(host, "/")
	var collectRes = make(map[string]bool)
	allDone := false
	for {
		allDone = true
		for pageLink, ok := range resource {
			if !ok {
				allDone = false
				log.Printf("Start fetch and store %s\n", pageLink)
				pageFind := handlePage(host, pageLink)
				for pageLinkSub := range pageFind {
					collectRes[pageLinkSub] = false
				}
				resource[pageLink] = true
				log.Printf("Done fetch and store %s\n", pageLink)
				log.Printf("Cache new links pointed from %s\n", pageLink)
			}
		}
		log.Printf("Mergeing Cache links to all links\n")
		for pageLinkSub := range collectRes {
			find := resource[pageLinkSub]
			if !find {
				resource[pageLinkSub] = false
			}
		}
		collectRes = make(map[string]bool)
		if allDone {
			break
		}
	}
}
