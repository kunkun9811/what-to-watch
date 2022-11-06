package main

import (
	"encoding/json"
	"errors"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
	"strings"
)

var validator *validator2.Validate

func init() {
	validator = validator2.New()
}

// Validate the object with rules associated with the object fields
func Validate(s interface{}) url.Values {
	errMsgs := url.Values{}

	validErrs := validator.Struct(s)

	fmt.Println(validErrs)

	if validErrs != nil {
		for _, err := range validErrs.(validator2.ValidationErrors) {
			errMsgs.Add(err.Field(), strings.ToUpper(err.ActualTag()))
		}
	}

	fmt.Println(errMsgs)

	return errMsgs
}

type ValidationError struct {
	Errors map[string][]string
}

func ValidateAndPopulateRequestBody(w http.ResponseWriter, r *http.Request, requestBody interface{}) error {
	if r.Body == nil {
		fmt.Println("Error: body is empty.")
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("bad request body")
	}

	// parse
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Println("Error: unable to deserialize request")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ValidationError{
			Errors: map[string][]string{
				"DeserializeError": {err.(*json.SyntaxError).Error()},
			},
		})
		return errors.New("bad request body")
	}

	if validationErrors := Validate(requestBody); len(validationErrors) > 0 {
		fmt.Println("Error: validationError: error occurred while validating the request.")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ValidationError{
			Errors: validationErrors,
		})
		return errors.New("bad request body")
	}

	return nil
}
