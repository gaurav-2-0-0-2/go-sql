package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"go-sql/controllers"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got / req\n")
	fmt.Printf("got a req on /\n")
}

func main() {
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/users", controllers.GetUsers)
	fmt.Println("Server started at http://localhost:4000")
	ServerErr := http.ListenAndServe(":4000", nil)
	if errors.Is(ServerErr, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if ServerErr != nil {
		fmt.Printf("error starting server: %s\n", ServerErr)
		os.Exit(1)
	}
}
