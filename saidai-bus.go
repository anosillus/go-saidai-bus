package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

// Time is structure of Hours "24" and Min "59".
type Time struct {
	Hour, Min int
}

// Day is current time data structure for Today.
type Day struct {
	Time
	Day, Month, Year     int
	WeekdayJp, WeekdayEn string
}

// var jpdays, endays = []string{0, 7}
var jpdays = [...]string{"日", "月", "火", "水", "木", "金", "土"}
var endays = [...]string{"Sun", "Mon", "Tue", "Wed", "Tho", "Fri", "Sat"}

// Today is today's data structure of Day.
var Today = Day{
	WeekdayJp: jpdays[time.Now().Weekday()],
	WeekdayEn: endays[time.Now().Weekday()],
	Day:       time.Now().Day(),
	Month:     int(time.Now().Month()),
	Year:      time.Now().Year(),
}

// Delay is bus delay time structure Used by Bus. Time and DelayMin.
type Delay struct {
	Time
	DelayMin int
}

// BusPlaned is structure of Time Used by Bus. Planed leave time.
type BusPlaned struct {
	Time
}

// BusArrival is structure of Time. Used by Bus. Time of arrive station.
type BusArrival struct {
	Time
}

// CSS is CSS position's data structure. Used by Init(), CSSN and Company.
type CSS struct {
	PlanedLeft, RealLeft, NonStepBus, BusArrival string
}

// CSSN is CSS lists. Used by Company and Init().
type CSSN []CSS

// Company is structure of strings "KKK", "国際興業バス" and CssN. Used by Scraping and Bus.
type Company struct {
	CompanyAbbr, CompanyName string
	CSSN
}

// Station is data sets relete with each Station.
type Station struct {
	Abbr, NameJp, NameEn, URLKkk, URLSb string
}

// Distination is structure of Bus distination and arrive time. Used by Bus.
type Distination struct {
	BusArrival
	Station
}

// MYN is Minami-Yono Nishi-gate bus station data structure.
var MYN = Station{
	Abbr:   "MYN",
	NameJp: "南与野駅西口",
	NameEn: "Minami-Yono station West gate",
	URLKkk: "",
	URLSb:  "http://transfer.navitime.biz/seibubus-dia/pc/location/BusLocationResult?startId=00111643&goalId=00111644",
}

// MYK is Minami-Yono Kita-gato bus station data structure.
var MYK = Station{
	Abbr:   "MYK",
	NameJp: "南与野駅北入口",
	NameEn: "Minami-Yono station North gate",
	URLKkk: "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2482&VID=lsc",
	URLSb:  "http://transfer.navitime.biz/seibubus-dia/pc/location/BusLocationResult?startId=00111643&goalId=00111639",
}

// KU is Kita-Urawa bus station data structure.
var KU = Station{
	Abbr:   "KU",
	NameJp: "北浦和駅",
	NameEn: "Kita-Urawa station",
	URLKkk: "",
	URLSb:  "",
}

// SHN is Shinogi station data structure.
var SHN = Station{
	Abbr:   "SHN",
	NameJp: "志木駅",
	NameEn: "Shinogi station",
	URLKkk: "",
	URLSb:  "",
}

// MYKpenalty is 5 min additional time of MYK of arrival MinamiYono Station in comparison of MYN.
var MYKpenalty = 5

// ScrapeDataNumber is the number how many data I scrape for same company same distination.
var ScrapeDataNumber = 3

var kkk, sb Company

// Stations is list of Station.
type Stations []Station

var stations Stations

// Init make stations.
func Init() {
	stations = append(stations, MYN)
	stations = append(stations, MYK)
	stations = append(stations, KU)
	stations = append(stations, SHN)

	kkk = InitKKK()
	sb = InitSB()
}

// BusList is list of Bus structure. It is the middle purpose.
type BusList []Bus

// Distinations is list of Distination.
type Distinations []Distination

// Bus is structure of bus related things data sets.
type Bus struct {
	Name string
	Company
	BusPlaned
	Delay
	Distinations
}

// ScrapeList is data structure of ScrapeStringi
type ScrapeList []ScrapeString

// ScrapeStringis data structure of scraping like "12:30", "12:35", "No" , "12:50".
type ScrapeString struct {
	Name       string
	PlanedLeft TimeStr
	RealLeft   TimeStr
	NonStepBus string
	BusArrival TimeStr
}

// TimeStr must be string "12:30", distinguish form [12 30] and "Next 12:30".
type TimeStr string

// Separator is ":". Both company are using "12:30" but [12 30] is better.
var Separator = ":"

// TimeSeparate is a method. TimeStr "12:30" make into "12", "30".
func (stime TimeStr) TimeSeparate() (string, string) {
	var h, m string
	l := strings.SplitN(string(stime), Separator, 2)
	h, m = l[0], l[1]
	return h, m
}

