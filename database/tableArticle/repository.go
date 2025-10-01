package tableArticle

import (
	"your_project_name/models"
	"your_project_name/utils/PaginateFunctions"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DatabaseTableArticle interface {
	// CREATE
	Create(row models.Article) (models.Article, error)

	// READ
	ReadAll() ([]models.Article, error)
	ReadByWhere(condition string, args ...interface{}) ([]models.Article, error)
	ReadFirstByWhere(condition string, args ...interface{}) (models.Article, error)
	ReadWhereByPaginate(opts PaginateFunctions.QueryOptions) (PaginateFunctions.PageResult[models.Article], error)

	// READ additionally
	ReadWhereIn(column string, values interface{}) ([]models.Article, error)
	ReadWhereNotIn(column string, values interface{}) ([]models.Article, error)
	ReadWhereBetween(column string, from, to interface{}) ([]models.Article, error)
	ReadWhereNull(column ...string) ([]models.Article, error)
	ReadWhereNotNull(column ...string) ([]models.Article, error)
	ReadByWhereAndNull(condition string, args []interface{}, nullFields ...string) ([]models.Article, error)
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
func (r *articleRepository) Create(row models.Article) (models.Article, error) {
	err := r.db.Create(&row).Error
	if err != nil {
		return row, err
	}

	return row, nil
}

// ====== READ ======
func (r *articleRepository) ReadAll() ([]models.Article, error) {
	var rows []models.Article
	err := r.db.Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *articleRepository) ReadByWhere(condition string, args ...interface{}) ([]models.Article, error) {
	var rows []models.Article
	err := r.db.Where(condition, args...).Find(&rows).Error

	return rows, err
}

func (r *articleRepository) ReadFirstByWhere(condition string, args ...interface{}) (models.Article, error) {
	var row models.Article
	err := r.db.Where(condition, args...).First(&row).Error

	return row, err
}

func (r *articleRepository) ReadWhereByPaginate(opts PaginateFunctions.QueryOptions) (PaginateFunctions.PageResult[models.Article], error) {
	return PaginateFunctions.GenericPaginate[models.Article](r.db, opts)
}

// ====== READ additionally ======
func (r *articleRepository) ReadWhereIn(column string, values interface{}) ([]models.Article, error) {
	var rows []models.Article
	err := r.db.Where(column+" IN ?", values).Find(&rows).Error

	return rows, err
}

func (r *articleRepository) ReadWhereNotIn(column string, values interface{}) ([]models.Article, error) {
	var rows []models.Article
	err := r.db.Where(column+" NOT IN ?", values).Find(&rows).Error

	return rows, err
}

func (r *articleRepository) ReadWhereBetween(column string, from, to interface{}) ([]models.Article, error) {
	var rows []models.Article
	err := r.db.Where(column+" BETWEEN ? AND ?", from, to).Find(&rows).Error

	return rows, err
}

// WHERE <field> IS NULL (bisa banyak field; semua digabung dengan AND)
func (r *articleRepository) ReadWhereNull(column ...string) ([]models.Article, error) {
	var rows []models.Article
	tx := r.db
	for _, f := range column {
		// map[string]interface{}{f: nil} => menghasilkan "<f> IS NULL"
		tx = tx.Where(map[string]interface{}{f: nil})
	}
	err := tx.Find(&rows).Error
	return rows, err
}

// WHERE <field> IS NOT NULL (bisa banyak field; semua digabung dengan AND) – aman dari SQL injection
func (r *articleRepository) ReadWhereNotNull(column ...string) ([]models.Article, error) {
	var rows []models.Article
	tx := r.db
	for _, f := range column {
		// clause.Expr dengan clause.Column memastikan nama kolom ditangani sebagai identifier, bukan value
		tx = tx.Where(clause.Expr{SQL: "? IS NOT NULL", Vars: []interface{}{clause.Column{Name: f}}})
	}
	err := tx.Find(&rows).Error
	return rows, err
}

// Kombinasi: WHERE <condition> AND (<field> IS NULL ...)
func (r *articleRepository) ReadByWhereAndNull(condition string, args []interface{}, nullFields ...string) ([]models.Article, error) {
	var rows []models.Article
	tx := r.db

	if condition != "" {
		tx = tx.Where(condition, args...)
	}
	for _, f := range nullFields {
		tx = tx.Where(map[string]interface{}{f: nil})
	}

	err := tx.Find(&rows).Error
	return rows, err
}

// LIKE in multiple columns: (title LIKE ? OR content LIKE ?)
func (r *articleRepository) ReadSearchLike(keyword string, cols ...string) ([]models.Article, error) {
	var rows []models.Article
	tx := r.db.Model(&models.Article{})
	if keyword != "" && len(cols) > 0 {
		tx = tx.Scopes(PaginateFunctions.ScopeSearchLike(keyword, cols...))
	}
	err := tx.Find(&rows).Error

	return rows, err
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
	var rows []models.Article
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
	err := tx.Find(&rows).Error

	return rows, err
}

// WITH (preload relation)
func (r *articleRepository) ReadWith(preloads []string, condition string, args ...interface{}) ([]models.Article, error) {
	var rows []models.Article
	tx := r.db.Model(&models.Article{})
	for _, p := range preloads {
		tx = tx.Preload(p)
	}
	if condition != "" {
		tx = tx.Where(condition, args...)
	}
	err := tx.Find(&rows).Error

	return rows, err
}

// ====== UPDATE ======
func (r *articleRepository) UpdateByID(id interface{}, fields map[string]interface{}) (models.Article, error) {
	var row models.Article

	if id == nil || len(fields) == 0 {
		return row, nil
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("id = ?", id).
			Updates(fields).Error; err != nil {
			return err
		}
		return tx.Where("id = ?", id).First(&row).Error
	})

	return row, err
}

func (r *articleRepository) UpdateBySlug(slug string, fields map[string]interface{}) (models.Article, error) {
	var row models.Article

	if slug == "" || len(fields) == 0 {
		return row, nil
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Article{}).
			Where("slug = ?", slug).
			Updates(fields).Error; err != nil {
			return err
		}
		return tx.Where("slug = ?", slug).First(&row).Error
	})

	return row, err
}

// ====== DELETE ======
func (r *articleRepository) DeleteBy(cols map[string]interface{}) (bool, error) {
	if len(cols) == 0 {
		return false, nil
	}
	tx := r.db.Where(cols).Delete(&models.Article{})

	return tx.RowsAffected > 0, tx.Error
}

func (r *articleRepository) DeleteWhere(condition string, args ...interface{}) (bool, error) {
	tx := r.db.Where(condition, args...).Delete(&models.Article{})

	return tx.RowsAffected > 0, tx.Error
}
