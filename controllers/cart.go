package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-ecommerce/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct{
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application{
	return Application{
		prodCollection: prodCollection,
		userCollection: userCollection
	}
}

func (app *Application) AddToCart() gin.HAndlerFunc{
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.NEw("product is empty"))
			return
		} 
		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.IndentedJSON((200, "Successfully added to cart"))
	}	
}

func (app *Application) RemoveItem() gin.HandlerFunc{
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.NEw("product is empty"))
			return
		} 
		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.RemoveProductFromCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
			}
			c.IndentedJSON((200, "Successfully removed from cart"))
	}	
}

func GetItemFromCart() gin.HandlerFunc{}

func (app *Application) BuyFromCart() gin.HandlerFunc{
	return func(c *gin.Context){
		userQueryID := c.Query("id")

		if userQueryID == ""{
			log.Println("user id is empty")
			- = c.ABortWithError(http.StatusBadREquest, errors.NEw("UserID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON((200, "Successfully placed the items"))
	}
}

func (app *Application) InstantBuy() gin.HadlerFunc{
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == ""{
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.NEw("product is empty"))
			return
		} 
		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortWithError(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cnacel()

		err = database.InstantBuy(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, "Successfully bought the product")

	}
}

