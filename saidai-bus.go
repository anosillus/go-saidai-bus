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

// Company is structure of strings "KKK", "国際興業バス" and CssN. Used by Scraping and Bus.
type Company struct {
	CompanyAbbr, CompanyName string
	CSSN
}

// CSS is CSS position's data structure. Used by CSSN and Company.
type CSS struct {
	PlanedLeft, RealLeft, NonStepBus, BusArrival string
}

// CSSN is CSS lists. Used by Company.
type CSSN []CSS

// Station is data sets relete with each Station.
type Station struct {
	Abbr, NameJp, NameEn, URLKkk, URLSb string
}

// Distination is structure of Bus distination and arrive time. Used by Bus.
type Distination struct {
	BusArrival
	Station
}

// Stations is list of Station.
type Stations []Station

var stations Stations

var kkk, sb Company

// Init make stations.
func Init() {
	stations = append(stations, MYN)
	stations = append(stations, MYK)
	stations = append(stations, KU)
	stations = append(stations, SHN)

	kkk = InitKKK()
	// sb = InitSB()
}

// MYKpenalty is 5 min additional time of MYK of arrival MinamiYono Station in comparison of MYN.
var MYKpenalty = 5

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

// BusList is list of Bus structure. It is the middle purpose.
type BusList []Bus

// Distinations is list of Distination.
type Distinations []Distination

// Bus is structure of bus related things data sets.
type Bus struct {
	Company
	BusPlaned
	Delay
	Distinations
}

// sb1 is 1st left bus information CSS data sets of Seibu bus on each URL.
var sb1 = CSS{
	PlanedLeft: "li#plot0 div.orvPane > div:nth-child(2)",
	RealLeft:   "li#plot0 div.orvPane > div:nth-child(3)",
	NonStepBus: "",
	BusArrival: "li#plot0 div.dnvPane > div:nth-child(2)",
}

// sb2 is 2nd left bus information CSS data sets of Seibu bus on each URL.
var sb2 = CSS{
	PlanedLeft: "li#plot1 div.orvPane > div:nth-child(2)",
	RealLeft:   "li#plot1 div.orvPane > div:nth-child(3)",
	NonStepBus: "",
	BusArrival: "li#plot1 div.dnvPane > div:nth-child(2)",
}

// sb3 is 3rd left bus information CSS data sets of Seibu bus on each URL.
var sb3 = CSS{
	PlanedLeft: "li#plot2 div.orvPane > div:nth-child(2)",
	RealLeft:   "li#plot2 div.orvPane > div:nth-child(3)",
	NonStepBus: "",
	BusArrival: "li#plot2 div.dnvPane > div:nth-child(2)",
}

// kkk1 is 1st left bus information CSS data sets of Kokusai bus on each URL.
var kkk1 = CSS{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(2) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(2) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(2) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(2) > td:nth-child(7)",
}

// kkk2 is 2nd left bus information CSS data sets of Kokusai bus on each URL.
var kkk2 = CSS{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(4) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(4) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(4) > td:nth-child(7)",
}

// kkk3 is 3rd left bus information CSS data sets of Kokusai bus on each URL.
var kkk3 = CSS{
	PlanedLeft: "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(1)",
	RealLeft:   "div#mainContents tr:nth-child(2) > td > table > tbody > tr:nth-child(6) > td:nth-child(2)",
	NonStepBus: "div#mainContents tr:nth-child(6) > td:nth-child(5)",
	BusArrival: "div#mainContents tr:nth-child(6) > td:nth-child(7)",
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

// InitKKK is Kokusai bus information initialize.
func InitKKK() Company {
	var kkkCSS CSSN
	kkkCSS = append(kkkCSS, kkk1)
	kkkCSS = append(kkkCSS, kkk2)
	kkkCSS = append(kkkCSS, kkk3)

	var kkk = Company{
		CompanyAbbr: "KKK",
		CompanyName: "国際興業",
		CSSN:        kkkCSS,
	}
	return kkk
}

// Scrape is Company method and return Bus Data.
func (c *Company) Scrape(station *Station) {
	fmt.Println(c)
	fmt.Println(station)
	var times TimeStr
	times = "12:30"
	times.Timetoi()
}

// GetData is Scrape and Format datas.
func GetData() {
	fmt.Println("Get Data Start")
	// fmt.Println(kkk)
	// fmt.Println(kkk.CSSN[0])
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
