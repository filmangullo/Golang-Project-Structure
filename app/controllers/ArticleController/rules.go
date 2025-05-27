package ArticleController

import (
	"strings"
)

/*
|--------------------------------------------------------------------------
| Rules
|--------------------------------------------------------------------------
|
| ValidateRuleInput performs custom validation on input data.
| This function is intended for applying advanced business rules
| that cannot be handled by basic struct tags, such as date checks,
| cross-field logic, or conditional requirements.
|
*/

func ValidatePostRequest(in PostRequest) map[string]string {
	errors := make(map[string]string)

	if len(strings.TrimSpace(in.Title)) < 2 {
		errors["Title"] = "Title must be at least 2 characters"
	}

	if !strings.Contains(in.Email, "@") {
		errors["Email"] = "Email must be valid"
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}

func ValidateGetRequest(in GetRequest) map[string]string {
	errors := make(map[string]string)

	if len(strings.TrimSpace(in.Slug)) < 2 {
		errors["Slug"] = "Slug must be at least 2 characters"
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
