package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"log"
	"go-sql/db"
	"go-sql/controllers"
	"go-sql/middleware"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got / req\n")
	fmt.Printf("got a req on /\n")
}

func main() {
	// Connecting Database
	err := db.Connectdb("./db/app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()
	// Routes Handling
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/users", middleware.CORS(controllers.GetAllUsers))
	http.HandleFunc("/create/user", middleware.CORS(controllers.CreateUser))
	http.HandleFunc("/user/", middleware.CORS(controllers.GetUserById))
	// Starting Server
	fmt.Println("Server started at http://localhost:4000")
	ServerErr := http.ListenAndServe(":4000", nil)
	if errors.Is(ServerErr, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if ServerErr != nil {
		fmt.Printf("error starting server: %s\n", ServerErr)
		os.Exit(1)
	}
}
