package repositories

import (
	"github.com/chitano/chatapp/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUserByID(id uint64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// Get User By Email
func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get User By ID
func (u *userRepository) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Get all Users
func (u *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Create a new user
func (u *userRepository) CreateUser(user *model.User) error {
	return u.db.Create(user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
