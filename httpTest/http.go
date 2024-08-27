package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	//url := "https://www.dongchedi.com/sales"
	url := "https://gorm.io/zh_CN/docs/"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest err = ", err)
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("client.Do(request) err = ", err)
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll(resp.Body) err = ", err)
		panic(err)
	}
	//fmt.Println("body = ", string(body))
	parse(string(body))
}

func parse(html string) {
	//替换空格
	html = strings.Replace(html, "\n", "", -1)
	fmt.Println("html = ", html)
	//正则匹配
	r_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	fmt.Println("r_sidebar = ", r_sidebar)

	sidebar := r_sidebar.FindString(html)
	fmt.Println("sidebar = ", sidebar)

	r_link := regexp.MustCompile(`href="(.*?)"`)
	fmt.Println("r_link = ", r_link)

	links := r_link.FindAllString(sidebar, -1)
	fmt.Println("links = ", links)

	//regexp.MustCompile()
	//r_sideName := regexp.MustCompile(`">(.*?)</a>`)
	r_sideName := regexp.MustCompile(`class="sidebar-link">(.*?)</a>`)
	fmt.Println("r_sideName = ", r_sideName)

	names := r_sideName.FindAllString(sidebar, -1)
	for _, name := range names {
		newName := strings.TrimPrefix(name, `class="sidebar-link">`)
		fmt.Println("newName = ", newName)
	}

	//strings.TrimSuffix()

	sideNames := r_sideName.FindAllString(sidebar, -1)
	fmt.Println("sideNames = ", sideNames)

	//url := "https://gorm.io/zh_CN/docs/"
	//
	//for _, link := range links {
	//	tmpLink := link[6 : len(link)-1]
	//	newUrl := url + tmpLink
	//	fmt.Println(newUrl)
	//}
}
