package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

type Time struct {
	MYK_t, MYN_t, KU_t int
}

func KKKgetData(url string) string {
	res, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer res.Body.Close()
	charset := "shift_jis"
	fmt.Println(charset)
	utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Text())
	fmt.Println(doc.Find("div#mainContents tr:nth-child(3) > td:nth-child(2)").Text())
	time := doc.Find("div#mainContents tr:nth-child(3) > td:nth-child(2)").Text()
	return time
}

func KKK() {
	urlmap := make(map[string]string)
	var urls = urlmap
	urls["MYK"] = "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2482&VID=lsc"
	urls["MYN"] = "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=3333&VID=lsc"
	urls["KU"] = "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2541&VID=lsc"
	// distination = map[int]string{0: "MYK", 1: "MYN", 2: "KU"}
	MYK := KKKgetData(urls["MYK"])
	fmt.Println(MYK)
}

// func GetDay() {
// 	t := time.Now()
// 	weekday := t.Weekday()
// 	day := t.Day()
// 	month := t.Month()
// 	year := t.Year()
// }

// func GetTime() {
// 	t := time.Now()
// 	hour := t.Hour()
// 	min := t.Minute()
// }

func main() {
	comapny := [2]string{"KKK", "SB"}
	fmt.Print(comapny[1])
	// GetDay()
	// t := Time{{1, 2, 3}, {2, 2, 2}, {2, 2, 2}}
	// fmt.Print(t)
}
