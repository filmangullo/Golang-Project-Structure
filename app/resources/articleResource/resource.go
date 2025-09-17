package articleResource

import (
	"time"
	"your_project_name/models"
)

type ArticleArray struct {
	ID          uint64    `json:"key"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	Author      *string   `json:"author"`
	Category    *string   `json:"category"`
	Tags        *string   `json:"tags"`
	IsPublished *bool     `json:"isPublished"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Resource(req models.Article) ArticleArray {
	res := ArticleArray{}

	res.ID = req.ID
	res.Title = req.Title
	res.Slug = req.Slug
	res.Content = req.Content
	res.Author = req.Author
	res.Category = req.Category
	res.Tags = req.Tags
	res.IsPublished = req.IsPublished
	res.PublishedAt = req.PublishedAt
	res.CreatedAt = req.CreatedAt
	res.UpdatedAt = req.UpdatedAt

	return res
}

func Resources(req []models.Article) []ArticleArray {
	var res []ArticleArray

	for _, rq := range req {
		res = append(res, Resource(rq))
	}

	return res
}
