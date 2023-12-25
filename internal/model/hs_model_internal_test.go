package model

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DimaKoz/spmon/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:exhaustruct
func TestHandshakeGetArticleBlocks(t *testing.T) {
	wDir := internal.GetWD()
	filePath := fmt.Sprintf("%s/%s", wDir, "testdata/hs/hs_sputnik_intl_en.json")
	file, err := os.ReadFile(filePath)

	require.NoError(t, err)
	var hsTest Handshake
	err = json.Unmarshal(file, &hsTest)
	require.NoError(t, err)

	tests := []struct {
		name            string
		sections        []Section
		wantBlockNumber int
	}{
		{
			name:            "ok",
			sections:        hsTest.Sections,
			wantBlockNumber: 14,
		},
		{
			name:            "no sections",
			sections:        []Section{},
			wantBlockNumber: 0,
		},
		{
			name:            "no feeds",
			sections:        []Section{{Feeds: []Feed{}}},
			wantBlockNumber: 0,
		},
		{
			name:            "no blocks",
			sections:        []Section{{Feeds: []Feed{{Blocks: []Block{}}}}},
			wantBlockNumber: 0,
		},

		{
			name: "wrong type",
			sections: []Section{{Feeds: []Feed{
				{
					Blocks: []Block{
						{
							ContentType: "wrong type",
						},
					},
				},
			}}},
			wantBlockNumber: 0,
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			handshake := &Handshake{
				Sections: tUnit.sections,
			}
			assert.Len(t, handshake.GetArticleBlocks(), tUnit.wantBlockNumber)
		})
	}
}
