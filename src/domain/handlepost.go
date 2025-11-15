package domain

import (
	"errors"
	"sync"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts   = make(map[int]Post)
	postsMu sync.Mutex
	lastID  = 0
)

func CreatePost(p Post) Post {
	postsMu.Lock()
	defer postsMu.Unlock()

	lastID++
	p.ID = lastID
	posts[p.ID] = p

	return p
}

func GetAllPosts() []Post {
	postsMu.Lock()
	defer postsMu.Unlock()

	result := make([]Post, 0, len(posts))
	for _, p := range posts {
		result = append(result, p)
	}
	return result
}

func GetPostByID(id int) (Post, error) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		return Post{}, errors.New("Post not found")
	}
	return p, nil
}

func UpdatePost(id int, p Post) (Post, error) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, exists := posts[id]
	if !exists {
		return Post{}, errors.New("Post not found")
	}

	p.ID = id
	posts[id] = p
	return p, nil
}

func DeletePost(id int) error {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, exists := posts[id]
	if !exists {
		return errors.New("Post not found")
	}

	delete(posts, id)
	return nil
}
