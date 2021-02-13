package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Hello hello
func Hello() string {
	return quote.Hello()
}

// Proverb proverb
func Proverb() string {
	return quoteV3.Concurrency()
}
