package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
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
	Id      string         `json:"id"`
	Answers []Answers      `json:"answers"`
	Family  QuestionFamily `json:"family"`
	Subtype string         `json:"subtype"`
	Heading string         `json:"heading"`
}
type Answers struct {
	ChoiceId   string `json:"choice_id"`
	RowId      string `json:"row_id"`
	SimpleText string `json:"simple_text"`
}
type Pages struct {
	Id        string      `json:"id"`
	Questions []Questions `json:"questions"`
}
type Responses struct {
	Id             string  `json:"id"`
	RecipientId    string  `json:"recipient_id"`
	ResponseStatus string  `json:"response_status"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	EmailAddress   string  `json:"email_address"`
	SurveyId       string  `json:"survey_id"`
	CollectorId    string  `json:"collector_id"`
	Language       string  `json:"language"`
	DateCreated    string  `json:"date_created"`
	DateModified   string  `json:"date_modified"`
	AnalyzeUrl     string  `json:"analyze_url"`
	Pages          []Pages `json:"pages"`
}

const jsonFile = "./data.json"

func main() {
	// Get Survey MonkeyJSON
	// Marshall it to struct
	// Map it to FHIR Resources
	// Produce POST FHIR Transaction for Resources

	file, err := os.Open(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var responses []Responses
	err = json.Unmarshal(byteValues, &responses)
	if err != nil {
		panic(err)
	}

	fhirBundle, err := MapSurveyMonkeyToFHIR(responses)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fhirBundle)

}
