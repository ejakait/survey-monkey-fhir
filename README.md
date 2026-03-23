# Survey Monkey to FHIR R4 Converter

Convert Survey Monkey survey responses to FHIR R4 compliant resources for integration with FHIR servers.

## Overview

This service transforms Survey Monkey API survey data into FHIR R4 resources including:
- **Questionnaire** - Survey structure definition
- **QuestionnaireResponse** - Completed survey responses
- **Patient** - Survey respondent information
- **Observation** - Individual answer records

## Installation

```bash
go install
```

## Usage

```bash
go run main.go --survey-monkey-data=<path> --output=<path>
```

## Configuration

Configure survey-to-FHIR mappings in `mapping.yml`.

## FHIR Output

Generates a FHIR Bundle transaction for POSTing directly to a FHIR server:

```bash
curl -X POST <fhir-server>/ \
  -H "Content-Type: application/fhir+json" \
  -d @output.json
```

## Requirements

- Go 1.25+
- Survey Monkey API data export
- FHIR R4 compatible server
