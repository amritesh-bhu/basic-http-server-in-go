package domain

import (
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

func GetPostByID(id int) (Post, bool) {
	postsMu.Lock()
	defer postsMu.Unlock()

	p, ok := posts[id]
	if !ok {
		return Post{}, false
	}
	return p, true
}

func UpdatePost(id int, p Post) (Post, bool) {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, exists := posts[id]
	if !exists {
		return Post{}, false
	}

	p.ID = id
	posts[id] = p
	return p, true
}

func DeletePost(id int) bool {
	postsMu.Lock()
	defer postsMu.Unlock()

	_, exists := posts[id]
	if !exists {
		return false
	}

	delete(posts, id)
	return true
}
