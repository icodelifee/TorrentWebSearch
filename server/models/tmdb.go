package models

type TMDBData struct {
	BackdropPath string   `json:"backdrop_path"`
	Genres       []string `json:"genres"`
	ID           int64    `json:"id"`
	Overview     string   `json:"overview"`
	Popularity   float64  `json:"popularity"`
	PosterPath   string   `json:"poster_path"`
	ReleaseDate  string   `json:"release_date"`
	Runtime      int64    `json:"runtime"`
	Tagline      string   `json:"tagline"`
	Title        string   `json:"title"`
	VoteAverage  float64  `json:"vote_average"`
	IMDBID       string   `json:"imdb_id"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TMDBMovie struct {
	Adult            bool        `json:"adult"`
	BackdropPath     string      `json:"backdrop_path"`
	Budget           int64       `json:"budget"`
	Genres           []Genre     `json:"genres"`
	Homepage         string      `json:"homepage"`
	ID               int64       `json:"id"`
	ImdbID           string      `json:"imdb_id"`
	OriginalLanguage string      `json:"original_language"`
	OriginalTitle    string      `json:"original_title"`
	Overview         string      `json:"overview"`
	Popularity       float64     `json:"popularity"`
	PosterPath       string      `json:"poster_path"`
	ReleaseDate      string      `json:"release_date"`
	Revenue          int64       `json:"revenue"`
	Runtime          int64       `json:"runtime"`
	Status           string      `json:"status"`
	Tagline          string      `json:"tagline"`
	Title            string      `json:"title"`
	Video            bool        `json:"video"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int64       `json:"vote_count"`
	ExternalIDS      ExternalIDS `json:"external_ids"`
}

type TMDBSeries struct {
	BackdropPath     string      `json:"backdrop_path"`
	EpisodeRunTime   []int64     `json:"episode_run_time"`
	FirstAirDate     string      `json:"first_air_date"`
	Genres           []Genre     `json:"genres"`
	Homepage         string      `json:"homepage"`
	ID               int64       `json:"id"`
	InProduction     bool        `json:"in_production"`
	Languages        []string    `json:"languages"`
	LastAirDate      string      `json:"last_air_date"`
	Name             string      `json:"name"`
	NumberOfEpisodes int64       `json:"number_of_episodes"`
	NumberOfSeasons  int64       `json:"number_of_seasons"`
	Overview         string      `json:"overview"`
	Popularity       float64     `json:"popularity"`
	PosterPath       string      `json:"poster_path"`
	Status           string      `json:"status"`
	Tagline          string      `json:"tagline"`
	Type             string      `json:"type"`
	VoteAverage      float64     `json:"vote_average"`
	VoteCount        int64       `json:"vote_count"`
	ExternalIDS      ExternalIDS `json:"external_ids"`
}
type ExternalIDS struct {
	ImdbID      string      `json:"imdb_id"`
	FreebaseMid string      `json:"freebase_mid"`
	FreebaseID  interface{} `json:"freebase_id"`
	TvdbID      int64       `json:"tvdb_id"`
	TvrageID    int64       `json:"tvrage_id"`
	FacebookID  string      `json:"facebook_id"`
	InstagramID string      `json:"instagram_id"`
	TwitterID   string      `json:"twitter_id"`
}