// Timetoi is a method of TimeStr returning int hour and int minuts.
func (stime TimeStr) Timetoi() (int, int) {
	h, m := stime.TimeSeparate()
	fmt.Println(h, m, "aaa")
	var a = 1
	var b = 1
	return a, b
}

// InitSB is CSS data and Company data making.
func InitSB() Company {
	var cssn CSSN
	for i := 0; i < ScrapeDataNumber; i++ {
		Num := i - 1
		NumS := strconv.Itoa(Num)
		var sbcss = CSS{
			PlanedLeft: "li#plot" + NumS + " div.orvPane > div:nth-child(2)",
			RealLeft:   "li#plot" + NumS + " div.orvPane > div:nth-child(3)",
			NonStepBus: "",
			BusArrival: "li#plot" + NumS + " div.dnvPane > div:nth-child(2)",
		}
		cssn = append(cssn, sbcss)
	}
	var sb = Company{
		CompanyAbbr: "SB",
		CompanyName: "西武バス",
		CSSN:        cssn,
	}
	return sb
}

// InitKKK is Kokusai bus information initialize.
func InitKKK() Company {
	var cssn CSSN
	for i := 0; i < ScrapeDataNumber; i++ {
		Num := i * 2
		NumS := strconv.Itoa(Num)
		var kkkcss = CSS{
			PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(" + NumS + ")> td:nth-child(1)",
			RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(" + NumS + ") > td:nth-child(2)",
			NonStepBus: "div#mainContents tr:nth-child(" + NumS + ") > td:nth-child(5)",
			BusArrival: "div#mainContents tr:nth-child(" + NumS + ") > td:nth-child(7)",
		}
		cssn = append(cssn, kkkcss)
	}
	var kkk = Company{
		CompanyAbbr: "KKK",
		CompanyName: "国際興業",
		CSSN:        cssn,
	}
	return kkk
}

// func (css *CSS) Scrape(station *Station) BusList {
// 	for i := 0; i < 3; i++
// 	switch c.CompanyAbbr{
// 	// case "KKK": url :=
// 	}
// 	url
// 	url = "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2482&VID=lsc"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		// handle error
// 	}
// 	defer res.Body.Close()
// 	var charset string
// 	charset = "shift_jis"
// 	// fmt.Println(charset)
// 	utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	doc, err := goquery.NewDocumentFromReader(utfBody)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// fmt.Println(doc.Text())
// 	fmt.Print(doc.Find("div#mainContents tr:nth-child(3) > td:nth-child(2)").Text())
//
// 	sh
//
// fmt.Println(c.CSSN[2].PlanedLeft, station.URLKkk)
// return kkkBus
// }

// (TimeStr, Bus)

func (c *Company) Access(url *string) ScrapeList {
	res, err := http.Get(*url)
	fmt.Println(res)
	if err != nil {
		fmt.Println("No NetConnection")
	}
	defer res.Body.Close()
	var charset string
	charset = "shift_jis"
	utfBody, err := iconv.NewReader(res.Body, charset, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Query Erro")
	}
	var scrapelist ScrapeList

	for i := 0; i < ScrapeDataNumber; i++ {
		var scrapestr = ScrapeString{
			PlanedLeft: TimeStr(doc.Find("c.CSSN[i].PlanedLeft").Text()),
			RealLeft:   TimeStr(doc.Find("c.CSSN[i].RealLeft").Text()),
			NonStepBus: doc.Find("c.CSNN[i].NonStepBus").Text(),
			BusArrival: TimeStr(doc.Find("c.CSNN[i].BusArrival").Text()),
		}
		scrapelist = append(scrapelist, scrapestr)
	}
	return scrapelist
}

// Scrape is Company method and return Bus Data.
func (c *Company) Scrape(station *Station) {
	switch c.CompanyAbbr {
	case "KKK":
		fmt.Println("kkk start")
		scrapelist := c.Access(&station.URLKkk)
		fmt.Println(&scrapelist)
	case "SB":
		scrapelist := c.Access(&station.URLSb)
		fmt.Println(scrapelist)
	default:
		fmt.Printf("Company Name Error")
	}
	// fmt.Println(c)
	// fmt.Println(station)
	// var times TimeStr
	// times = "12:30"
	// fmt.Println(scrapelist)
}

// GetData is Scrape and Format datas.
func GetData() {
	fmt.Println("Get Data Start")
	// fmt.Println(kkk)
	// fmt.Println(kkk.CSSN[0])
	// for
	kkk.Scrape(&MYN)
	fmt.Println("Get Data Finish")
}

func main() {
	fmt.Println("Main Start")
	Init()
	GetData()
	// OrdarBusList()
	// ShowData()
	fmt.Println("Main End")
}
