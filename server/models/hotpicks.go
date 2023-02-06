package models

type HotPicks struct {
	UUID     string    `json:"uuid"`
	Torrents []Torrent `json:"torrents"`
	Imdb     struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Year     string `json:"year"`
		Hqposter string `json:"hqposter"`
		Genre    string `json:"genre"`
		Runtime  string `json:"runtime"`
		Rating   string `json:"rating"`
		Backdrop string `json:"backdrop"`
		Plot     string `json:"plot"`
	} `json:"imdb"`
}
