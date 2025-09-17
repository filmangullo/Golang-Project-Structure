package tableUser

import (
	"your_project_name/models"

	"gorm.io/gorm"
)

type DatabaseTableUser interface {
	Create(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func CallUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
