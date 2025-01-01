package auth

import (
	"GoogleProject/db"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt" // Correct import for bcrypt
)

type registerBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Region   string `json:"region"`
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var requestBody registerBody
	client := db.GetClient()
	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return
	}

	// Create a new user document
	newUser := bson.M{
		"username": requestBody.Username,
		"email":    requestBody.Email,
		"password": string(hashedPassword),
		"region":   requestBody.Region,
	}

	// Insert the document into the "Users" collection
	collection := client.Database("GoogleBuisn").Collection("Users")
	_, err = collection.InsertOne(r.Context(), newUser)            
	if err != nil {
		http.Error(w, "Failed to add data to the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully!")
}


func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var requestBody loginBody

	client := db.GetClient()

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	collection := client.Database("GoogleBuisn").Collection("Users")
	result, err := findUser(w, collection, requestBody.Username)
	if err != nil {
		http.Error(w, "Couldn't find the user", http.StatusInternalServerError)
		return
	}

	if VerifyPassword(result["password"].(string), requestBody.Password) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Login correct")
	} else {
		http.Error(w, "Trouble with verifying password", http.StatusBadRequest)
	}
}
func findUser(w http.ResponseWriter , collection *mongo.Collection, username string) (bson.M, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var result bson.M
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusBadRequest)
		} else {
			http.Error(w, "Problem with finding user", http.StatusInternalServerError)
		}
	}
	return result, err
}

func VerifyPassword(hashedPassword string , password string ) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
