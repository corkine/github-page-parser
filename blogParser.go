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

var href = regexp.MustCompile(`(?:href|src)="(/.*?)"`)

var categoriesHash = regexp.MustCompile(`(/.*?)/#\S+`)

var cssUrl = regexp.MustCompile(`(?:url|URL)\(['|"]([a-zA-Z0-9.]+).*?['|"]\)`)

func downloadResource(host string, src string) []byte {
	fullUrl := host + src
	log.Printf("http request for %s start\n", fullUrl)
	get, err := http.Get(fullUrl)
	if err != nil {
		log.Printf("http request %s err in get\n", fullUrl)
		return nil
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		log.Printf("http request %s err in read\n", fullUrl)
		return nil
	}
	return all
}

//对于一些页面进行链接的特殊处理 (带 #hash 的删除 #hash，/ 直接跳过 etc)
func specUrlHandler(old string) (needHandle bool, newName string) {
	if old == "/" {
		return false, ""
	} else if categoriesHash.MatchString(old) {
		findRes := categoriesHash.FindAllStringSubmatch(old, -1)
		if len(findRes) >= 1 && len(findRes[0]) >= 2 {
			log.Printf("Replace %s to %v", old, findRes[0][1])
			return true, findRes[0][1]
		} else {
			log.Fatalf("Replace %s failed, regExp can't match!", old)
			return true, old
		}
	}
	return true, old
}

func findHTML(origin string, data []byte) map[string]bool {
	in := string(data)
	matched := href.FindAllStringSubmatch(in, -1)
	var result = make(map[string]bool)
	for _, item := range matched {
		if len(item) <= 1 {
			log.Fatalf("Can't handle match: %v", item)
		}
		find := item[1]
		//log.Printf("Find Link: %s", find)
		pass, renamedFind := specUrlHandler(find)
		if pass {
			result[renamedFind] = false
		} else {
			log.Printf("Skip %s because policy disallowed it.", find)
		}
	}
	return result
}

func findCSS(origin string, data []byte) map[string]bool {
	in := string(data)
	matched := cssUrl.FindAllStringSubmatch(in, -1)
	var result = make(map[string]bool)
	for _, item := range matched {
		if len(item) <= 1 {
			log.Fatalf("Can't handle match: %v", item)
		}
		find := item[1]
		//字体文件所在目录为发现的 CSS 目录
		find = path.Join(path.Dir(origin), find)
		log.Printf("Find URL in CSS: %s", find)
		pass, renamedFind := specUrlHandler(find)
		if pass {
			result[renamedFind] = false
		} else {
			log.Printf("Skip %s because policy disallowed it.", find)
		}
	}
	return result
}

func writeResource2File(root string, src string, in []byte) {
	var paths string
	if path.Ext(src) != "" { ///assert/prism.css -> /assert
		paths = path.Dir(src)
	} else { ///page14 -> /page14
		paths = src
	}
	fullPaths := path.Join(root, paths) // /Users/mazhangjing/gen/page14
	log.Printf("Handleing src %s at path %s\n", src, fullPaths)
	//for .css resource, fetch dir and try to create it,
	//if dir exist, do nothing
	err := os.MkdirAll(fullPaths, os.ModePerm)
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
		log.Fatalf("Write %s error: %s\n", src, err)
		return
	}
}

//一般情况下将获取到的资源当做 html 查找 href，但也不一定，比如对于资源文件 .css 可直接
func specFindAction(resource string) func(string, []byte) map[string]bool {
	if path.Ext(resource) == "" {
		log.Printf("Finding pointer %s as HTML file to fetch href/src start", resource)
		return findHTML
	}
	if strings.HasSuffix(strings.ToUpper(resource), ".CSS") {
		log.Printf("Finding pointer %s as CSS file to fetch url start", resource)
		return findCSS
	}
	return func(string, []byte) map[string]bool {
		log.Printf("Skip finding pointer in %s because it is a resource file", resource)
		return make(map[string]bool)
	}
}

func handleResource(host string, path string, root string) map[string]bool {
	pageData := downloadResource(host, path)
	if pageData == nil {
		log.Fatalf("http fetch for %s failure, exit now", path)
	}
	writeResource2File(root, path, pageData)
	pageRes := specFindAction(path)(path, pageData)
	return pageRes
}

func main() {
	var host = "https://blog.mazhangjing.com"
	//home, _ := os.UserHomeDir()
	var root = path.Join("/", "www", "wwwroot", "blog-cn.mazhangjing.com", "gen")
	err := os.RemoveAll(root)
	if err != nil {
		log.Fatalf("Can't remove old root %s %v", root, err)
		return
	}
	err = os.Mkdir(root, fs.ModePerm)
	if err != nil {
		log.Fatalf("Can't mkdir root %s %v", root, err)
		return
	}
	var resource = handleResource(host, "/", root)
	var collectRes = make(map[string]bool)
	allDone := false
	for {
		allDone = true
		for pageLink, ok := range resource {
			if !ok {
				allDone = false
				log.Printf("Start fetch and store %s\n", pageLink)
				pageFind := handleResource(host, pageLink, root)
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
