package main

import (
	"fmt"
	"net/http"

	routehandlers "example.com/go-mongo-crud/src/route-handlers"
)

func main() {

	http.HandleFunc("/posts", routehandlers.HandleGetPosts)
	http.HandleFunc("/post", routehandlers.HandleGetPost)
	http.HandleFunc("/create", routehandlers.HandleCreatePost)
	http.HandleFunc("/update", routehandlers.HandleUpdatePost)
	http.HandleFunc("/delete", routehandlers.HandleDeletePost)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at port 8080")
}
