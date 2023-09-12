package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/minhdao911/rss-aggregator-go/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	decoderErr := decoder.Decode(&params)
	if decoderErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", decoderErr))
		return
	}

	user, dbErr := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if dbErr != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", dbErr))
		return
	}

	respondWithJSON(w, 200, dbUserToUser(user))
}