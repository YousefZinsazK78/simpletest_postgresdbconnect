package api

import (
	"encoding/json"
	"migrationtest/db/userstore"
	"net/http"
)

type userHandler struct {
	userstorer userstore.UserStorer
}

func NewUserHandler(userstorer userstore.UserStorer) *userHandler {
	return &userHandler{
		userstorer: userstorer,
	}
}

func (u *userHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	users, err := u.userstorer.GetUsers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"result": users,
	})
}

// func (u *userHandler) HandleInsert(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		// users, err := u.userstorer.InsertUser()
// 		// if err != nil {
// 		// 	w.Header().Set("Content-Type", "application/json")
// 		// 	w.WriteHeader(http.StatusInternalServerError)
// 		// 	json.NewEncoder(w).Encode(map[string]string{
// 		// 		"error": err.Error(),
// 		// 	})
// 		// 	return
// 		// }
// 		// w.Header().Set("Content-Type", "application/json")
// 		// w.WriteHeader(http.StatusOK)
// 		// json.NewEncoder(w).Encode(map[string]any{
// 		// 	"result": users,
// 		// })
// 	}

// }
