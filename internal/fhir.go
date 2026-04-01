package internal

import (
	log "log/slog"

	r4 "github.com/gofhir/models/r4"
)

type Converter interface {
	JsonConverter(*FHIRConverter) (*r4.Bundle, error)
	XMLConverter(*FHIRConverter) (*r4.Bundle, error)
}

type FHIRConverter struct {
	FHIRVersion      string
	SurveyMonkeyJson []Responses
}

func NewJsonFHIRConverter(s []Responses) *FHIRConverter {
	return &FHIRConverter{
		FHIRVersion:      "4.0.1",
		SurveyMonkeyJson: s,
	}
}

func (f FHIRConverter) JSONConverter() (*r4.Bundle, error) {
	bundle := &r4.Bundle{}
	for _, response := range f.SurveyMonkeyJson {
		questionnaireResponse := r4.QuestionnaireResponse{
			Id: &response.Id,
		}
		for _, page := range response.Pages {
			for _, question := range page.Questions {

				question.Heading = RemoveHTMLTags(question.Heading)
				for i, answer := range question.Answers {
					question.Answers[i].SimpleText = RemoveHTMLTags(answer.SimpleText)
				}

				questionnaireResponseItem := r4.QuestionnaireResponseItem{
					LinkId: &question.Id,
				}

				questionnaireResponseItem.Text = &question.Heading

				for _, answer := range question.Answers {
					splitQuestionAnswers, err := SeparateSimpleTextResponse(answer)
					if err != nil {
						log.Error("error separating simple text response", "error", err)
						continue
					}
					if len(splitQuestionAnswers) == 2 {
						questionnaireResponseItem.Answer = append(questionnaireResponseItem.Answer, r4.QuestionnaireResponseItemAnswer{
							Id:          &splitQuestionAnswers[0],
							ValueString: &splitQuestionAnswers[1],
						})
					} else {
						questionnaireResponseItem.Answer = append(questionnaireResponseItem.Answer, r4.QuestionnaireResponseItemAnswer{
							ValueString: &splitQuestionAnswers[0],
						})
					}
				}

				questionnaireResponse.Item = append(questionnaireResponse.Item, questionnaireResponseItem)
			}
		}

		bundle.Entry = append(bundle.Entry, r4.BundleEntry{Resource: &questionnaireResponse})
	}

	return bundle, nil
}
