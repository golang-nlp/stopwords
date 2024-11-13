package stopwords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStopWord(t *testing.T) {
	tt := []struct {
		lang string
		word string
		want bool
	}{
		{"fr", "au", true},
		{"fr", "aux", true},
		{"fr", "avec", true},
		{"fr", "ce", true},
		{"fr", "ces", true},
		{"fr", "Voiture", false},
		{"bad", "bad", false},
		{"en", "the", true},
		{"en", "car", false},
	}

	for _, tc := range tt {
		t.Run(tc.word, func(t *testing.T) {
			got := IsStopWord(tc.lang, tc.word)
			assert.Equal(t, tc.want, got)
		})
	}
}
