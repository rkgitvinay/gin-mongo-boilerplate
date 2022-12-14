package services

import (
	"fmt"
	"gin-mongo-boilerplate/src/user/dto"
	"gin-mongo-boilerplate/src/user/models"
	"log"
	"strconv"
	"time"

	pagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AllTodo(requestFilter map[string]interface{}) ([]models.Todo, pagination.PaginationData) {
	var users []models.Todo

	filter := bson.M{}

	if requestFilter["status"] != "" {
		filter["status"] = requestFilter["status"]
	}

	page, _ := strconv.ParseInt(requestFilter["page"].(string), 10, 64)
	limit, _ := strconv.ParseInt(requestFilter["limit"].(string), 10, 64)

	paginatedData, err := pagination.New(models.TodoCollection.Collection).
		Page(page).
		Limit(limit).
		Sort("created_at", -1).
		Decode(&users).
		Filter(filter).
		Find()

	if err != nil {
		panic(err)
	}
	return users, paginatedData.Pagination
}

func CreateATodo(createTodoDto dto.CreateTodoRequest) models.Todo {
	user := models.Todo{
		Id:        primitive.NewObjectID(),
		Task:      createTodoDto.Task,
		Status:    createTodoDto.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := models.TodoCollection.InsertOne(user)
	if err != nil || result == nil {
		panic(err)
	}

	return user
}

func UpdateATodo(userId string, updateTodoDto dto.UpdateTodoRequest) (models.Todo, error) {

	objId, _ := primitive.ObjectIDFromHex(userId)

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := models.TodoCollection.FindOneAndUpdate(
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.M{
				"task":       updateTodoDto.Task,
				"status":     updateTodoDto.Status,
				"updated_at": time.Now(),
			}},
		},
		&opt,
	)

	if result.Err() != nil {
		log.Println("Err ", result.Err())
		return models.Todo{}, result.Err()
	}

	var user models.Todo
	if err := result.Decode(&user); err != nil {
		return models.Todo{}, err
	}

	return user, nil
}

func ATodo(userId string) models.Todo {
	var user models.Todo

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := models.TodoCollection.FindOne(bson.M{"_id": objId}).Decode(&user)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return user
}

func DeleteATodo(userId string) bool {
	objId, _ := primitive.ObjectIDFromHex(userId)

	result := models.TodoCollection.FindOneAndDelete(bson.D{{"_id", objId}})

	if result.Err() != nil {
		return false
	}

	return true
}
