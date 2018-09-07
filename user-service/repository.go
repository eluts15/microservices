// user-service/repository.go
package main

import (
	pb "github.com/eluts15/microservices/shipping/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) (*pb.User, error)
	GetEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var user []*pb.User
	if err := repo.db.Find(&Users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&User).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
}
