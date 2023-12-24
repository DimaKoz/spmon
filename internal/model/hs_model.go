package model

import "time"

type Handshake struct {
	UpdatedAt        int        `json:"updatedAt"`
	UpdatedAtIso8601 time.Time  `json:"updatedAtIso8601"`
	Version          string     `json:"version"`
	Issuer           string     `json:"issuer"`
	Sources          Sources    `json:"sources"`
	Sections         []Sections `json:"sections"`
}

type Sources struct {
	Articles string `json:"articles"`
	Lists    string `json:"lists"`
	Images   string `json:"images"`
}

type DataURLQueryArgs struct {
	Limit        string `json:"limit"`
	Lists        string `json:"lists"`
	UpdatedSince string `json:"updatedSince"`
	Issuer       string `json:"issuer"`
}

type Blocks struct {
	ID                 string           `json:"id"`
	ContentType        string           `json:"contentType"`
	SortBy             string           `json:"sortBy"`
	SortDir            string           `json:"sortDir"`
	DataURLQueryArgs   DataURLQueryArgs `json:"dataUrlQueryArgs"`
	Hash               string           `json:"hash"`
	DataURLQueryString string           `json:"dataUrlQueryString"`
	DataURL            string           `json:"dataUrl"`
}

type Feeds struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Blocks []Blocks `json:"blocks"`
}

type Sections struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Feeds []Feeds `json:"feeds"`
}
