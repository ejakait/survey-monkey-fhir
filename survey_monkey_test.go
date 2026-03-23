package main

import (
	"log"
	"slices"
	"testing"
)

func TestSeparateSimpleTextResponse(t *testing.T) {
	a := Answers{
		choice_id:   "12345",
		row_id:      "54321",
		simple_text: "Question | Text Answer",
	}
	testCases := []struct {
		name       string
		simpleText string
		want       []string
	}{
		{"Basic Simple Text", a.simple_text, []string{"Question", "Text Answer"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := a.SeparateSimpleTextResponse()
			log.Printf("Found question %s and answer %s", got[0], got[1])
			if !slices.Equal(got, tt.want) {
				t.Errorf("SeparateSimpleTextResponse(%v) = %v want %v", a, got, tt.want)
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

	log.Printf("Got=%s", got)
	if got != want {
		t.Errorf("RemoveHTMLTags got=%s, want=%s", got, want)
	}
}
