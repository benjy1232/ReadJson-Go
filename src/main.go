package main

import "fmt"

func main() {
	YTJsons := ReadYTJson("music-history.json")
	for i := 0; i < len(YTJsons); i++ {
		fmt.Println(YTJsons[i].Header)
	}
}
