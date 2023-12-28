package article

type Poster struct {
	URL          string `json:"url"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	AverageColor string `json:"averageColor"`
}

type Hires struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Mimetype     string `json:"mimetype"`
	AverageColor string `json:"averageColor"`
}

type Cover struct {
	ID          string `json:"id"`
	Etag        string `json:"etag"`
	Hash        string `json:"hash"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
	Hires       Hires  `json:"hires"`
	Title       string `json:"title"`
	Copyright   string `json:"copyright"`
	Description string `json:"description"`
	Issuer      string `json:"issuer"`
	SourceURL   string `json:"sourceUrl"`
	DownloadURL string `json:"downloadUrl"`
	Poster      Poster `json:"poster"`
}
