package models

import "encoding/json"

func UnmarshalSnowfl(data []byte) (Snowfl, error) {
	var r Snowfl
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Snowfl) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Snowfl struct {
	Name    string `json:"name"`
	Magnet  string `json:"magnet"`
	Size    string `json:"size"`
	Seeder  int64  `json:"seeder"`
	Leecher int64  `json:"leecher"`
	Type    string `json:"type"`
	Site    string `json:"site"`
	URL     string `json:"url"`
	Trusted bool   `json:"trusted"`
	Nsfw    bool   `json:"nsfw"`
	Age     string `json:"age"`
}
