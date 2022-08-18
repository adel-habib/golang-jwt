package helpers

import (
	"time"

	"github.com/adel-habib/golang-jwt/database"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var secret = "MYSECRET"

func GenerateAllTokens(email string, first_name string, last_name string, user_type string, user_id string) (token string, ref_token string, err error) {

	claims := &SignedDetails{
		Email:      email,
		First_name: first_name,
		Last_name:  last_name,
		Uid:        user_id,
		User_type:  user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour).Unix(),
		},
	}

	referesh_claims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24).Unix(),
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(secret))
	ref_token, err = jwt.NewWithClaims(jwt.SigningMethodES256, referesh_claims).SignedString([]byte(secret))

	return token, ref_token, err
}
