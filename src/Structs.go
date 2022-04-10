package main

type Track_metadata struct {
	Track_name  string
	Artist_name string
}

type ListenBrainz struct {
	Track_metadata Track_metadata
	Listened_at    int64
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
