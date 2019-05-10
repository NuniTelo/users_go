package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-contacts/models"
	"log"
)

func GetUsers (c *gin.Context) {
	users := models.Users{} //this will be a list of users

	userReference := &users  //get the reference here  STRUCT.METHOD() call

	err := userReference.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.Error(errors.New("500" + err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"data":users,
	})
}

func GetUserById (c *gin.Context){

	userID := c.Param("id")	//let's get the user id on pramas

	user := &models.User{}   //this is just the address on memeory of the user

	err := user.GetUserBasedOnID(userID)
	if err != nil {
		log.Println(err)
		c.Error(errors.New("Error" + err.Error()))
	}

	c.JSON(200,gin.H{
		"user":user,
	})
}

func GetUserByName(c *gin.Context) {
	userName := c.Param("name")
	userFromDB := &models.User{}

	err := userFromDB.GetUserByName(userName)
	if err != nil {
		log.Println(err)
		c.Error(errors.New("Error on :"  + err.Error()))
	}

	c.JSON(200,gin.H{
		"data":userFromDB,
	})
}

func DeleteUserFormDatabase(c *gin.Context) {
	userID := c.Param("userID")

	userResponse := &models.User{}

	err := userResponse.DeleteUserFormDatabase(userID)
	if err != nil {
		log.Println(err)
		c.Error(errors.New("Error on : " + err.Error()))
	}
	c.JSON(200,gin.H{
		"data":userResponse,
	})
}

func SaveUser(c *gin.Context) {
	var bodyRequest struct{
		Username string `json:"username" form:"username"`
		Name string `json:"name" form:"name"`
		Lastname string `json:"lastname" form:"lastname"`
		Age int `json:"age" form:"age"`
	}

	err := c.ShouldBindBodyWith(&bodyRequest,binding.JSON)
	if err != nil {
		log.Println(err)
		c.Error(errors.New("400| " + err.Error() ))
		return
	}

	userModel := models.User{
		Name:bodyRequest.Name,
		Username:bodyRequest.Username,
		Lastname:bodyRequest.Lastname,
		Age:bodyRequest.Age,
	}

	err = userModel.SaveUser()
	if err != nil {
		log.Println(err)
		c.Error(errors.New("Error : " + err.Error() ))
		return
	}

	c.JSON(200,gin.H{
		"success":"ok",
	})

}