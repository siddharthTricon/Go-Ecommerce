package controllers

import (
	"github.com/gin-gonic/gin"
)

func HashPassword(password string) string {

}

func VerifyPassword(usserPassword string, givenPassword string) (bool, string) {}

func SignUp() gin.HandlerFunc {}

func Login() gin.HandlerFunc {}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}
