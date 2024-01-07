package model

type CheckUnit struct {
	URL              string
	Issuer           string
	ArticleID        string
	CheckTime        int64
	ServerAnswer     int
	ServerAnswerText string
	ImageSize        int64
}

func NewCheckUnit(url string, articleID string, issuer string) *CheckUnit {
	return &CheckUnit{
		URL:       url,
		Issuer:    issuer,
		ArticleID: articleID,
		ImageSize: 0, CheckTime: 0, ServerAnswer: 0, ServerAnswerText: "",
	}
}
