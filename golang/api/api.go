package api

import (
	"encoding/json"
	"net/http"
)


type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//Sucsess code 200
	Code int

// the balance
	Balance int64
}


type Error struct{
	// error code
	Code int
// error message
	Message string
}

func WriteError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-type", "application/json");
	w.WriteHeader(code);
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		WriteError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		WriteError(w, "an unexpected error occured",http.StatusInternalServerError)
	}
)