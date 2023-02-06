package models

// AnimeData struct
type AnimeData struct {
	Title       string        `json:"title"`
	ShowID      string        `json:"showID"`
	Description string        `json:"desc"`
	Poster      string        `json:"image"`
	Episodes    []EpisodeData `json:"episodes"`
}

// EpisodeData of anime struct
type EpisodeData struct {
	EpName   string `json:"epname"`
	Three60p string `json:"360p"`
	Four80p  string `json:"480p"`
	Seven20p string `json:"720p"`
	Ten80p   string `json:"1080p"`
}
