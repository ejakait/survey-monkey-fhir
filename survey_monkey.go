package main

import (
	"fmt"
	"strings"
)

type QuestionFamily string

const (
	SingleChoice QuestionFamily = "single_choice"
	Matrix       QuestionFamily = "matrix"
)

func (qf QuestionFamily) String() string {
	return string(qf)
}

type Questions struct {
	id      string
	answers []Answers
	family  QuestionFamily
	subtype string
	heading string
}
type Answers struct {
	choice_id   string
	row_id      string
	simple_text string
}
type Pages struct {
	id        string
	questions []Questions
}
type Responses struct {
	id              string
	recipient_id    string
	response_status string
	first_name      string
	last_name       string
	email_address   string
	survey_id       string
	collector_id    string
	language        string
	date_created    string
	date_modified   string
	analyze_url     string
	pages           []Answers
}

func (a Answers) SeperateSimpleTextResponse() ([]string, error) {

	var questionAnswers []string
	questionAnswers = strings.Split(a.simple_text, "|")

	if len(questionAnswers) > 2 {
		return questionAnswers, fmt.Errorf("more than 2 sections were found in the simple text: %s", a.simple_text)

	}
	strippedQuestionAnswers := make([]string, 0, len(questionAnswers))
	for _, qa := range questionAnswers {

		strippedQuestionAnswers = append(strippedQuestionAnswers, strings.TrimSpace(qa))
	}
	return strippedQuestionAnswers, nil
}
