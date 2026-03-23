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
