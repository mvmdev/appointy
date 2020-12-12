package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvmdev/appointy/helper"
	"github.com/mvmdev/appointy/models"
	"github.com/mvmdev/appointy/new_helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionUser = helper.ConnectDB()
var collectionContact = new_helper.ConnectDB()

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collectionUser.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	result, err := collectionUser.InsertOne(context.TODO(), user)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)

	result, err := collectionContact.InsertOne(context.TODO(), contact)

	if err != nil {
		new_helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collectionContact.FindOne(context.TODO(), filter).Decode(&contact)

	if err != nil {
		new_helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/users", createUser).Methods("POST")
	route.HandleFunc("/users/{id}", getUser).Methods("GET")
	route.HandleFunc("/contacts", createContact).Methods("POST")
	route.HandleFunc("/contacts/{id}", getContact).Methods("GET")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, route))

}
