package controllers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/free-go-foundation/go-auth/models"
	"github.com/free-go-foundation/go-auth/services"
)

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	requestBody := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		log.Error(err)
		http.Error(w, "RequestParseError", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	requestUser, err := DataBase.FindUserByEmail(requestBody.Username)
	if err != nil {
		log.Error(err)
		http.Error(w, "InvalidUserError", http.StatusBadRequest)
		return
	}

	token, err := services.IssueToken(&requestUser, requestBody.Password)
	if err != nil {
		log.Error(err)
		http.Error(w, "TokenCreationError", http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	_, err = w.Write([]byte(token))
	if err != nil {
		log.Error(err)
		http.Error(w, "ResponseError", http.StatusInternalServerError)
	}
})
