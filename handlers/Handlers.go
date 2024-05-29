package handlers

import (
	"encoding/json"
	"github.com/janicaleksander/GoInterviewTask1/types"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func MakeHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadGateway, types.ErrorAPI{Error: types.FUNPROB})
		}
	}
}

func HandleGetJoke(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return WriteJson(w, http.StatusBadRequest, types.ErrorAPI{types.UNS})
	}

	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, types.ErrorAPI{types.EAPI})
	}
	defer resp.Body.Close()

	newJoke := new(types.Joke)
	if err = json.NewDecoder(resp.Body).Decode(newJoke); err != nil {
		return WriteJson(w, http.StatusBadRequest, types.ErrorAPI{types.EAPI})
	}

	return WriteJson(w, http.StatusOK, newJoke)
}

func WriteJson(w http.ResponseWriter, statusCode int, val any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(val)
}
