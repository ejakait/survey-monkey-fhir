package main

import (
	"fmt"
	"strings"

	htmlsanitizer "github.com/microcosm-cc/bluemonday"
)

func SeparateSimpleTextResponse(a Answers) ([]string, error) {

	var questionAnswers []string
	questionAnswers = strings.Split(a.SimpleText, "|")

	if len(questionAnswers) > 2 {
		return questionAnswers, fmt.Errorf("more than 2 sections were found in the simple text: %s", a.SimpleText)

	}
	strippedQuestionAnswers := make([]string, 0, len(questionAnswers))
	for _, qa := range questionAnswers {

		strippedQuestionAnswers = append(strippedQuestionAnswers, strings.TrimSpace(qa))
	}
	return strippedQuestionAnswers, nil
}

func RemoveHTMLTags(target string) string {
	sanitizer := htmlsanitizer.StripTagsPolicy()
	target = strings.TrimSpace(sanitizer.Sanitize(target))
	r := strings.NewReplacer(
		`\u003cbr\u003e`, "\n",
		`\u003c`, "<",
		`\u003e`, ">",
		`\u0026`, "&",
	)
	target = r.Replace(target)
	return target
}
