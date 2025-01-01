package controllers

import (
	"encoding/json"
	"net/http"
	"go-sql/db"
)

type User struct {
	Id    int		`json:"id"`
	Name  string	`json:"name"`
	Email string	`json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id,name,email FROM users;")
	if err != nil {
		http.Error(w, "Failed to query databse", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Failed to read rows", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
