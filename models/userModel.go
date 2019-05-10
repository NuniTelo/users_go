package models

import (
	"gopkg.in/mgo.v2/bson"
)

/**
This is the model for the user
 */



/**
This is the stuct for a user
 */
type User struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Name string `json:"name" bson:"name"`
	Lastname string `json:"lastname" bson:"lastname"`
	Age int `json:"age" bson:"age"`
}

/**
This is an array of users
 */
type Users[]User

/**
This is the save function fot the user
 */
func (user *User) SaveUser() error {
	err := checkMongoDBNotNull() //Let's check if the database is null or not
	err = mongoDB.C(user_collection).Insert(user) //We set a user to
	return err
}

/**
This is for getting a user based on id
 */
func (user *User) GetUserBasedOnID(id string) error {
	err := checkMongoDBNotNull();
	err = mongoDB.C(user_collection).FindId(bson.ObjectIdHex(id)).One(user)
	return err
}


/**
This is for getting a user based on it's name
 */
func (user *User) GetUserByName(name string) error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(user_collection).Find(bson.M{"name":name}).One(user)
	return err
}


/**
This is the method for deleting a user form the database
 */
func (user *User) DeleteUserFormDatabase(id string) error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(user_collection).RemoveId(id)
	return err
}

/**
This will get all users
 */
func (users *Users) GetAllUsers() error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(user_collection).Find(bson.M{}).All(users)
	return err
}

