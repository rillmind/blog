package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service
}

func NewController(service Service) Controller {
	return Controller{
		service,
	}
}

func (pc *Controller) CreatePost(ctx *gin.Context) {
	post := Model{}

	err := ctx.BindJSON(&post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedPost, err := pc.Service.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	postRes := ModelResponse{
		ID:      insertedPost.ID,
		Title:   insertedPost.Title,
		Slug:    insertedPost.Slug,
		Content: insertedPost.Content,
	}

	ctx.JSON(http.StatusCreated, postRes)
}

func (pc *Controller) ReadPosts(ctx *gin.Context) {
	posts, err := pc.Service.ReadPosts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	postsRes := make([]ModelResponse, 0, len(posts))

	for _, post := range posts {
		postModelResponse := ModelResponse{
			ID:        post.ID,
			Title:     post.Title,
			Slug:      post.Slug,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			DeletedAt: post.DeletedAt,
		}

		postsRes = append(postsRes, postModelResponse)
	}

	ctx.JSON(http.StatusOK, postsRes)
}
