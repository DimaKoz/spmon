package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
var (
	testMedia0    = Media{ID: "some_id"}
	testMedia1    = Media{ID: "some_id_1"}
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
