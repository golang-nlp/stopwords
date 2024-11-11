package stopwords

import "github.com/golang-nlp/stopwords/data"

func IsStopWord(lang, word string) bool {
	_, ok := data.Languages[lang][word]

	return ok
}
