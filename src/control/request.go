package control

import (
	"BotSpider/src/structure"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/yaml.v3"
)

func GetEnv() structure.Env {
	data, err := ioutil.ReadFile("../etc/ua.yaml")
	var env structure.Env
	err = yaml.Unmarshal(data, &env)
	if err != nil {
		panic(err)
	}
	return env
}

func RepaetPost() {
	env := GetEnv()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "123456")
	sql := "'union select admin from users;#"
	params := url.Values{}
	params.Set("q", sql)

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	weakpwdUrl := env.Host + "/bodgeit/login.jsp"
	sqlinjectUrl := env.Host + "/bodgeit/search.jsp" + params.Encode()

	rand.Seed(time.Now().Unix())

	if rand.Intn(2) == 0 {
		fmt.Println("现在是弱口令攻击")
		resp, err := client.PostForm(weakpwdUrl, data)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	} else {
		fmt.Println("现在是sql注入攻击", sqlinjectUrl)
		req, err := http.NewRequest("GET", sqlinjectUrl, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.200 Safari/537.36 Qaxbrowser")
		req.Header.Set("Accept-Language", "en-US,en;q=0.8")
		req.Header.Set("Referer", sqlinjectUrl)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败：", err)
			return
		}
		defer resp.Body.Close()
	}

	time.Sleep(time.Duration(env.Timeout) * time.Second)
}

func SimpulateRequest() {
	env := GetEnv()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	data, err := ioutil.ReadFile("../etc/ua.yaml")
	if err != nil {
		panic(err)
	}
	var configList structure.Config
	err = yaml.Unmarshal(data, &configList)
	if err != nil {
		panic(err)
	}

	// 输出每个User-Agent
	for _, userAgent := range configList.UserAgent {
		client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
		fmt.Println(userAgent)
		// 发送 GET 请求
		req, err := http.NewRequest("GET", env.Host+"/ESAPI-Java-SwingSet-Interactive/main", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("User-Agent", userAgent)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		time.Sleep(time.Duration(env.Timeout) * time.Second)
	}
	//userAgent := "Sosospider+(+http://help.soso.com/webspider.htm)"//搜搜
	//userAgent := "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)"//雅虎
	//userAgent := "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0 Fiddler Spider" //Fiddler
	//userAgent := "InfoNaviRobot/0.9 (+http://www.infoseek.co.jp/)"//InfoNaviRobot
	//userAgent := "Internet Ninja/1.0"//Internet Ninja
	//userAgent := "Kenjin Spider/0.1" //kenjin
	//userAgent := "lexiBot/1.0"//lexibot
	//userAgent := "Mozilla/4.0 (compatible; MSIE 5.01; Windows 95; MSIECrawler)" //msie
	//userAgent := "NICErsPRO"
	//userAgent := "NPBot"
	//userAgent := "Navroad"

	// 读取响应
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // 输出响应
	// fmt.Println(string(body))
}
