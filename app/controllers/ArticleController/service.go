package ArticleController

import (
	"fmt"
	"your_project_name/utils/DatetimeFunctions"
	"your_project_name/utils/StringsFunctions"
)

func WelcomeService(input PostRequest) (*PostRequest, error) {
	var response PostRequest

	fmt.Println(DatetimeFunctions.FormatDateStyle("2025-07-19 00:00:00", "Y-m-d"))
	fmt.Println(StringsFunctions.ToInt8("42"))
	fmt.Println(StringsFunctions.ToInt16("42"))
	fmt.Println(StringsFunctions.ToInt32("42"))
	fmt.Println(StringsFunctions.ToInt64("42"))
	fmt.Println(StringsFunctions.ToInt("42"))
	return &response, nil
}

/*
# -----------------------------------------------------------------------------
# GET
# -----------------------------------------------------------------------------
*/
func GetArticleService(input GetRequest) (*GetResponse, error) {
	var response GetResponse

	return &response, nil
}
