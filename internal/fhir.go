package internal

import (
	"fmt"
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
func TranslateSurveyMonkeyStatuses(responseStatus string) r4.QuestionnaireResponseStatus {
	switch responseStatus {
	case "completed":
		return r4.QuestionnaireResponseStatusCompleted
	case "partial", "incomplete":
		return r4.QuestionnaireResponseStatusInProgress
	case "draft":
		return r4.QuestionnaireResponseStatusAmended
	default:
		return r4.QuestionnaireResponseStatusCompleted
	}
}
func (f FHIRConverter) JSONConverter() (*r4.Bundle, error) {
	bundleType := r4.BundleTypeBatch

	bundle := &r4.Bundle{}
	bundle.Type = &bundleType
	for _, response := range RemoveHTMLTags(f.SurveyMonkeyJson) {
		responseStatus := TranslateSurveyMonkeyStatuses(response.ResponseStatus)
		questionnaireResponse := r4.QuestionnaireResponse{
			Id:           &response.Id,
			Authored:     &response.DateCreated,
			Status:       &responseStatus,
			ResourceType: "QuestionnaireResponse",
		}
		for _, page := range response.Pages {
			for _, question := range page.Questions {

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

		putVerb := r4.HTTPVerbPut
		questionnaireResponseUrl := fmt.Sprintf("%s/sm-%v", questionnaireResponse.ResourceType, response.Id)
		bundle.Entry = append(bundle.Entry, r4.BundleEntry{
			Resource: &questionnaireResponse,
			Request: &r4.BundleEntryRequest{
				Method: &putVerb,
				Url:    &questionnaireResponseUrl,
			},
		})
	}

	return bundle, nil
}
