package routehandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"example.com/go-mongo-crud/src/domain"
	"example.com/go-mongo-crud/src/helper"
)

func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts := domain.GetAllPosts()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "all posts",
		"posts":   posts,
	})
}

func HandleGetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		helper.WriteError(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	post, err := domain.GetPostByID(id)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"messgage": fmt.Sprintf("post for id: %s", id),
		"post":     post,
	})
}

func HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	var p domain.Post
	json.NewDecoder(r.Body).Decode(&p)

	if p.Title == "" || p.Content == "" {
		helper.WriteError(w, "content or title is empty", http.StatusBadRequest)
		return
	}

	post := domain.CreatePost(p)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "post created",
		"post":    post,
	})
}

func HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		helper.WriteError(w, "missing parameter id", http.StatusBadRequest)
		return
	}

	var p domain.Post
	errJson := json.NewDecoder(r.Body).Decode(&p)

	if errJson != nil {
		helper.WriteError(w, "invalid json body", http.StatusInternalServerError)
		return
	}

	post, err := domain.UpdatePost(id, p)
	if err != nil {
		helper.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "post updated",
		"post":    post,
	})
}

func HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		helper.WriteError(w, "missing parameter id", http.StatusBadRequest)
		return
	}

	post := domain.DeletePost(id)
	if post != nil {
		helper.WriteError(w, post.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "post deleted",
	})
}
