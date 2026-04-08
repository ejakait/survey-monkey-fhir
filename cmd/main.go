package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	fhir "github.com/ejakait/survey-monkey-fhir/internal/fhir"
	"github.com/ejakait/survey-monkey-fhir/internal/survey"
)

func main() {
	jsonFile := flag.String("json", "sample/input/survey_monkey.json", "path to the Survey Monkey JSON file")
	flag.Parse()
	file, err := os.Open(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var responses []survey.Responses
	err = json.Unmarshal(byteValues, &responses)
	if err != nil {
		log.Fatal(err)
	}

	jsonFHIRConverter := fhir.NewJsonFHIRConverter(responses)

	fhirBundle, err := jsonFHIRConverter.JsonConverter()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bundle has %d entries\n", len(fhirBundle.Entry))
	formatted, _ := json.MarshalIndent(fhirBundle, "", "  ")

	err = os.WriteFile("sample/output/fhir_bundle.json", formatted, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
