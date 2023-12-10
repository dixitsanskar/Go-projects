package services

import (
	"context"

	"example.com/letter-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx 			context.Context
}
//service layer will interact with database layer so it has databse object 

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService{
	return &UserServiceImpl {
		usercollection: usercollection,
		ctx: ctx,
	}
}//acts like an constructor

//implementing userfunction using function below
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error){
	return nil, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error){
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error{
	return nil 
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}