package repository

import (
	"testing"

	"github.com/DimaKoz/spmon/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testChkUnit0 = model.NewCheckUnit("https://url.com", "articleId0", "issuer0")
	testChkUnit1 = model.NewCheckUnit("https://url.com", "articleId1", "issuer1")
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

func TestClearCheckUnit(t *testing.T) {
	testKey := "some_key0"

	AddCheckUnit(testKey, testChkUnit0)

	require.NotEmpty(t, chkUnitStorage.storage)
	ClearUnitStorage()

	assert.Empty(t, chkUnitStorage.storage)
}

func TestGetAllCheckUnits(t *testing.T) {
	require.Empty(t, GetAllChkUnits())

	want := []model.CheckUnit{*testChkUnit0, *testChkUnit1}

	AddCheckUnit("some_key0", testChkUnit0)
	AddCheckUnit("testKey", testChkUnit1)
	got := GetAllChkUnits()

	assert.NotEmpty(t, got)
	assert.Equal(t, want, got)
}
