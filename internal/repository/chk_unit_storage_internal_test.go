package repository

import (
	"testing"

	"github.com/DimaKoz/spmon/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetCheckUnit(t *testing.T) {
	testKey := "some_key"
	AddCheckUnit(testKey, nil)
	got, err := GetCheckUnit(testKey)
	require.Error(t, err)
	assert.Nil(t, got)
}

func TestAddGetCheckUnit(t *testing.T) {
	testKey := "some_key0"
	testChkUnit0 := model.NewCheckUnit("https://url.com", "articleId0", "issuer0")
	testChkUnit1 := model.NewCheckUnit("https://url.com", "articleId1", "issuer1")
	want0 := model.NewCheckUnit("https://url.com", "articleId0", "issuer0")
	want1 := model.NewCheckUnit("https://url.com", "articleId1", "issuer1")
	AddCheckUnit(testKey, nil)
	AddCheckUnit(testKey, testChkUnit0)
	got, err := GetCheckUnit(testKey)
	assert.Equal(t, want0, got)
	require.NoError(t, err)
	AddCheckUnit(testKey, testChkUnit1)
	got, err = GetCheckUnit(testKey)
	assert.Equal(t, want1, got)
	require.NoError(t, err)
}
