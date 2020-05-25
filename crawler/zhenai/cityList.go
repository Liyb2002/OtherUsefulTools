package zhenai

import(
	"fmt"
	"regexp"
	//"github.com/fetcher"
)

func Try() string{
	return "zhenai"
}

func ParseCityList(contents []byte) [][][]byte{
	re:=regexp.MustCompile(`(<a href=")(http://www.zhenai.com/zhenghun/[^"]+)(" data-v-2cb5b6a2>)([^<]+)(</a>)`)

	CityMatch := re.FindAllSubmatch(contents, -1)
	return CityMatch

}

//Find people in a city
func ParseSingleCity(contents []byte, cityName string) [][][]byte{
	fmt.Println("newPage in", cityName, "!")
	re:=regexp.MustCompile(`(<th><a href=")(http://album.zhenai.com/u/[^"]+)`)
	PeoplePage := re.FindAllSubmatch(contents, -1)
	return PeoplePage
}