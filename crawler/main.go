package main

import (
	"fmt"
//	"net/http"
//	"io/ioutil"
//	"regexp"
	"strconv"
	"github.com/fetcher"
	"github.com/zhenai"
//	"reflect"
)

func main(){
	MyChan := make(chan[]byte)
	go fetcher.Fetch("https://www.zhenai.com/zhenghun", MyChan)
	allByte:=<-MyChan

	cityList:=zhenai.ParseCityList(allByte)
	fmt.Println("hi")
	for m := range cityList {
	fmt.Println(string(cityList[m][0]))}

	//Print out cityList
	//cityList[2]: url; cityList[4]: city name; 
	
	Totalcount:=0
	for m := range cityList {
	//Deal with a single city
	TemptCityUrl:=cityList[m][2]
	cityName:=string(cityList[m][4])
	fmt.Println("City is:", cityName)
	peopleInCity:= zhenai.ParseSingleCity(TemptCityUrl, cityName)

	//Add pages
	TTemptCityUrl:=string(TemptCityUrl)+"/"
	page:=2
	for {
		NTemptCity:=TTemptCityUrl+strconv.Itoa(page)
		go fetcher.Fetch(string(NTemptCity), MyChan)
		NTemptCityPeople:=<-MyChan
		if NTemptCityPeople==nil{
			break
		}
		//Save all people in this city into S
		s:=zhenai.ParseSingleCity(NTemptCityPeople, cityName)
		fmt.Println(string(s[0][0]))
		for m:=range s{
			peopleInCity=append(peopleInCity, s[m])
		}
		page=page+1
	}
	count:=0
	for m:=range peopleInCity{
		count=count+1
		Totalcount=Totalcount+1
	
		fmt.Println(string(peopleInCity[m][2]))
	}
	fmt.Println("People in",cityName, count)
}
fmt.Println("Total People in Zhenai", Totalcount)
}
