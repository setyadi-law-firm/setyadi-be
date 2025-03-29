package auth

import (
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(user *User) error
	FindByEmail(email string) (*User, error)
}

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) AuthRepository {
	return &GormAuthRepository{db}
}

func (r *GormAuthRepository) Save(user *User) error {
	return r.db.Create(user).Error
}

func (r *GormAuthRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
