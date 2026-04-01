package internal

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
