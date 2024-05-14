package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rohitkrcodes/go_aggregator/internal/database"
)

type feedFollowParam struct {
	FeedID uuid.UUID `json:"feed_id"`
}

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	params := feedFollowParam{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create feed follower: %s", err))
		return
	}

	respondWithJSON(w, 201, ModelWrapperFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not get feed follows: %s", err))
		return
	}

	respondWithJSON(w, 201, ModelWrapperAllFeedFollows(feeds))
}

func (apiCfg *apiConfig) handlerDeleteFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFolowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFolowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not parse feed follow id: %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not delete feed follow: %s", err))
		return
	}

	respondWithJSON(w, 201, struct{}{})
}
