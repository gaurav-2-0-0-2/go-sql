package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"log"
	"go-sql/db"
	"go-sql/controllers"
)


func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got / req\n")
	fmt.Printf("got a req on /\n")
}

func main() {
	err := db.Connectdb("./db/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/users", controllers.GetAllUsers)
	http.HandleFunc("/create/user", controllers.CreateUser)
	fmt.Println("Server started at http://localhost:4000")
	ServerErr := http.ListenAndServe(":4000", nil)
	if errors.Is(ServerErr, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if ServerErr != nil {
		fmt.Printf("error starting server: %s\n", ServerErr)
		os.Exit(1)
	}
}
