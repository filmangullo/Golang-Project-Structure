package ArticleController

import (
	"errors"
	"log"
	"time"
	"your_project_name/app/resources/articleResource"
	"your_project_name/models"
	"your_project_name/utils/PaginateFunctions"
	"your_project_name/utils/StringsFunctions"
)

func CreateService(input CreateRequest) (*CreateResponse, error) {
	article := models.Article{
		Title:       input.Title,
		Slug:        StringsFunctions.Slug(input.Title, '-'),
		Content:     input.Content,
		Author:      &input.Author,
		Category:    &input.Category,
		Tags:        &input.Tags,
		IsPublished: input.IsPublished,
		PublishedAt: time.Now(),
	}

	postArticle, err := tblArticle.Create(article)
	if err != nil {
		return nil, errors.New("failed to create article: " + err.Error())
	}

	response := CreateResponse{
		ArticleArray: articleResource.Resource(postArticle),
	}

	return &response, nil
}

/*
# -----------------------------------------------------------------------------
# LIST
# -----------------------------------------------------------------------------
*/
func ListArticleService(input ListRequest) (*ListResponse, error) {
	var response ListResponse

	// ambil halaman 1, 10 per page, filter by status
	page, err := tblArticle.ReadWhereByPaginate(PaginateFunctions.QueryOptions{
		Where:     "is_published = ?",
		WhereArgs: []any{true},
		Page:      StringsFunctions.IntCoalescePositive(input.Page, 1),
		PerPage:   10,
		Order:     "created_at DESC",
		Window:    2, // display ±2 numbers from the active page
	})
	if err != nil {
		log.Println("paginate error:", err)
	}

	// insert pagination results into response
	response = ListResponse{
		Results:  articleResource.Resources(page.Results),
		Page:     page.Page,
		PerPage:  page.PerPage,
		Total:    page.Total,
		LastPage: page.LastPage,
		HasNext:  page.HasNext,
		HasPrev:  page.HasPrev,
		Labels:   page.Labels,
	}

	return &response, nil
}

/*
# -----------------------------------------------------------------------------
# GET
# -----------------------------------------------------------------------------
*/
func GetArticleService(input GetRequest) (*GetResponse, error) {
	article, err := tblArticle.ReadByWhere("slug = ?", input.Slug)
	if err != nil {
		return nil, errors.New("failed to get article: " + err.Error())
	}

	response := GetResponse{
		articleResource.Resources(article),
	}

	return &response, nil
}

/*
# -----------------------------------------------------------------------------
# Update
# -----------------------------------------------------------------------------
*/
func UpdateArticleService(input UpdateRequest) (*UpdateResponse, error) {
	article, err := tblArticle.UpdateByID(input.ID, map[string]any{
		"title":        input.Title,
		"slug":         StringsFunctions.Slug(input.Title, '-'),
		"content":      input.Content,
		"author":       input.Author,
		"category":     input.Category,
		"tags":         input.Tags,
		"is_published": input.IsPublished,
	})
	if err != nil {
		return nil, errors.New("failed to update article: " + err.Error())
	}

	response := UpdateResponse{
		articleResource.Resource(article),
	}

	return &response, nil
}

/*
# -----------------------------------------------------------------------------
# Delete
# -----------------------------------------------------------------------------
*/
func DeleteArticleService(input DeleteRequest) (*DeleteResponse, error) {
	article, err := tblArticle.DeleteBy(map[string]interface{}{"id": input.ID})
	if err != nil {
		return nil, errors.New("failed to delete article: " + err.Error())
	}

	var msg string
	if !article {
		msg = "data has not been successfully deleted"
	} else {
		msg = "data deleted successfully"
	}

	response := DeleteResponse{
		Messages: map[string]string{
			"general": msg,
		},
	}

	return &response, nil
}
