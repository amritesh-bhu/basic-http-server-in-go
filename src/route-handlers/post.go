package routehandlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go-mongo-crud/src/domain"
)

func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts := domain.GetAllPosts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func HandleGetPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	post, ok := domain.GetPostByID(id)
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var p domain.Post
	json.NewDecoder(r.Body).Decode(&p)

	post := domain.CreatePost(p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var p domain.Post
	json.NewDecoder(r.Body).Decode(&p)

	post, ok := domain.UpdatePost(id, p)
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	ok := domain.DeletePost(id)
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post deleted"))
}
