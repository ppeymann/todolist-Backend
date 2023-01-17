package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todoList/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "todoList"
const colName = "myTodo"

//! connecting to DB

var collection *mongo.Collection

func init(){
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connection is Successful...")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is READY ...")
}

// * HELPER *

func insertTodo(todo model.Todos){
	inserted, err := collection.InsertOne(context.Background(), todo)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Insert One Todo",inserted.InsertedID)
}

func updateTodo(todoId string){
	id, err := primitive.ObjectIDFromHex(todoId)
	if err!=nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id":id}
	update:=  bson.M{"$set":bson.M{"finished":true}}

	result, err := collection.UpdateOne(context.Background(), filter,update)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Update is Successful",result.ModifiedCount)
}

func deleteOneTodo(todoId string){
	id, err := primitive.ObjectIDFromHex(todoId)
	if err!=nil{
		log.Fatal(err)
	}

	filter := bson.M{"_id":id}
	res, err := collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Delete is Successfil",res.DeletedCount)
}

func deleteAllTodo()int64{
	res, err := collection.DeleteMany(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Delete All is Successful",res.DeletedCount)
	return res.DeletedCount
}

func getAllTodo()[]primitive.M{
	curs, err := collection.Find(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var allTodo []primitive.M

	for curs.Next(context.Background()){
		var myTodo bson.M
		err:=curs.Decode(&myTodo)
		if err!=nil{
			log.Fatal(err)
		}

		allTodo = append(allTodo, myTodo)
	}
	defer curs.Close(context.Background())

	return allTodo
}

func GetTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")

	allTodo:= getAllTodo()

	json.NewEncoder(w).Encode(allTodo)
}

func CreateOneTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var todo model.Todos
	_ = json.NewDecoder(r.Body).Decode(&todo)
	insertTodo(todo)
	json.NewEncoder(w).Encode(todo)
}

func MarkTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")

	params := mux.Vars(r)

	updateTodo(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params := mux.Vars(r)

	deleteOneTodo(params["id"])
}

func DeleteTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count := deleteAllTodo()
	json.NewEncoder(w).Encode(count)
}