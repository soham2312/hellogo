package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/google/uuid"
	"github.com/soham2312/rsagg/internal/database"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {	
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		resonsewitherror(w, 400, fmt.Sprintf("Error parsing json,%v",err))
		return
	}

	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if nil != err {
		resonsewitherror(w, 400, fmt.Sprintf("Error creating user,%v",err))
		return
	}

	respondwithJSON(w, 200, databaseUserToUser(user))

}