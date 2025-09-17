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

func ValidateCreateRequest(in CreateRequest) map[string]string {
	errors := make(map[string]string)

	if len(strings.TrimSpace(in.Title)) < 2 {
		errors["Title"] = "title must be at least 2 characters"
	}

	if in.IsPublished == nil {
		errors["IsPublished"] = "is published is required"
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}

func ValidateListRequest(in ListRequest) map[string]string {
	errors := make(map[string]string)

	if in.Page != nil {
		if *in.Page <= 0 {
			errors["Page"] = "page must be greater than 0"
		}
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}

func ValidateGetRequest(in GetRequest) map[string]string {
	errors := make(map[string]string)

	if len(errors) == 0 {
		return nil
	}
	return errors
}

func ValidateUpdateRequest(in UpdateRequest) map[string]string {
	errors := make(map[string]string)

	if len(strings.TrimSpace(in.Title)) < 2 {
		errors["Title"] = "title must be at least 2 characters"
	}

	titleExists, err := tblArticle.ExistsByWhere("title = ? AND id != ?", in.Title, in.ID)
	if err != nil {
		errors["Title"] = err.Error()
	}
	if titleExists {
		errors["Title"] = "the title is available, please do not duplicate the title"
	}

	if len(errors) == 0 {
		return nil
	}
	return errors
}
