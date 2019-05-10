package controllers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-contacts/models"
	"gopkg.in/mgo.v2/bson"
)

// GetTopics is a function that can return all of tables
func GetTopics(c *gin.Context) {
	topics := models.Topics{} //lets create a getter for the topics

	pointerOfTopics := &topics

	err := pointerOfTopics.GetAll()  //mqs methods aksesohen keshtu atehere : go

	if err != nil {
		log.Println(err)
		c.Error(errors.New("500|" + err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"data": topics,
	})
}


//GetTopicByID is a function that can return Topic by topic's id
func GetTopicByID(c *gin.Context) {
	topicID := c.Param("_id")  //params on the url
	topic := &models.Topic{} //create the referece so we can catch the motherfucking method

	err := topic.GetOneByID(topicID)  //go access the model here
	if err != nil {
		log.Println(err)
		c.Error(errors.New("500|" + err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"data": topic,
	})
}

// CreateTopic save topic to db
func CreateTopic(c *gin.Context) {
	var bodyRequest struct {
		Name string `form:"name" json:"name" binding:"required"`
	}
	err := c.ShouldBindWith(&bodyRequest, binding.JSON)
	if err != nil {
		log.Println(err)
		c.Error(errors.New("400|" + err.Error()))
		return
	}
	newTopic := &models.Topic{
		Name: bodyRequest.Name,
	}
	err = newTopic.Save()
	if err != nil {
		log.Println(err)
		c.Error(errors.New("500|" + err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"body": "ok",
	})
}

// DeleteTopic is a function that will recieve and delete
func DeleteTopic(c *gin.Context) {
	topic := &models.Topic{
		ID: bson.ObjectIdHex(c.Param("_id")),
	}
	err := topic.Delete()
	if err != nil {
		log.Println(err)
		c.Error(errors.New("500|" + err.Error()))
		return
	}
	c.JSON(200, gin.H{
		"body": "ok",
	})
}
