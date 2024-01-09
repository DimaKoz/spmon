package internal

import (
	"testing"

	"github.com/DimaKoz/spmon/internal/model"
	"github.com/DimaKoz/spmon/internal/model/article"
	"github.com/DimaKoz/spmon/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:exhaustruct
var (
	link00 = "https://mobileimg.img.ria.ru/image?etag=3814079297&" +
		"id=1919134947&issuer=rian&_sig=PShPCXlT53hfBna_sKJppg"
	link01 = "https://mobileimg.img.ria.ru/image?etag=3814079297&h=996&" +
		"id=1919134947&issuer=rian&s=fit&w=996&_sig=BjhpcA2Z8-oQIvGjnHcNyw"
	link02 = "https://mobileimg.img.ria.ru/image?etag=3814079297&" +
		"id=1919134947&issuer=rian&s=fit&w=640&_sig=jeK_5jv6x8pOIQ7VUF_MmA"

	link10 = "https://mobileimg.img.ria.ru/image?etag=2376645882&" +
		"id=1919146348&issuer=rian&_sig=fzn8l8HKNw_ofcBenyFmbA"
	link11 = "https://mobileimg.img.ria.ru/image?etag=2376645882&h=996&" +
		"id=1919146348&issuer=rian&s=fit&w=996&_sig=86S9z75G4vBuNfuJEtBTrQ"
	link12 = "https://mobileimg.img.ria.ru/image?etag=2376645882&" +
		"id=1919146348&issuer=rian&s=fit&w=640&_sig=0-8Lrs6ZyO_A29UxwsmPuA"

	testMediaRian0 = article.Media{
		ID:          "1919134947",
		Type:        "image",
		Issuer:      "rian",
		Copyright:   "фото очевидца",
		SourceURL:   link00,
		DownloadURL: link01,
		Poster: article.Poster{
			URL: link02,
		},
	}

	testMediaRian1 = article.Media{
		ID:          "1919146348",
		Type:        "image",
		Issuer:      "rian",
		Copyright:   "ГУ МЧС России по Белгородской области",
		SourceURL:   link10,
		DownloadURL: link11,
		Poster: article.Poster{
			URL: link12,
		},
	}

	testHeadlineRian0 = article.Headline{
		Title: "Headline Rian Title0",
		Cover: testMediaRian1,
	}

	testPublicationTime = article.PublicationDateTime{
		Unix: 1703947740,
	}

	testArticleRianBody0 = []article.Body{
		{Headline: testHeadlineRian0},
		{PublicationDateTime: testPublicationTime},
	}
)

func TestUpdateCheckUnitsRepo(t *testing.T) {
	testArticleRian0 := article.Article{ //nolint:exhaustruct
		ID:     "1919152045",
		Issuer: "rian",
		Cover:  testMediaRian0,
		Body:   testArticleRianBody0,
	}

	type args struct {
		article article.Article
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		want    []model.CheckUnit
	}{
		{
			name: "RianArticle0",
			args: args{
				testArticleRian0,
			},
			wantLen: 6,
			want: []model.CheckUnit{
				*model.NewCheckUnit(link12, testArticleRian0.ID, testArticleRian0.Issuer),
				*model.NewCheckUnit(link11, testArticleRian0.ID, testArticleRian0.Issuer),
				*model.NewCheckUnit(link10, testArticleRian0.ID, testArticleRian0.Issuer),
				*model.NewCheckUnit(link00, testArticleRian0.ID, testArticleRian0.Issuer),
				*model.NewCheckUnit(link01, testArticleRian0.ID, testArticleRian0.Issuer),
				*model.NewCheckUnit(link02, testArticleRian0.ID, testArticleRian0.Issuer),
			},
		},
	}
	for _, tCase := range tests {
		tCase := tCase
		t.Run(tCase.name, func(t *testing.T) {
			repository.ClearUnitStorage()
			emptyChkUnits := repository.GetAllChkUnits()
			require.Empty(t, emptyChkUnits)
			UpdateCheckUnitsRepo(tCase.args.article)
			got := repository.GetAllChkUnits()
			assert.Len(t, got, tCase.wantLen)
			for _, wantItem := range tCase.want {
				assert.Contains(t, got, wantItem)
			}
		})
	}
}
