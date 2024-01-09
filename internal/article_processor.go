package internal

import (
	"github.com/DimaKoz/spmon/internal/helper"
	"github.com/DimaKoz/spmon/internal/model"
	"github.com/DimaKoz/spmon/internal/model/article"
	"github.com/DimaKoz/spmon/internal/repository"
)

func UpdateCheckUnitsRepo(article article.Article) {
	media := article.GetMedia1()
	var item model.CheckUnit
	for _, mediaItem := range media {
		chkUnits := helper.GetCheckUnits(mediaItem)
		for _, chkUnitItem := range chkUnits {
			item = chkUnitItem
			repository.AddCheckUnit(chkUnitItem.URL, &item)
		}
	}
}
