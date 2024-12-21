package validator

import (
	"context"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type Validator interface {
	Valid(context.Context) Evaluator
}

type Evaluator map[string]string

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (e *Evaluator) AddFieldError(key, message string) {
	if *e == nil {
		*e = make(map[string]string)
	}

	if _, exists := (*e)[key]; !exists {
		(*e)[key] = message
	}
}

func (e *Evaluator) CheckField(ok bool, key string, message string) {
	if !ok {
		e.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, max int) bool {
	return utf8.RuneCountInString(value) <= max
}

func MinChars(value string, min int) bool {
	return utf8.RuneCountInString(value) >= min
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func MinPrice(value float64) bool {
	return value > 0
}

func MinAuctionDuration(value time.Time, now time.Time) bool {
	const minAuctionDuration = 2 * time.Hour

	return value.Sub(now) >= minAuctionDuration
}
