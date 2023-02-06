package models

type Torrent struct {
	Title  string `json:"title"`
	Seeds  string `json:"seeds"`
	Leechs string `json:"leechs"`
	Size   string `json:"size"`
	Date   string `json:"added"`
	Link   string `json:"link"`
	Magnet string `json:"magnet"`
}
