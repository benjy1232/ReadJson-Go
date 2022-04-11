package main

type Track_metadata struct {
	Track_name  string `json:"track_name"`
	Artist_name string `json:"artist_name"`
}

type ListenBrainz struct {
	Track_metadata Track_metadata `json:"track_metadata"`
	Listened_at    int64          `json:"listened_at"`
}

type YTJson struct {
	Header           string
	Title            string
	TitleUrl         string
	Subtitles        []Subtitles
	Time             string
	Products         []string
	ActivityControls []string
}

type Subtitles struct {
	Name string
	Url  string
}

type Sent struct {
	Listen_type string         `json:"listen_type"`
	Payload     []ListenBrainz `json:"payload"`
}
