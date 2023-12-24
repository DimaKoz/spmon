package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"spmon/internal"
	"spmon/internal/model"
)

func TestAddGetHs(t *testing.T) {
	type args struct {
		key string
		hs  *model.Handshake
	}
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
		tt := tCase
		t.Run(tt.name, func(t *testing.T) {
			AddHs(tt.args.key, nil)
			got, err := GetHs(tt.args.key)
			assert.Error(t, err)
			assert.Nil(t, got)
			AddHs(tt.args.key, tt.args.hs)
			got, err = GetHs(tt.args.key)
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, 1703246443, got.UpdatedAt)
		})
	}
}
