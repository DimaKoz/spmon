package model

const ContentTypeArticle = "articles"

type Handshake struct {
	UpdatedAt        int       `json:"updatedAt"`
	UpdatedAtIso8601 string    `json:"updatedAtIso8601"`
	Version          string    `json:"version"`
	Issuer           string    `json:"issuer"`
	Sources          Sources   `json:"sources"`
	Sections         []Section `json:"sections"`
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

type Block struct {
	ID                 string           `json:"id"`
	ContentType        string           `json:"contentType"`
	SortBy             string           `json:"sortBy"`
	SortDir            string           `json:"sortDir"`
	DataURLQueryArgs   DataURLQueryArgs `json:"dataUrlQueryArgs"`
	Hash               string           `json:"hash"`
	DataURLQueryString string           `json:"dataUrlQueryString"`
	DataURL            string           `json:"dataUrl"`
}

type Feed struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Blocks []Block `json:"blocks"`
}

type Section struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Feeds []Feed `json:"feeds"`
}

func (section *Section) GetBlocks() []Block {
	if section == nil {
		return []Block{}
	}
	result := make([]Block, 0)
	for _, feed := range section.Feeds {
		if len(feed.Blocks) > 0 {
			result = append(result, feed.Blocks...)
		}
	}

	return result
}

func (handshake *Handshake) GetArticleBlocks() []Block {
	result := make([]Block, 0)
	if handshake == nil {
		return result
	}
	for _, section := range handshake.Sections {
		blocks := section.GetBlocks()

		for _, block := range blocks {
			if block.ContentType == ContentTypeArticle {
				result = append(result, block)
			}
		}
	}

	return result
}
