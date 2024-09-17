package middleware

import (
	"errors"
	"net/http"

	"github.com/Realrubr2/ramon-coderen/golang/api"
	"github.com/Realrubr2/ramon-coderen/golang/internal/tools"
	log "github.com/sirupsen/logrus"
)


var unAuthorizedError = errors.New("invalid user or token");
func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username");
		var token = r.Header.Get("authorrization")
		var err error
		if username == "" || token ==""{
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil{
			api.InternalErrorHandler(w)
			return
		}

		var LoginDetails *tools.LoginDetails
		LoginDetails = (*database).GetUserLoginDetails(username)

		if(LoginDetails == nil || (token !=(*LoginDetails).AuthToken)){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}
		next.ServeHTTP(w,r)
	})
}