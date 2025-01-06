package controllers

import (
	"context"
	"go/token"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"honnef.co/go/tools/analysis/facts/generated"
)

func HashPassword(password string) string {

}

func VerifyPassword(usserPassword string, givenPassword string) (bool, string) {}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateionErr := Validate.Struct(user)
		if validateionErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err:= UserCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0{
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}
		
		count, err = UserCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
		defer cancel()

		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0{
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone no. already exists"})
			return
		}

		password, err := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.User_ID.Hex()

		token, refreshToken, _ :=  generate.ToeknGenerator(*user.Email, *user.First_Name, *user.Last_Name)
		user.Token = &token
		user.Refresh_Token = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"the user did not get created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated,"Successfully signed in")
		
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"login or password incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founder.Password)
		defer cancel()

		if !PasswordIsValid{
			c.JSON{http.StatusInternalServerError, gin.H{"error":msg}}
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email,*founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, *founduser.User_ID)
		c.JSON(http.StatusFound, founduser)
	}
}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}
