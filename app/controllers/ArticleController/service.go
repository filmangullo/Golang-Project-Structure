package ArticleController

func WelcomeService(input PostRequest) (*PostRequest, error) {
	var response PostRequest

	// fmt.Println(StringsFunctions.Lower(input.Title))
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
