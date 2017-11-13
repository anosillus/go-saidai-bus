package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

func main() {
	var url string
	url = "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2482&VID=lsc"
	res, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer res.Body.Close()
	var charset string
	charset = "shift_jis"
	// fmt.Println(charset)
	utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(doc.Text())
	aaa := doc.Find("div#mainContents tr:nth-child(10) > td:nth-child(2)").Text()
	fmt.Println(aaa)
	// fmt.Println(type(aaa))
	fmt.Println(reflect.TypeOf(aaa))
	fmt.Println("")
	fmt.Println(reflect.TypeOf(""))
	//人類に感謝
}
