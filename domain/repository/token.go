package repository

import (
	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/INVITATION-RPC-AUTH/pkg/database"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Get() (companies []models.Token, err error)
	Insert(member models.Token) (*models.Token, error)
	Update(member models.Token, id int) (*models.Token, error)
	Delete(id int) (bool, error)
	GetById(id int) (*models.Token, error)
	GetByUserId(userId int) (resp *models.Token, err error)
}

type tokenRepository struct {
	H *gorm.DB
}

func NewTokenRepository() *tokenRepository {
	return &tokenRepository{
		H: database.DB,
	}
}

func (r *tokenRepository) Get() (tokens []models.Token, err error) {

	return tokens, nil
}

func (r *tokenRepository) Insert(token models.Token) (*models.Token, error) {
	err := r.H.Save(&token).Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *tokenRepository) Update(token models.Token, id int) (*models.Token, error) {

	return &token, nil
}

func (r *tokenRepository) Delete(id int) (bool, error) {
	err := r.H.Delete(&models.Token{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *tokenRepository) GetById(id int) (token *models.Token, err error) {

	return token, nil
}

func (r *tokenRepository) GetByUserId(userId int) (token *models.Token, err error) {

	if result := r.H.Where(&models.Token{UserID: userId}).First(&token); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return token, nil
}
