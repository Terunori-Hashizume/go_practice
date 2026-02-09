package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		var emptyFldErr EmptyFieldError
		errMsgs := make([]string, 0, 10)
		missingFields := make([]string, 0, 5)
		for _, err := range ValidateEmployee(emp).Unwrap() {
			if errors.Is(err, InvalidIdError) {
				errMsgs = append(errMsgs, InvalidIdError.Error())
			} else if errors.As(err, &emptyFldErr) {
				missingFields = append(missingFields, strings.Split(err.Error(), "missing ")[1])
			}
		}
		if len(missingFields) > 0 {
			errMsgs = append(errMsgs, "missing "+strings.Join(missingFields, ", "))
		}
		if len(errMsgs) == 0 {
			fmt.Printf("record %d: %+v\n", count, emp)
		} else {
			fmt.Printf("record %d: %+v error: %s\n", count, emp, strings.Join(errMsgs, " "))
		}
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID        = regexp.MustCompile(`\w{4}-\d{3}`)
	InvalidIdError = errors.New("invalid ID")
)

type ValidationError struct {
	Errors []error
}

func (v ValidationError) Error() string {
	if len(v.Errors) == 0 {
		return ""
	} else {
		return errors.Join(v.Errors...).Error()
	}
}

func (v ValidationError) Unwrap() []error {
	return v.Errors
}

type EmptyFieldError struct {
	FieldName string
}

func (err EmptyFieldError) Error() string {
	return fmt.Sprintf("missing %s", err.FieldName)
}

func ValidateEmployee(e Employee) ValidationError {
	errs := make([]error, 5)
	if len(e.ID) == 0 {
		errs = append(errs, EmptyFieldError{FieldName: "ID"})
	}
	if !validID.MatchString(e.ID) {
		errs = append(errs, InvalidIdError)
	}
	if len(e.FirstName) == 0 {
		errs = append(errs, EmptyFieldError{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		errs = append(errs, EmptyFieldError{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		errs = append(errs, EmptyFieldError{FieldName: "Title"})
	}
	return ValidationError{Errors: errs}
}
