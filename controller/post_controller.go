package controller

import (
	
	"encoding/json"
	"net/http"
	"go-rest-api/entity"
	"go-rest-api/repository"
	"go-rest-api/service"
	
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type postController struct{}

var (
	postService service.PostService = service.NewPostService(repository.NewFirestorePostRepository())
)

func NewPostController() PostController {
	return &postController{}
}
func (p *postController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"errors": "Error getting posts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	// Encode writes the JSON encoding of v to the stream, followed by a newline character.
	json.NewEncoder(resp).Encode(posts)
}

func (p *postController) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"errors": "marshall error"}`))
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"errors": "marshall error"}`))
		return
	}
	_, err = postService.Create(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"errors": "marshall error"}`))
		return
	}
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"errors": "Error saving posts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}