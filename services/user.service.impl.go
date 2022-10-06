package services

import (
	"context"
	"errors"

	"github.com/MountainGator/playlist_CRUD/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	store          *sessions.CookieStore
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, store *sessions.CookieStore, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		store:          store,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User, c *gin.Context) error {
	var temp *models.User
	var err error
	query := bson.D{bson.E{Key: "username", Value: user.Username}}

	err = u.usercollection.FindOne(u.ctx, query).Decode(&temp)

	if err != nil {
		_, er := u.usercollection.InsertOne(u.ctx, user)
		session, ses_err := u.store.Get(c.Request, "session")
		if ses_err != nil {
			return ses_err
		}
		session.Values["user"] = temp.Username
		session.Save(c.Request, c.Writer)
		if er != nil {
			return er
		}

	} else if temp.Username == user.Username {
		return errors.New("user already exists")
	}
	return nil
}

func (u *UserServiceImpl) UserLogin(name *string, pwd string, c *gin.Context) error {
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: name}}
	if err := u.usercollection.FindOne(u.ctx, query).Decode(&user); err != nil {
		return err
	}

	pwd_err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd))

	if pwd_err != nil {
		return pwd_err
	}

	session, ses_err := u.store.Get(c.Request, "session")
	if ses_err != nil {
		return ses_err
	}

	session.Values["user"] = name
	session.Save(c.Request, c.Writer)
	return nil
}

func (u *UserServiceImpl) GetUserDetails(user_name *string) (*models.User, error) {
	var user *models.User
	filter := bson.D{primitive.E{Key: "username", Value: user_name}}
	if err := u.usercollection.FindOne(u.ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) Logout(c *gin.Context) error {
	session, err := u.store.Get(c.Request, "session")
	if err != nil {
		return err
	}
	session.Values["user"] = nil
	session.Save(c.Request, c.Writer)
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{primitive.E{Key: "_id", Value: user.Id}}
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "username", Value: user.Username},
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
