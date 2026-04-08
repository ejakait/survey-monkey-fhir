package survey

import (
	"fmt"
	"strings"

	htmlsanitizer "github.com/microcosm-cc/bluemonday"
)

var sanitizer = htmlsanitizer.StripTagsPolicy()

func StripHTML(s string) string {
	replacer := strings.NewReplacer(
		`\u003cbr\u003e`, "\n",
		`\u003c`, "<",
		`\u003e`, ">",
		`\u0026`, "&",
	)
	return sanitizer.Sanitize(replacer.Replace(s))
}

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
	for i, response := range r {
		r[i].FirstName = strings.TrimSpace(StripHTML(response.FirstName))
		r[i].LastName = strings.TrimSpace(StripHTML(response.LastName))
		r[i].EmailAddress = strings.TrimSpace(StripHTML(response.EmailAddress))
		for j, page := range response.Pages {
			for k, question := range page.Questions {
				r[i].Pages[j].Questions[k].Heading = StripHTML(question.Heading)
				for l, answer := range question.Answers {
					if answer.SimpleText != "" {
						r[i].Pages[j].Questions[k].Answers[l].SimpleText = strings.TrimSpace(StripHTML(answer.SimpleText))
					}
				}
			}
		}
	}
	return r
}
