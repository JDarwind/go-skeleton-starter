package httpkit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"github.com/go-playground/validator/v10"
)


type ValidationMessages interface {
	Messages() map[string]map[string]string
}

var validate = validator.New()

func ValidateRequest[T any](r *http.Request, req *T) (any, map[string][]string) {
	if r.Method != http.MethodGet {
		_ = r.ParseForm() 
		_ = json.NewDecoder(r.Body).Decode(req)
	}

	bindQueryParams(r, req)

	if err := validate.Struct(req); err != nil {
		return nil, formatErrors(err, req)
	}
	return any(*req), nil
}

func formatErrors(err error, req any) map[string][]string {
	out := map[string][]string{}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return out
	}

	t := reflect.TypeOf(req).Elem()

	var custom map[string]map[string]string
	if m, ok := req.(ValidationMessages); ok {
		custom = m.Messages()
	}

	for _, e := range validationErrors {
		field, _ := t.FieldByName(e.StructField())

		jsonName := field.Tag.Get("json")
		if jsonName == "" || jsonName == "-" {
			jsonName = strings.ToLower(e.StructField())
		}

		rule := e.Tag()

		if custom != nil {
			if msg, ok := custom[jsonName][rule]; ok {
				out[jsonName] = append(out[jsonName], msg)
				continue
			}
		}

		out[jsonName] = append(out[jsonName], defaultMessage(jsonName, e))
	}

	return out
}

func defaultMessage(field string, e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required", field)
	case "email":
		return fmt.Sprintf("The %s field must be a valid email address", field)
	case "min":
		return fmt.Sprintf(
			"The %s field must be at least %s characters long",
			field, e.Param(),
		)
	case "gte":
		return fmt.Sprintf(
			"The %s field must be greater than or equal to %s",
			field, e.Param(),
		)
	default:
		return fmt.Sprintf(
			"The %s field is invalid (%s)",
			field, e.Tag(),
		)
	}
}

func bindQueryParams(r *http.Request, dst any) {
	values := r.URL.Query()

	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}

		val := values.Get(tag)
		if val == "" {
			continue
		}

		f := v.Field(i)
		if !f.CanSet() {
			continue
		}

		switch f.Kind() {
		case reflect.String:
			f.SetString(val)
		case reflect.Int:
			if n, err := strconv.Atoi(val); err == nil {
				f.SetInt(int64(n))
			}
		case reflect.Bool:
			if b, err := strconv.ParseBool(val); err == nil {
				f.SetBool(b)
			}
		}
	}
}
