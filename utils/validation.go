package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateStruct(c *gin.Context, s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			if field.Interface() == "" {
				fieldName := strings.ToLower(v.Type().Field(i).Name)
				return NewValidationError(fieldName + " cannot be empty")
			}
		}
	}

	return nil
}

// NewValidationError creates a new validation error
func NewValidationError(message string) error {
	return errors.New(message)
}
