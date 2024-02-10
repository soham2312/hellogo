package main

import (
	"net/http"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {	
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&name)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json,%v",err))
		return
	}

	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if nil != err {
		respondWithError(w, 400, fmt.Sprintf("Error creating user,%v",err))
		return
	}

	respondwithJSON(w, 200, user)

}