package article

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
func TestMediaHasID(t *testing.T) {
	tests := []struct {
		name  string
		media Media
		want  bool
	}{
		{
			name: "HasId() == true",
			want: true,
			media: Media{
				ID: "hasID",
			},
		},
		{
			name: "no ID",
			want: false,
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			got := tUnit.media.hasID()
			assert.Equal(t, tUnit.want, got)
		})
	}
}
