package credentials

import (
	"errors"
	"gorm.io/gorm"
	"todolist/database/models"
	"todolist/domain"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(credentials *domain.Credentials) error {
	return r.db.Create(&models.Credentials{
		UserId:   credentials.UserId,
		Username: credentials.Username,
		Password: credentials.Password,
	}).Error
}

func (r *DBRepository) Get(username, password string) (string, error) {
	var creds models.Credentials
	r.db.Where(&models.Credentials{Username: username}).First(&creds)
	if creds.Password == password {
		return creds.UserId, nil
	}

	return "", errors.New("неверный логин или пароль")
}
