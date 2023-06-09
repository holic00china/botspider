package main

import (
	"BotSpider/src/control"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

// var wg sync.WaitGroup
type Switchbtn struct {
	Button int `yaml:"SWITCH"`
}

func main() {
	log.Println(time.Now())
	data, err := ioutil.ReadFile("../etc/ua.yaml")
	var switchbtn Switchbtn
	err = yaml.Unmarshal(data, &switchbtn)
	if err != nil {
		panic(err)
	}
	if switchbtn.Button == 0 {
		for {
			control.RepaetPost()
		}
	} else if switchbtn.Button == 1 {
		for {
			control.SimpulateRequest()
		}
	}
}
