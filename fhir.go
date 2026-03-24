package main

import (
	"encoding/json"
	"fmt"

	r4pb "github.com/gofhir/models/r4"
)

func MapSurveyMonkeyToFHIR(responses []Responses) (*r4pb.Bundle, error) {
	bundle := &r4pb.Bundle{}
	for _, response := range responses {
		questionnaireResponse := r4pb.QuestionnaireResponse{
			Id: &response.Id,
		}
		for _, page := range response.Pages {
			for _, question := range page.Questions {

				question.Heading = RemoveHTMLTags(question.Heading)
				for i, answer := range question.Answers {
					question.Answers[i].SimpleText = RemoveHTMLTags(answer.SimpleText)
				}

				questionnaireResponseItem := r4pb.QuestionnaireResponseItem{
					LinkId: &question.Id,
				}

				questionnaireResponseItem.Text = &question.Heading

				for _, answer := range question.Answers {
					splitQuestionAnswers, err := SeparateSimpleTextResponse(answer)
					if err != nil {
						continue
					}
					if len(splitQuestionAnswers) == 2 {
						questionnaireResponseItem.Answer = append(questionnaireResponseItem.Answer, r4pb.QuestionnaireResponseItemAnswer{
							Id:          &splitQuestionAnswers[0],
							ValueString: &splitQuestionAnswers[1],
						})
					} else {
						questionnaireResponseItem.Answer = append(questionnaireResponseItem.Answer, r4pb.QuestionnaireResponseItemAnswer{
							ValueString: &splitQuestionAnswers[0],
						})
					}
				}

				questionnaireResponse.Item = append(questionnaireResponse.Item, questionnaireResponseItem)
			}
			bundle.Entry = append(bundle.Entry, r4pb.BundleEntry{Resource: &questionnaireResponse})
		}
	}
	fmt.Printf("Bundle has %d entries\n", len(bundle.Entry))
	data, err := json.MarshalIndent(bundle, "", "  ")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	return bundle, nil
}
