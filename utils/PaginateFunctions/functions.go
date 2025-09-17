package PaginateFunctions

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type QueryOptions struct {
	Where      string
	WhereArgs  []interface{}
	Search     string
	SearchCols []string

	Select   string
	Preloads []string
	Order    string

	Page    int
	PerPage int

	// Jumlah nomor di kiri-kanan halaman aktif (default: 2 jika 0)
	Window int
}

type PageLabel struct {
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

type PageResult[T any] struct {
	Results  []T         `json:"results"`
	Page     int         `json:"page"`
	PerPage  int         `json:"perPage"`
	Total    int64       `json:"total"`
	LastPage int         `json:"lastPage"`
	Labels   []PageLabel `json:"labels"`
	HasNext  bool        `json:"hasNext"`
	HasPrev  bool        `json:"hasPrev"`
}

// ScopeSearchLike builds a grouped LIKE (col1 LIKE ? OR col2 LIKE ? ...)
func ScopeSearchLike(keyword string, cols ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" || len(cols) == 0 {
			return db
		}
		pattern := "%" + keyword + "%"
		query := "("
		args := make([]interface{}, 0, len(cols))
		for i, c := range cols {
			if i > 0 {
				query += " OR "
			}
			query += c + " LIKE ?"
			args = append(args, pattern)
		}
		query += ")"
		return db.Where(query, args...)
	}
}

// buildLabelsWindowed membangun labels dengan window & ellipsis.
func buildLabelsWindowed(current, last, window int) []PageLabel {
	labels := make([]PageLabel, 0, 10)

	// Previous
	labels = append(labels, PageLabel{
		Label:  "« Previous",
		Active: current > 1,
	})

	if last <= 1 {
		labels = append(labels, PageLabel{Label: "1", Active: current == 1})
		labels = append(labels, PageLabel{Label: "Next »", Active: false})
		return labels
	}

	// Halaman pertama selalu ada
	labels = append(labels, PageLabel{Label: "1", Active: current == 1})

	// Hitung rentang window
	start := current - window
	if start < 2 {
		start = 2
	}
	end := current + window
	if end > last-1 {
		end = last - 1
	}

	// Ellipsis setelah 1 jika perlu
	if start > 2 {
		labels = append(labels, PageLabel{Label: "…", Active: false})
	}

	// Nomor dalam window
	for i := start; i <= end; i++ {
		labels = append(labels, PageLabel{
			Label:  fmt.Sprintf("%d", i),
			Active: i == current,
		})
	}

	// Ellipsis sebelum last jika perlu
	if end < last-1 {
		labels = append(labels, PageLabel{Label: "…", Active: false})
	}

	// Halaman terakhir (jika > 1)
	if last > 1 {
		labels = append(labels, PageLabel{
			Label:  fmt.Sprintf("%d", last),
			Active: current == last,
		})
	}

	// Next
	labels = append(labels, PageLabel{
		Label:  "Next »",
		Active: current < last,
	})

	return labels
}

// GenericPaginate works for any model type T
func GenericPaginate[T any](db *gorm.DB, opts QueryOptions) (PageResult[T], error) {
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PerPage <= 0 {
		opts.PerPage = 10
	}
	if opts.Window <= 0 {
		opts.Window = 2
	}

	tx := db.Model(new(T))

	if opts.Select != "" {
		tx = tx.Select(opts.Select)
	}
	if opts.Where != "" {
		tx = tx.Where(opts.Where, opts.WhereArgs...)
	}
	if opts.Search != "" && len(opts.SearchCols) > 0 {
		tx = tx.Scopes(ScopeSearchLike(opts.Search, opts.SearchCols...))
	}

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return PageResult[T]{}, err
	}

	if opts.Order != "" {
		tx = tx.Order(opts.Order)
	}
	for _, p := range opts.Preloads {
		tx = tx.Preload(p)
	}

	offset := (opts.Page - 1) * opts.PerPage
	var rows []T
	if err := tx.Offset(offset).Limit(opts.PerPage).Find(&rows).Error; err != nil {
		return PageResult[T]{}, err
	}

	last := int(math.Ceil(float64(total) / float64(opts.PerPage)))

	labels := buildLabelsWindowed(opts.Page, last, opts.Window)

	return PageResult[T]{
		Results:  rows,
		Page:     opts.Page,
		PerPage:  opts.PerPage,
		Total:    total,
		LastPage: last,
		Labels:   labels,
		HasNext:  opts.Page < last,
		HasPrev:  opts.Page > 1,
	}, nil
}
