package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
var (
	testMedia0    = Media{ID: "some_id"}
	testMedia1    = Media{ID: "some_id_1"}
	testMedia2    = Media{ID: "some_id_2"}
	testMedia3    = Media{ID: "some_id_3"}
	testMedia4    = Media{ID: "some_id_4"}
	testMedia5    = Media{ID: "some_id_5"}
	testMediaNoID = Media{}
)

//nolint:exhaustruct
func TestHeadlineGetMedia(t *testing.T) {
	tests := []struct {
		name     string
		headline Headline
		want     []Media
	}{
		{
			name: "has media",
			headline: Headline{
				Cover:   testMedia0,
				Authors: []Author{{Avatar: testMedia1}, {Avatar: testMediaNoID}},
			},
			want: []Media{testMedia0, testMedia1},
		},
		{
			name:     "no media",
			headline: Headline{Cover: testMediaNoID},
			want:     []Media{},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			got := tUnit.headline.getMedia()
			assert.Equal(t, tUnit.want, got)
		})
	}
}

func TestGetMediaFromSliceMedia(t *testing.T) {
	tests := []struct {
		name  string
		media []Media
		want  []Media
	}{
		{
			name:  "all_cases",
			media: []Media{testMedia0, testMediaNoID, testMedia1},
			want:  []Media{testMedia0, testMedia1},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			assert.Equal(t, tUnit.want, getMedia(tUnit.media), "getMedia(%v)", tUnit.media)
		})
	}
}

//nolint:exhaustruct
func TestBodyGetMedia(t *testing.T) {
	body := Body{
		Headline: Headline{
			Cover:   testMedia0,
			Authors: []Author{{Avatar: testMedia1}},
		},
		Authors:    []Author{{Avatar: testMedia1}},
		Medias:     []Media{testMedia2, testMediaNoID},
		Article:    MentionedArticle{ID: "MentionedArticle0", URL: "http://example.com", Cover: testMedia3},
		References: []Reference{{ID: "Reference0", URL: "http://example.com", Cover: testMedia4}},
		Photobook:  []Media{testMedia5, testMediaNoID, testMedia3},
	}

	tests := []struct {
		name string
		body Body
		want []Media
	}{
		{
			name: "body with media",
			body: body,
			want: []Media{
				testMedia0,
				testMedia1,
				testMedia1,
				testMedia2,
				{ID: "some_id_3", ArticleID: "MentionedArticle0"},
				testMedia4,
				testMedia5,
				testMedia3,
			},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			got := tUnit.body.getMedia()
			assert.Equal(t, tUnit.want, got)
		})
	}
}

//nolint:exhaustruct
func TestFillArticleID(t *testing.T) {
	tests := []struct {
		name    string
		article Article
		media   []Media
		want    []Media
	}{
		{
			name:    "with id and without",
			article: Article{ID: "articleId0"},
			media: []Media{
				{ID: "test0"},
				{ID: "test1", ArticleID: "articleId1"},
				{ID: "test2"},
				{ID: "test3"},
				{ID: "test4"},
				{ID: "test5"},
				{ID: "test6"},
				{ID: "test7"},
				{ID: "test8"},
				{ID: "test9"},
				{ID: "test10"},
			},
			want: []Media{
				{ID: "test0", ArticleID: "articleId0"},
				{ID: "test1", ArticleID: "articleId1"},
				{ID: "test2", ArticleID: "articleId0"},
				{ID: "test3", ArticleID: "articleId0"},
				{ID: "test4", ArticleID: "articleId0"},
				{ID: "test5", ArticleID: "articleId0"},
				{ID: "test6", ArticleID: "articleId0"},
				{ID: "test7", ArticleID: "articleId0"},
				{ID: "test8", ArticleID: "articleId0"},
				{ID: "test9", ArticleID: "articleId0"},
				{ID: "test10", ArticleID: "articleId0"},
			},
		},
	}

	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			fillArticleID(tUnit.article, tUnit.media)
			assert.Equal(t, tUnit.want, tUnit.media)
		})
	}
}
