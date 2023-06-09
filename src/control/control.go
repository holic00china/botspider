package control

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetResponse() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	requestUrl := "https://172.24.178.124:8002/"
	// 发送Get请求
	rsp, err := client.Get(requestUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//content := string(body)
	defer rsp.Body.Close()
	//fmt.Println(content)
	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		doc.Find("center>table").NextAll().Each(func(i int, table *goquery.Selection) {
			table.Find("tbody tr td a:nth-child(2)").Each(func(j int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				fmt.Println(href)
				// res, err := client.Get(requestUrl + href)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// defer res.Body.Close()
			})
		})
	}
}
