package internal

import (
	"fmt"
	"strings"

	htmlsanitizer "github.com/microcosm-cc/bluemonday"
)

var sanitizer = htmlsanitizer.StripTagsPolicy()

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

func RemoveHTMLTags(r []Responses) []Responses {
	replacer := strings.NewReplacer(
		`\u003cbr\u003e`, "\n",
		`\u003c`, "<",
		`\u003e`, ">",
		`\u0026`, "&",
	)
	for i, response := range r {
		r[i].FirstName = strings.TrimSpace(sanitizer.Sanitize(response.FirstName))
		r[i].LastName = strings.TrimSpace(sanitizer.Sanitize(response.LastName))
		r[i].EmailAddress = strings.TrimSpace(sanitizer.Sanitize(response.EmailAddress))
		for j, page := range response.Pages {
			for k, question := range page.Questions {
				r[i].Pages[j].Questions[k].Heading = sanitizer.Sanitize(replacer.Replace(question.Heading))
				for l, answer := range question.Answers {
					if answer.SimpleText != "" {
						r[i].Pages[j].Questions[k].Answers[l].SimpleText = strings.TrimSpace(sanitizer.Sanitize(replacer.Replace(answer.SimpleText)))
					}
				}
			}
		}
	}
	return r
}
