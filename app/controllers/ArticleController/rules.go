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
	rule_err := make(map[string]string)

	if len(strings.TrimSpace(in.Title)) < 2 {
		rule_err["Title"] = "title must be at least 2 characters"
	}

	if in.IsPublished == nil {
		rule_err["IsPublished"] = "is published is required"
	}

	if len(rule_err) == 0 {
		return nil
	}
	return rule_err
}

func ValidateListRequest(in ListRequest) map[string]string {
	rule_err := make(map[string]string)

	if in.Page != nil {
		if *in.Page <= 0 {
			rule_err["Page"] = "page must be greater than 0"
		}
	}

	if len(rule_err) == 0 {
		return nil
	}
	return rule_err
}

func ValidateGetRequest(in GetRequest) map[string]string {
	rule_err := make(map[string]string)

	if len(rule_err) == 0 {
		return nil
	}
	return rule_err
}

func ValidateUpdateRequest(in UpdateRequest) map[string]string {
	rule_err := make(map[string]string)

	if len(strings.TrimSpace(in.Title)) < 2 {
		rule_err["Title"] = "title must be at least 2 characters"
	}

	titleExists, err := tblArticle.ExistsByWhere("title = ? AND id != ?", in.Title, in.ID)
	if err != nil {
		rule_err["Title"] = err.Error()
	}
	if titleExists {
		rule_err["Title"] = "the title is available, please do not duplicate the title"
	}

	if len(rule_err) == 0 {
		return nil
	}
	return rule_err
}

func ValidateDeleteRequest(in DeleteRequest) map[string]string {
	rule_err := make(map[string]string)

	idExists, err := tblArticle.ExistsByWhere("id = ?", in.ID)
	if err != nil {
		rule_err["ID"] = err.Error()
	}
	if !idExists {
		rule_err["ID"] = "the ID no longer exists, please enter the correct ID."
	}

	if len(rule_err) == 0 {
		return nil
	}
	return rule_err
}
