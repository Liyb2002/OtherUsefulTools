package fetcher

import(
	"net/http"
	"fmt"
	"io/ioutil"
)

func Try() string{
	return "hll"
}

func Fetch(url string, myChan chan <-[]byte){

	req, _ := http.NewRequest("GET",url, nil)
	// 比如说设置个token
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
    // 再设置个json
  //  req.Header.Set("Content-Type","application/json")
    resp, err := (&http.Client{}).Do(req)
	if err!=nil{
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{
		fmt.Println("exit")
		myChan<- nil
		//fmt.Println(resp.StatusCode)
		return
	}

	allByte, _:=ioutil.ReadAll(resp.Body)
	myChan<- allByte
}