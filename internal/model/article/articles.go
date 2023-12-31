package article

import "time"

type ServerAnswer struct {
	Context            CtxArticle         `json:"context"`
	InsertionMode      string             `json:"insertionMode"`
	UpdateURLQueryArgs UpdateURLQueryArgs `json:"updateUrlQueryArgs"`
	UpdateURL          string             `json:"updateUrl"`
	Articles           []Article          `json:"articles"`
}

type CtxArticle struct {
	Limit        string `json:"limit"`
	Lists        string `json:"lists"`
	UpdatedSince string `json:"updatedSince"`
	Issuer       string `json:"issuer"`
}

type UpdateURLQueryArgs struct {
	Limit  string `json:"limit"`
	Lists  string `json:"lists"`
	Issuer string `json:"issuer"`
}

type Author struct {
	URL              string `json:"url"`
	LastName         string `json:"lastName"`
	Biography        string `json:"biography"`
	FirstName        string `json:"firstName"`
	DisplayName      string `json:"displayName"`
	Name             string `json:"name"`
	NameComment      string `json:"name-comment"` //nolint:tagliatelle
	Avatar           Media  `json:"avatar"`
	AvatarURL        string `json:"avatarUrl"`
	AvatarURLComment string `json:"avatarUrl-comment"` //nolint:tagliatelle
}

type Headline struct {
	Title       string   `json:"title"`
	Lead        string   `json:"lead"`
	Authors     []Author `json:"authors"`
	PublishedAt int      `json:"publishedAt"`
	ModifiedAt  int      `json:"modifiedAt"`
	Cover       Media    `json:"cover"`
}

type PublicationDateTime struct {
	Unix int `json:"unix"`
}

type ModificationDateTime struct {
	Unix int `json:"unix"`
}

type Reference struct {
	ID                 string    `json:"id"`
	URL                string    `json:"url"`
	Title              string    `json:"title"`
	Issuer             string    `json:"issuer"`
	PublishedAt        int       `json:"publishedAt"`
	PublishedAtIso8601 time.Time `json:"publishedAtIso8601"`
	Cover              Media     `json:"cover"`
}

type Body struct {
	Type                 string               `json:"type"`
	Headline             Headline             `json:"headline,omitempty"`
	PublicationDateTime  PublicationDateTime  `json:"publicationDateTime,omitempty"`
	ModificationDateTime ModificationDateTime `json:"modificationDateTime,omitempty"`
	Authors              []Author             `json:"authors,omitempty"`
	Lead                 string               `json:"lead,omitempty"`
	HTML                 string               `json:"html,omitempty"`
	Medias               []Media              `json:"medias,omitempty"`
	Quote                string               `json:"quote,omitempty"`
	Placement            string               `json:"placement,omitempty"`
	Article              MentionedArticle     `json:"article,omitempty"`
	Subtype              string               `json:"subtype,omitempty"`
	Lists                []List               `json:"lists,omitempty"`
	References           []Reference          `json:"references,omitempty"`
	Photobook            []Media              `json:"photobook,omitempty"`
}

type List struct {
	ID                string `json:"id"`
	Type              string `json:"type"`
	URL               string `json:"url"`
	Logo              Poster `json:"logo,omitempty"`
	Group             string `json:"group,omitempty"`
	Title             string `json:"title,omitempty"`
	Announce          string `json:"announce,omitempty"`
	Subtitle          string `json:"subtitle,omitempty"`
	Description       string `json:"description,omitempty"`
	ArticleToListType string `json:"articleToListType,omitempty"`
	Issuer            string `json:"issuer"`
}

type MentionedArticle struct {
	ID                 string    `json:"id"`
	URL                string    `json:"url"`
	Title              string    `json:"title"`
	Issuer             string    `json:"issuer"`
	PublishedAt        int       `json:"publishedAt"`
	PublishedAtIso8601 time.Time `json:"publishedAtIso8601"`
	Cover              Media     `json:"cover"`
}

type Article struct {
	ID          string      `json:"id"`
	Issuer      string      `json:"issuer"`
	Type        string      `json:"type"`
	Priority    int         `json:"priority,omitempty"`
	Title       string      `json:"title"`
	Subtitle    string      `json:"subtitle"`
	Lead        string      `json:"lead,omitempty"`
	URL         string      `json:"url,omitempty"`
	Cover       Media       `json:"cover,omitempty"`
	Body        []Body      `json:"body"`
	PublishedAt int         `json:"publishedAt,omitempty"`
	ModifiedAt  int         `json:"modifiedAt"`
	References  []Reference `json:"references,omitempty"`
	Lists       []List      `json:"lists"`
	Tags        []List      `json:"tags,omitempty"`
	Authors     []Author    `json:"authors,omitempty"`
}

func (headline Headline) getMedia() []Media {
	result := make([]Media, 0)

	if headline.Cover.hasID() {
		result = append(result, headline.Cover)
	}

	for _, author := range headline.Authors {
		if author.Avatar.hasID() {
			result = append(result, author.Avatar)
		}
	}

	return result
}

func getMedia(medias []Media) []Media {
	result := make([]Media, 0)

	for _, media := range medias {
		if media.hasID() {
			result = append(result, media)
		}
	}

	return result
}

func (body Body) getMedia() []Media {
	result := make([]Media, 0)

	if hMedia := body.Headline.getMedia(); len(hMedia) > 0 {
		result = append(result, hMedia...)
	}

	for _, author := range body.Authors {
		if author.Avatar.hasID() {
			result = append(result, author.Avatar)
		}
	}

	if media := getMedia(body.Medias); len(media) > 0 {
		result = append(result, media...)
	}

	if body.Article.Cover.hasID() {
		body.Article.Cover.ArticleID = body.Article.ID
		result = append(result, body.Article.Cover)
	}

	for _, ref := range body.References {
		if ref.Cover.hasID() {
			result = append(result, ref.Cover)
		}
	}

	if photobook := getMedia(body.Photobook); len(photobook) > 0 {
		result = append(result, photobook...)
	}

	return result
}

func (article Article) GetMedia1() []Media {
	result := make([]Media, 0)
	if article.Cover.hasID() {
		result = append(result, article.Cover)
	}

	for _, body := range article.Body {
		if media := body.getMedia(); len(media) > 0 {
			result = append(result, media...)
		}
	}

	for _, ref := range article.References {
		if ref.Cover.hasID() {
			result = append(result, ref.Cover)
		}
	}

	for _, author := range article.Authors {
		if author.Avatar.hasID() {
			result = append(result, author.Avatar)
		}
	}

	fillArticleID(article, result)

	return result
}

func fillArticleID(article Article, media []Media) {
	for i := 0; i < len(media); i++ {
		if media[i].ArticleID == "" {
			media[i].ArticleID = article.ID
		}
	}
}
