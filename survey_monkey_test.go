package main

import (
	"errors"
	"slices"
	"testing"
)

func TestSeparateSimpleTextResponse(t *testing.T) {
	testCases := []struct {
		name       string
		simpleText string
		want       []string
		wantErr    error
	}{
		{
			name:       "Basic Simple Text",
			simpleText: "Question | Text Answer",
			want:       []string{"Question", "Text Answer"},
			wantErr:    nil,
		},
		{
			name:       "Single Value No Delimiter",
			simpleText: "Single Value",
			want:       []string{"Single Value"},
			wantErr:    nil,
		},
		{
			name:       "With Extra Whitespace",
			simpleText: "  Question  |  Answer  ",
			want:       []string{"Question", "Answer"},
			wantErr:    nil,
		},
		{
			name:       "Too Many Sections",
			simpleText: "Q | A | Extra",
			want:       []string{"Q ", " A ", " Extra"},
			wantErr:    errors.New("more than 2 sections were found in the simple text: Q | A | Extra"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			a := Answers{simple_text: tt.simpleText}
			got, err := a.SeparateSimpleTextResponse()

			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuestionFamilyString(t *testing.T) {
	testCases := []struct {
		qf   QuestionFamily
		want string
	}{
		{SingleChoice, "single_choice"},
		{Matrix, "matrix"},
		{QuestionFamily("custom"), "custom"},
	}

	for _, tt := range testCases {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.qf.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRemoveHTMLTags(t *testing.T) {
	question := Questions{
		heading: "\u003Cspan style=\"font-size: 14pt; font-family: arial, helvetica, sans-serif;\"\u003E\u003Cstrong\u003ETesting\u003C/strong\u003E\u003C/span\u003E",
	}
	got := question.RemoveHTMLTags()
	want := "Testing"

	if got != want {
		t.Errorf("RemoveHTMLTags got=%s, want=%s", got, want)
	}
}
