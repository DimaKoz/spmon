package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/DimaKoz/spmon/internal"
	"github.com/DimaKoz/spmon/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAddGetHs(t *testing.T) {
	type args struct {
		key string
		hs  *model.Handshake
	}
	expectedTimestamp := 1703246443
	wDir := internal.GetWD()
	filePath := fmt.Sprintf("%s/%s", wDir, "testdata/hs/sh_sputnik_intl_en.json")
	file, err := os.ReadFile(filePath)
	// println("use test file:", filePath, ", ", file)
	if !assert.NoError(t, err) {
		return
	}
	var hsTest model.Handshake
	err = json.Unmarshal(file, &hsTest)
	if !assert.NoError(t, err) {
		return
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				key: "sputnik_intl_en",
				hs:  &hsTest,
			},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			AddHs(tUnit.args.key, nil)
			got, err := GetHs(tUnit.args.key)
			assert.Error(t, err)
			assert.Nil(t, got)
			AddHs(tUnit.args.key, tUnit.args.hs)
			got, err = GetHs(tUnit.args.key)
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, expectedTimestamp, got.UpdatedAt)
		})
	}
}
