package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter ListenBrainz ID: ")
	var id string
	fmt.Scanln(&id)
	YTJsons := ReadYTJson("music-history.json")
	listen := ConvertToListenBrainz(YTJsons, "music-uploads-metadata.csv")
	SendToListenBrainz(listen, id)
}
