package repository

import (
	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/INVITATION-RPC-AUTH/pkg/database"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Get() (companies []models.User, err error)
	Insert(member models.User) (*models.User, error)
	Update(member models.User, id int32) (*models.User, error)
	Delete(id uint32) (bool, error)
	GetById(id uint32) (*models.User, error)
	FindUserByEmail(email string) (resp *models.User, err error)
}

type authRepository struct {
	H *gorm.DB
}

func NewAuthRepository() *authRepository {
	return &authRepository{
		H: database.DB,
	}
}

func (r *authRepository) Get() (users []models.User, err error) {

	return users, nil
}

func (r *authRepository) Insert(user models.User) (*models.User, error) {
	err := r.H.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) Update(user models.User, id int32) (*models.User, error) {

	return &user, nil
}

func (r *authRepository) Delete(id uint32) (bool, error) {
	err := r.H.Delete(&models.User{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *authRepository) GetById(id uint32) (user *models.User, err error) {

	return user, nil
}

func (r *authRepository) FindUserByEmail(email string) (user *models.User, err error) {

	if result := r.H.Where(&models.User{Email: email}).First(&user); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return user, nil
}
