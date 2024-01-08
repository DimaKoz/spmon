package helper

import (
	"reflect"
	"testing"

	"github.com/DimaKoz/spmon/internal/model"
	"github.com/DimaKoz/spmon/internal/model/article"
	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
func TestGetCheckUnits(t *testing.T) {
	wantCheckUnit0 := model.CheckUnit{
		URL:       "https://example.com",
		ArticleID: "ArticleId0",
		Issuer:    "Issuer0",
	}

	type args struct {
		media article.Media
	}
	tests := []struct {
		name string
		args args
		want []model.CheckUnit
	}{
		{
			name: "EmptyArticleId",
			args: args{media: article.Media{}},
			want: []model.CheckUnit{},
		},
		{
			name: "EmptyIssuer",
			args: args{media: article.Media{}},
			want: []model.CheckUnit{},
		},
		{
			name: "MediaTypeImage",
			args: args{media: article.Media{
				ArticleID:   "ArticleId0",
				Issuer:      "Issuer0",
				Type:        "image",
				Poster:      article.Poster{URL: "https://example.com"},
				SourceURL:   "https://example.com",
				DownloadURL: "https://example.com",
			}},
			want: []model.CheckUnit{wantCheckUnit0, wantCheckUnit0, wantCheckUnit0},
		},
		{
			name: "MediaTypeVideo",
			args: args{media: article.Media{
				ArticleID:   "ArticleId0",
				Issuer:      "Issuer0",
				Type:        "video",
				Poster:      article.Poster{URL: "https://example.com"},
				SourceURL:   "https://example.com",
				DownloadURL: "https://example.com",
			}},
			want: []model.CheckUnit{wantCheckUnit0},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			if got := GetCheckUnits(tUnit.args.media); !reflect.DeepEqual(got, tUnit.want) {
				assert.Equalf(t, tUnit.want, got, "GetCheckUnits() = %v, want %v", got, tUnit.want)
			}
		})
	}
}
