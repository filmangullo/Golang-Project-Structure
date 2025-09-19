package tableArticle

import (
	"your_project_name/models"
	"your_project_name/utils/PaginateFunctions"

	"gorm.io/gorm"
)

type DatabaseTableArticle interface {
	// CREATE
	Create(article models.Article) (models.Article, error)

	// READ
	ReadAll() ([]models.Article, error)
	ReadByWhere(condition string, args ...interface{}) ([]models.Article, error)
	ReadFirstByWhere(condition string, args ...interface{}) (models.Article, error)
	ReadWhereByPaginate(opts PaginateFunctions.QueryOptions) (PaginateFunctions.PageResult[models.Article], error)

	// READ additionally
	ReadWhereIn(column string, values interface{}) ([]models.Article, error)
	ReadWhereNotIn(column string, values interface{}) ([]models.Article, error)
	ReadWhereBetween(column string, from, to interface{}) ([]models.Article, error)
	ReadWhereNull(column string) ([]models.Article, error)
	ReadWhereNotNull(column string) ([]models.Article, error)
	ReadSearchLike(keyword string, cols ...string) ([]models.Article, error)
	CountByWhere(condition string, args ...interface{}) (int64, error)
	ExistsByWhere(condition string, args ...interface{}) (bool, error)

	// READ optional
	ReadSelectOrder(selectExpr, orderBy string, condition string, args ...interface{}) ([]models.Article, error)
	ReadWith(preloads []string, condition string, args ...interface{}) ([]models.Article, error)

	// UPDATE
	UpdateByID(id interface{}, fields map[string]interface{}) (models.Article, error)
	UpdateBySlug(slug string, fields map[string]interface{}) (models.Article, error)

	// DELETE
	DeleteBy(fields map[string]interface{}) (bool, error)
	DeleteWhere(condition string, args ...interface{}) (bool, error)
}

type articleRepository struct {
	db *gorm.DB
}

func CallArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

// ====== CREATE ======
func (r *articleRepository) Create(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

// ====== READ ======
func (r *articleRepository) ReadAll() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) ReadByWhere(condition string, args ...interface{}) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(condition, args...).Find(&articles).Error

	return articles, err
}

func (r *articleRepository) ReadFirstByWhere(condition string, args ...interface{}) (models.Article, error) {
	var article models.Article
	err := r.db.Where(condition, args...).First(&article).Error

	return article, err
}

func (r *articleRepository) ReadWhereByPaginate(opts PaginateFunctions.QueryOptions) (PaginateFunctions.PageResult[models.Article], error) {
	return PaginateFunctions.GenericPaginate[models.Article](r.db, opts)
}

// ====== READ additionally ======
func (r *articleRepository) ReadWhereIn(column string, values interface{}) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(column+" IN ?", values).Find(&articles).Error

	return articles, err
}

func (r *articleRepository) ReadWhereNotIn(column string, values interface{}) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(column+" NOT IN ?", values).Find(&articles).Error

	return articles, err
}

func (r *articleRepository) ReadWhereBetween(column string, from, to interface{}) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(column+" BETWEEN ? AND ?", from, to).Find(&articles).Error

	return articles, err
}

func (r *articleRepository) ReadWhereNull(column string) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(column + " IS NULL").Find(&articles).Error

	return articles, err
}

func (r *articleRepository) ReadWhereNotNull(column string) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Where(column + " IS NOT NULL").Find(&articles).Error

	return articles, err
}

// LIKE in multiple columns: (title LIKE ? OR content LIKE ?)
func (r *articleRepository) ReadSearchLike(keyword string, cols ...string) ([]models.Article, error) {
	var articles []models.Article
	tx := r.db.Model(&models.Article{})
	if keyword != "" && len(cols) > 0 {
		tx = tx.Scopes(PaginateFunctions.ScopeSearchLike(keyword, cols...))
	}
	err := tx.Find(&articles).Error

	return articles, err
}

func (r *articleRepository) CountByWhere(condition string, args ...interface{}) (int64, error) {
	var total int64
	err := r.db.Model(&models.Article{}).Where(condition, args...).Count(&total).Error

	return total, err
}

func (r *articleRepository) ExistsByWhere(condition string, args ...interface{}) (bool, error) {
	var total int64
	err := r.db.Model(&models.Article{}).Where(condition, args...).Limit(1).Count(&total).Error

	return total > 0, err
}

// ====== READ optional ======
// SELECT + ORDER (+ optional WHERE)
func (r *articleRepository) ReadSelectOrder(selectExpr, orderBy string, condition string, args ...interface{}) ([]models.Article, error) {
	var articles []models.Article
	tx := r.db.Model(&models.Article{})
	if selectExpr != "" {
		tx = tx.Select(selectExpr)
	}
	if condition != "" {
		tx = tx.Where(condition, args...)
	}
	if orderBy != "" {
		tx = tx.Order(orderBy)
	}
	err := tx.Find(&articles).Error

	return articles, err
}

// WITH (preload relation)
func (r *articleRepository) ReadWith(preloads []string, condition string, args ...interface{}) ([]models.Article, error) {
	var articles []models.Article
	tx := r.db.Model(&models.Article{})
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	if condition != "" {
		tx = tx.Where(condition, args...)
	}
	err := tx.Find(&articles).Error

	return articles, err
}

// ====== UPDATE ======
func (r *articleRepository) UpdateByID(id interface{}, fields map[string]interface{}) (models.Article, error) {
	var updated models.Article

	if id == nil || len(fields) == 0 {
		return updated, nil
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("id = ?", id).
			Updates(fields).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", id).First(&updated).Error
	})

	return updated, err
}

func (r *articleRepository) UpdateBySlug(slug string, fields map[string]interface{}) (models.Article, error) {
	var updated models.Article

	if slug == "" || len(fields) == 0 {
		return updated, nil
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("slug = ?", slug).
			Updates(fields).Error; err != nil {
			return err
		}
		return tx.Where("slug = ?", slug).First(&updated).Error
	})

	return updated, err
}

// ====== DELETE ======
func (r *articleRepository) DeleteBy(fields map[string]interface{}) (bool, error) {
	if len(fields) == 0 {
		return false, nil
	}
	tx := r.db.Where(fields).Delete(&models.Article{})

	return tx.RowsAffected > 0, tx.Error
}

func (r *articleRepository) DeleteWhere(condition string, args ...interface{}) (bool, error) {
	tx := r.db.Where(condition, args...).Delete(&models.Article{})

	return tx.RowsAffected > 0, tx.Error
}
