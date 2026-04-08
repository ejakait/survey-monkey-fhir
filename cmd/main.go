package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	converter "github.com/ejakait/survey-monkey-fhir/internal"
)

// const jsonFile = "sample/input/survey_monkey.json"

func main() {
	// Get Survey MonkeyJSON
	// Marshall it to struct
	// Map it to FHIR Resources
	// Produce POST FHIR Transaction for Resources
	jsonFile := flag.String("json", "sample/input/survey_monkey.json", "path to the Survey Monkey JSON file")
	flag.Parse()
	file, err := os.Open(*jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValues, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	var responses []converter.Responses
	err = json.Unmarshal(byteValues, &responses)
	if err != nil {
		panic(err)
	}
	jsonFHIRConverter := converter.NewJsonFHIRConverter(
		responses,
	)

	fhirBundle, err := jsonFHIRConverter.JSONConverter()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bundle has %d entries\n", len(fhirBundle.Entry))

}
