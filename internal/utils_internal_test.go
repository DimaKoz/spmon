package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWD(t *testing.T) {
	tests := []struct {
		name          string
		want          string
		notEmptyCheck bool
	}{
		{name: "spmon 1st time", want: "spmon", notEmptyCheck: false},
		{name: "spmon 2nd time", want: "spmon", notEmptyCheck: true},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			got := GetWD()
			if tUnit.notEmptyCheck {
				assert.NotEmpty(t, workDir)
			}
			assert.Contains(t, got, tUnit.want)
		})
	}
}
