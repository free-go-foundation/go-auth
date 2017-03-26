package main

import (
	"net/http"
	"os"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/free-go-foundation/go-auth/config"
	"github.com/free-go-foundation/go-auth/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(mConfig.Signing.SecretKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

var mConfig *config.Config = nil

func init() {
	mConfig = config.DevConfig
}

func main() {
	r := mux.NewRouter()

	r.Handle("/user/", jwtMiddleware.Handler(controllers.GetUser)).Methods("GET")
	r.Handle("/user", controllers.CreateUser).Methods("POST")
	r.Handle("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.Handle("/login", controllers.Login).Methods("POST")

	http.ListenAndServe(mConfig.Env.Port, handlers.LoggingHandler(os.Stdout, r))
}

