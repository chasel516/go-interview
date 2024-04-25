package main

import (
	"fmt"
)

func main() {
	// url := "http://google.co.jp"

	// resp, _ := http.Get(url)
	// defer resp.Body.Close()

	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray))

	// url := "http://google.co.jp"

	// resp, _ := http.Get(url)
	// defer resp.Body.Close()

	// byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(byteArray)) // htmlをstringで取得

	// req,_:=http.NewRequest("GEt",)

	merchantName4Rune := []rune("有限公司")[:4] //slice bounds out of range [:4] with capacity 3
	merchantName := string(merchantName4Rune)
	fmt.Println(merchantName)

}
