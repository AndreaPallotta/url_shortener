package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {

	links := map[string]string{
		"https://www.google.com": "2dDEQAS1",
		"https://www.amazon.com": "CJatrA6W",
		"https://github.com":     "cCd1qpQg",
	}

	for fullLink, expectedShort := range links {
		actualShort := GenerateShortLink(fullLink, UserId)
		assert.Equal(t, actualShort, expectedShort)
	}
}
