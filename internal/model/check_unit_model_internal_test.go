package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:exhaustruct
func TestNewCheckUnit(t *testing.T) {
	want := &CheckUnit{URL: "http://example.com", ArticleID: "articleId1", Issuer: "issuer0"}
	assert.Equal(t, want, NewCheckUnit("http://example.com", "articleId1", "issuer0"))
}
