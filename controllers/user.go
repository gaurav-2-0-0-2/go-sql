package controllers

import (
	"encoding/json"
	"net/http"
	"go-sql/db"
	"fmt"
)

type User struct {
	Id    	int			`json:"id"`
	Name  	string	`json:"name"`
	Email  	string	`json:"email"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w,"Person: %v\n", u)

	_, err = db.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?);", u.Name, u.Email)
	if err != nil {
		http.Error(w, "Failed to insert into db", http.StatusInternalServerError)
		return
	}
	fmt.Println("User created")
}
