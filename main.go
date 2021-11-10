package main

import (
	"context"
	"encoding/json"
	"fmt"
	"mux"
	"net/http"

	"github.com/stretchr/testify/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("abc")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("could not connect to mongo")
	}

	fmt.Println("connected to momgoDB")

	collection := client.Database("test").Collection("UserManagement")
}

func main() {
	//connect to DB
	ConnectToMongo()

	//set up the router
	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	http.ListenAndServe("", router)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user model.User

	_ := json.NewDecoder(req.Body).Decode(&user)

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("Record not inserted")
	}

	json.NewEncoder(res).Encode(result)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	var user model.User

	userID := mux.Vars(r)["id"]

	filter := bson.M{"id": userID}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		fmt.Println("Record not found")
	}
}


func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var user model.User

	userID := mux.Vars(r)["id"]

	filter := bson.M{"id": userID}

	_ := json.NewDecoder(req.Body).Decode(&user)

	newUserData := bson.D{
		{
			"$set": bson.D{
				"id": user.ID,
				"name": user.Name,
				"email": user.Email,
			}
		}
	}

	err := collection.FindAndUpdate(context.TODO(), filter, newUserData).Decode(&user)

	if err != nil {
		fmt.Println("Record not updated")
	}
}