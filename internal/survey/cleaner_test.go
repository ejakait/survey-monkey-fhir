package survey

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
			a := Answers{SimpleText: tt.simpleText}
			got, err := SeparateSimpleTextResponse(a)

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

func TestStripHTML(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "HTML tags stripped",
			input: "<span style=\"font-size: 14pt;\"><strong>Testing</strong></span>",
			want:  "Testing",
		},
		{
			name:  "No HTML returns unchanged",
			input: "Plain text",
			want:  "Plain text",
		},
		{
			name:  "Escaped br tag converted",
			input: "Line 1\u003cbr\u003eLine 2",
			want:  "Line 1Line 2",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := StripHTML(tt.input)
			if got != tt.want {
				t.Errorf("StripHTML() = %q, want %q", got, tt.want)
			}
		})
	}
}
