package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
func TestHeadlineGetMedia(t *testing.T) {
	media0 := Media{ID: "some_id"}
	media1 := Media{ID: "some_id_1"}
	mediaNoID := Media{}
	tests := []struct {
		name     string
		headline Headline
		want     []Media
	}{
		{
			name: "has media",
			headline: Headline{
				Cover:   media0,
				Authors: []Author{{Avatar: media1}, {Avatar: mediaNoID}},
			},
			want: []Media{media0, media1},
		},
		{
			name:     "no media",
			headline: Headline{Cover: mediaNoID},
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
