package constants

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError mengekstrak error validasi menjadi map yang mudah dibaca
func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	switch e := err.(type) {

	// Validator error (misalnya `binding:"required"`)
	case validator.ValidationErrors:
		for _, fe := range e {
			field := jsonFieldNameFromStruct(fe.StructField(), fe.StructNamespace())
			errors[field] = generateValidationMessage(fe)
		}

	// Error unmarshalling tipe data (misalnya string → int)
	case *json.UnmarshalTypeError:
		field := jsonFieldNameFromPath(e.Field)
		errors[field] = fmt.Sprintf("invalid type: expected %s", e.Type.String())

	// Error syntax JSON (misalnya kurung tutup kurang)
	case *json.SyntaxError:
		errors["body"] = "invalid JSON syntax"

	// General fallback parsing error
	default:
		msg := err.Error()

		// Tangani error `cannot unmarshal ... into Go struct field ...`
		if strings.Contains(msg, "cannot unmarshal") && strings.Contains(msg, "into Go struct field") {
			field := extractFieldFromUnmarshalError(msg)
			errors[field] = "invalid type or format"
		} else {
			errors["general"] = msg
		}
	}

	return errors
}

// generateValidationMessage mengubah error tag menjadi pesan human readable
func generateValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}

// jsonFieldNameFromPath mengekstrak nama field terakhir dari path JSON
// Contoh: "TestRequest.postView" → "postView"
func jsonFieldNameFromPath(path string) string {
	parts := strings.Split(path, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return path
}

// jsonFieldNameFromStruct mengambil nama field JSON dari tag `json` (jika ada)
func jsonFieldNameFromStruct(structField, namespace string) string {
	// Ambil nama field dari bagian terakhir namespace
	fieldParts := strings.Split(namespace, ".")
	fieldName := fieldParts[len(fieldParts)-1]

	// Jika kamu ingin menggunakan tag `json` di sini (opsional, perlu reflect struct-nya)
	// Maka kamu bisa extend fungsi ini, contohnya:
	// getJSONTagName(structType, structField)

	return fieldName
}

// extractFieldFromUnmarshalError mengekstrak nama field dari pesan error JSON parsing
func extractFieldFromUnmarshalError(msg string) string {
	// Contoh error:
	// "json: cannot unmarshal string into Go struct field TestRequest.postView of type int"
	parts := strings.Split(msg, "into Go struct field")
	if len(parts) < 2 {
		return "unknown"
	}
	fieldPart := strings.TrimSpace(parts[1])
	fieldInfo := strings.Split(fieldPart, " ")[0]
	return jsonFieldNameFromPath(fieldInfo)
}
