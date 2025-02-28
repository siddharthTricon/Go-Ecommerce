package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddharthTricon/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc{}

func EditHomeAdress() gin.HandlerFunc{}

func EditWorkAddress() gin.HandlerFunc{}

func DeleteAddress() gin.HandlerFunc{
	return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		usert_id, err := primitive.ObejectIDFromHex(user_id)
		if err != nil{
			c.IndentedJSON(500, "Invalid Server Error")
		}

		var ctx,  cancel = context.Background(), WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "id", Value: usert_id}}
		update := bson.D{{Key:"$set", Value:bson.D{primitive.E{Key:"address", Value: addresses}}}}
		_, err := UserCollection.UpdateOne(ctx, filter, update) 
		if err != nil{
			c.IndentedJSON(404, "Wrong Command")
			return
		}

		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Deleted")
	}
}
