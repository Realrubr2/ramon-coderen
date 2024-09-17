package middleware

import (
	"errors"
	"net/http"

	"github.com/realrubr2/golang/api"
	"github.com/realrubr2/golang/internal/tools"
	"github.com/realrubr2/ramon-coderen/api"
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
			api.internalErrorHandler(w)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if(loginDetails == nil || (token !=(*loginDetails).AuthToken)){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}
		next.ServeHTTP(w,r)
	})
}