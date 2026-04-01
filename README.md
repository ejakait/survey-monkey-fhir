# Survey Monkey to FHIR R4 Converter

[![CI](https://github.com/ejakait/survey-monkey-fhir/actions/workflows/ci.yml/badge.svg)](https://github.com/ejakait/survey-monkey-fhir/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/ejakait/survey-monkey-fhir/branch/main/graph/badge.svg)](https://codecov.io/gh/github.com/ejakait/survey-monkey-fhir/)

Convert Survey Monkey survey responses to FHIR R4 compliant resources for integration with FHIR servers.

## Overview

This service transforms Survey Monkey API survey data into FHIR R4 resources:
- **QuestionnaireResponse** - Completed survey responses bundled for FHIR transaction

## Installation

```bash
go mod download
```

## Usage

```bash
go run cmd/main.go
```

The converter reads from `sample/input/survey_monkey.json` by default and outputs a FHIR Bundle.

## Features

- Converts Survey Monkey JSON responses to FHIR R4 QuestionnaireResponse resources
- Generates FHIR Bundle suitable for POSTing to a FHIR server
- HTML tag sanitization using bluemonday
- Batch processing of survey responses

## Project Structure

```
.
├── cmd/main.go              # Entry point
├── internal/
│   ├── fhir.go             # FHIR conversion logic
│   ├── models.go           # Survey Monkey data models
│   ├── survey_cleaner.go   # HTML sanitization utilities
│   └── *_test.go           # Unit tests
├── sample/input/           # Sample Survey Monkey data
└── mapping.yml             # Field mapping configuration (future)
```

## FHIR Output

The converter produces a FHIR Bundle containing QuestionnaireResponse resources:

```bash
curl -X POST <fhir-server>/ \
  -H "Content-Type: application/fhir+json" \
  -d @output.json
```

## Requirements

- Go 1.21+
- Survey Monkey API data export
- FHIR R4 compatible server
