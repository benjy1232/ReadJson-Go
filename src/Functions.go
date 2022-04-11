package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// ReadYTJson Reads an exported JSON from YouTube containing the history
// Converts said JSON into array of YTJson
func ReadYTJson(filename string) []YTJson {
	var YTJsons []YTJson
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	error := json.Unmarshal(byteValue, &YTJsons)
	if error != nil {
		fmt.Println(error)
	}
	return YTJsons
}

// ConvertToListenBrainz Converts a list of YTJson to a list of ListenBrainz in order to get ready to send
func ConvertToListenBrainz(ytJsons []YTJson, uploadedMusic string) []ListenBrainz {
	var FinalList []ListenBrainz
	uploadedSongs := readCSV(uploadedMusic)
	layout := "2006-01-02T15:04:05.000Z"

	for i := 0; i < len(ytJsons); i++ {
		if strings.Contains(ytJsons[i].Title, "Visited YouTube") ||
			strings.Contains(ytJsons[i].Title, "https://") ||
			ytJsons[i].Header == "YouTube" ||
			ytJsons[i].Title == "Watched a video that has been removed" {
			continue
		}
		var toAdd ListenBrainz

		toAdd.Track_metadata.Track_name = strings.TrimPrefix(ytJsons[i].Title, "Watched ")

		if strings.Contains(ytJsons[i].Subtitles[0].Name, "Music Library Uploads") {
			for _, songInfo := range uploadedSongs {
				if songInfo[0] == toAdd.Track_metadata.Track_name {
					toAdd.Track_metadata.Artist_name = songInfo[2]
					break
				}
				toAdd.Track_metadata.Artist_name = ytJsons[i].Subtitles[0].Name
			}
		} else {
			toAdd.Track_metadata.Artist_name = strings.TrimSuffix(ytJsons[i].Subtitles[0].Name, " - Topic")
		}

		t, err := time.Parse(layout, ytJsons[i].Time)

		if err != nil {
			backlayout := "2006-01-02T15:04:05Z"
			t, err = time.Parse(backlayout, ytJsons[i].Time)
		}

		toAdd.Listened_at = t.Unix()
		FinalList = append(FinalList, toAdd)
	}
	return FinalList
}

func readCSV(filename string) [][]string {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	csvSongs, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return csvSongs
}

func SendToListenBrainz(Songs []ListenBrainz, id string) {

	var Sending Sent
	clear := true
	Sending.Listen_type = "import"
	j := 0

	for _, currSong := range Songs {
		if clear {
			clear = false
			Sending.Payload = nil
		}
		Sending.Payload = append(Sending.Payload, currSong)

		if sent, _ := json.Marshal(Sending); len(sent) > 10200 {
			j++
			clear = true
			err := os.WriteFile("./JsonOut/"+strconv.Itoa(j)+".json", sent, 0644)
			if j%30 == 0 {
				time.Sleep(10 * time.Second)
			}
			if err != nil {
				fmt.Println(err)
			}
			url := "http://localhost:8100/1/submit-listens"
			var Token = "Token " + id
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(sent))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Add("Authorization", Token)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		}
	}
	sent, _ := json.Marshal(Sending)
	j++
	err := os.WriteFile("./JsonOut/"+strconv.Itoa(j)+".json", sent, 0644)
	if err != nil {
		fmt.Println(err)
	}
	url := "http://localhost:8100/1/submit-listens"
	var Token = "Token " + id
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sent))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
