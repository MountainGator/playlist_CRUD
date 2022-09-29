package services

import (
	"context"
	"errors"

	"github.com/MountainGator/playlist_CRUD/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {

	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}
func (u *UserServiceImpl) UserLogin(name *string, pwd string) error {
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	pwd_err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd))

	if pwd_err != nil {
		return pwd_err
	}
	return err
}
func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{primitive.E{Key: "name", Value: user.Name}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "name", Value: user.Name},
				primitive.E{Key: "age", Value: user.Age},
				primitive.E{Key: "username", Value: user.Username},
				primitive.E{Key: "email", Value: user.Email},
			},
		},
	}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("couldn't find user")
	}
	return nil
}
func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("error. could not delete user")
	}
	return nil
}
