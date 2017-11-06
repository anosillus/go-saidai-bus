package main

import (
	"fmt"
	"strings"
	"time"
)

// Time is structure of Hours "24" and Min "59".
type Time struct {
	Hour, Min int
}

// Day is current time data structure for Today.
type Day struct {
	Time
	Day, Month, Year     int
	WeekdayJP, WeekdayEn string
}

// var jpdays, endays = []string{0, 7}
var jpdays = [...]string{"日", "月", "火", "水", "木", "金", "土"}
var endays = [...]string{"Sun", "Mon", "Tue", "Wed", "Tho", "Fri", "Sat"}

var Today = Day{
	WeekdayJP: jpdays[time.Now().Weekday()],
	WeekdayEn: endays[time.Now().Weekday()],
	Day:       time.Now().Day(),
	// Month:     time.Now().Month(),
	Year: time.Now().Year(),
}

// Delay is Time plus DelayMin.
type BusDelay struct {
	Time
	DelayMin int
}

// Ride is structure of bus ride time.
type BusLeft struct {
	Time
}

type BusArrival struct {
	Time
}

// Company is structure of strings. "KKK", "国際興業バス" and "http://www.kokusaibus.com/blsys/".
type Company struct {
	CompanyAbbr, CompanyName string
	CssN
}

type Urls []Url

type Url struct {
	Station
	DistUrl string
}

// Css is Css position of the data.
type Css struct {
	PlanedLeft, RealLeft, NonStepBus, BusArrival string
}

type CssN []Css

type Station struct {
	Abbr, NameJp, NameEn, UrlKKK, UrlSB string
}

type Distination struct {
	BusArrival
	Station
}

//type Stations []Station

var MYKpenalty int = 5

var MYN Station = Station{
	Abbr:   "MYN",
	NameJp: "南与野駅西口",
	NameEn: "Minami-Yono station West gate",
	UrlKKK: "",
	UrlSB:  "",
}

var MYK Station = Station{
	Abbr:   "MYK",
	NameJp: "南与野駅北入口",
	NameEn: "Minami-Yono station North gate",
	UrlKKK: "http://www.kokusaibus.com/blsys/loca?EID=nt&DSMK=0015&ASMK=2482&VID=lsc",
	UrlSB:  "",
}

var KU Station = Station{
	Abbr:   "KU",
	NameJp: "北浦和駅",
	NameEn: "Kita-Urawa station",
	UrlKKK: "",
	UrlSB:  "",
}

var SHN Station = Station{
	Abbr:   "SHN",
	NameJp: "志木駅",
	NameEn: "Shinogi station",
	UrlKKK: "",
	UrlSB:  "",
}

type BusList []Bus

type Distinations []Distination

type Bus struct {
	Company
	BusLeft
	BusDelay
	Distinations
}

var kkk1 = Css{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(2) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(2) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(2) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(2) > td:nth-child(7)",
}

var kkk2 = Css{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(4) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(4) > td:nth-child(7)",
}
var kkk3 = Css{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(6) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(6) > td:nth-child(7)",
}

var TimeS string
var Separator string = ":"

func TimeSToI(time, separator string) {
	if strings.Index(time, ":") != 2 {
		fmt.Println("Scrape null")
	}
	else{
	fmt.Println(strings.SplitN(time, separator, 2))
	}
	// var l [3]string
	// l = strings.SplitN(time, separator, 2)
	// if l[0]
	// fmt.println(l)
}

func InitKKK() Company {
	var kkkCSS CssN
	kkkCSS = append(kkkCSS, kkk1)
	kkkCSS = append(kkkCSS, kkk2)
	kkkCSS = append(kkkCSS, kkk3)

	var kkk = Company{
		CompanyAbbr: "KKK",
		CompanyName: "国際興業",
		CssN:        kkkCSS,
	}
	return kkk
}

func (c *Company) Scrape(station *Station) {
	fmt.Println(c)
	fmt.Println(station)
}

func GetData() {
	var kkk Company
	kkk = InitKKK()
	// sb = InitSB()
	fmt.Println(kkk)
	// fmt.Println(kkk.Urls[0].DistUrl)
	TimeS = "12:30"
	// kkk.Scrape(MYN)
	TimeSToI(TimeS, Separator)
}

func main() {
	GetData()
}
