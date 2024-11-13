# Stopwords for Golang ![Last release](https://img.shields.io/github/release/golang-nlp/stopwords.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/golang-nlp/stopwords)](https://goreportcard.com/report/github.com/golang-nlp/stopwords)

| Branch | Status                                                                                                                                                | Coverage                                                                                                                                         |
| ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| master | [![Go](https://github.com/golang-nlp/stopwords/actions/workflows/go.yml/badge.svg)](https://github.com/golang-nlp/stopwords/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/golang-nlp/stopwords/master.svg)](https://coveralls.io/github/golang-nlp/stopwords?branch=master) |

```sh
go get -u github.com/golang-nlp/stopwords
```

```go
import (
    "fmt"

    "github.com/golang-nlp/stopwords"
)

func main() {
    fmt.Print(stopwords.IsStopWord("fr", "avec")) // true
}

```

## License

stopwords is licensed under [the MIT license](LICENSE.md).
