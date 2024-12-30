package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type registerBody struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Region string  `json:"region"`
}
type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Region string  `json:"region"`
}

func Handleregister(w http.ResponseWriter, r *http.Request) {
	var requestBody registerBody

	err := json.NewDecoder(r.Body).Decode(&requestBody);
	if err != nil {
		fmt.Println(err.Error())
	}

	
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {}
