package helper

import (
	"github.com/DimaKoz/spmon/internal/model"
	"github.com/DimaKoz/spmon/internal/model/article"
)

func GetCheckUnits(media article.Media) []model.CheckUnit {
	result := make([]model.CheckUnit, 0)
	if media.ArticleID == "" || media.Issuer == "" {
		return result
	}

	if media.Type == "image" {
		if media.SourceURL != "" {
			sourceCheckUnit := model.NewCheckUnit(media.SourceURL, media.ArticleID, media.Issuer)
			result = append(result, *sourceCheckUnit)
		}
		if media.DownloadURL != "" {
			downloadCheckUnit := model.NewCheckUnit(media.DownloadURL, media.ArticleID, media.Issuer)
			result = append(result, *downloadCheckUnit)
		}
	}
	if media.Poster.URL != "" {
		posterCheckUnit := model.NewCheckUnit(media.Poster.URL, media.ArticleID, media.Issuer)
		result = append(result, *posterCheckUnit)
	}

	return result
}
