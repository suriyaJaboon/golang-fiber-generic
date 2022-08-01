package handles

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var val = validator.New()

type ErrorValidation struct {
	Tag     string `json:"tag"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func Validation(i any) []*ErrorValidation {
	var errors []*ErrorValidation
	if err := val.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var element ErrorValidation
			element.Tag = e.Tag()
			element.Name = strings.ToLower(e.Field())
			element.Message = "Error:Field validation for '" + e.Field() + "' failed on the '" + e.Tag() + "' tag"
			errors = append(errors, &element)

			fmt.Println(e.Namespace())
			fmt.Println(e.Field())
			fmt.Println(e.StructNamespace())
			fmt.Println(e.StructField())
			fmt.Println(e.Tag())
			fmt.Println(e.ActualTag())
			fmt.Println(e.Kind())
			fmt.Println(e.Type())
			fmt.Println(e.Value())
			fmt.Println(e.Param())
			fmt.Println()
		}
	}

	return errors
}
