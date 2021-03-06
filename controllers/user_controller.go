package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/free-go-foundation/go-auth/services"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var GetUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	username := user.(*jwt.Token).Claims.(jwt.MapClaims)["username"]

	retrievedUser, err := DataBase.FindUserByEmail(username.(string))
	if err != nil {
		log.Error(err)
		http.Error(w, "InvalidUserError", http.StatusBadRequest)
		return
	}

	payload, _ := json.Marshal(retrievedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write([]byte(payload))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})

var CreateUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	newUser, creationErr := services.CreateUser(r)
	if creationErr != nil {
		log.Error(creationErr)
		http.Error(w, creationErr.Error(), http.StatusBadRequest)
		return
	}

	token, err := services.GenerateJWTToken(newUser.Username)
	if err != nil {
		log.Error(err)
		http.Error(w, "TokenCreationError", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte(token))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})

var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	err := services.DeleteUserById(id)
	if err != nil {
		log.Error(err)
		http.Error(w, "UserDeletionError", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte("User Deleted"))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})
