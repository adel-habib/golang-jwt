package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/adel-habib/golang-jwt/database"
	"github.com/adel-habib/golang-jwt/helpers"
	"github.com/adel-habib/golang-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Hashpassword() {}

func VerifyPassword(givenPassword string, actualPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(givenPassword))
	if err != nil {
		return false, "incorrect password"
	}
	return true, "password correct!"
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if count != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered!"})
			return
		}
		user.Created_at = time.Now()
		user.Updated_at = time.Now()
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token, refereshtoken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, user.User_id)
		user.Token = &token
		user.Referesh_token = &refereshtoken

		resIN, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User was not creatred!"})
			return
		}
		c.JSON(http.StatusOK, resIN)
	}
}

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User
		var foundUser models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)

	}
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Param("user_id")
		err := helpers.MatchUserTypeToUid(c, userid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err = userCollection.FindOne(ctx, bson.M{"user_id": userid}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)

	}
}

func getUsers() {}
