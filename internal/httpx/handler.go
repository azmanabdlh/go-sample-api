package httpx

import (
	"errors"
	"strconv"

	"github.com/azmanabdlh/go-sample-api/internal/book"
	"github.com/azmanabdlh/go-sample-api/internal/provider"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *book.Service
	jwt     provider.JsonWebToken
}

type SearchFilter struct {
	Title  string
	Author string

	Page  int
	Limit int

	Sort  string
	Order string
}

func NewHandler(service *book.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type TokenRequest struct {
	UserID string `json:"user_id"`
}

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	book, err := h.service.Create(
		c.Request.Context(),
		req.Title,
		req.Author,
		req.Year,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Book Created",
		"id":      book.ID,
	})
}

func (h *Handler) FindByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.FindByID(
		c.Request.Context(),
		id,
	)

	if errors.Is(err, book.ErrBookNotFound) {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, data)
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := h.service.Update(
		c.Request.Context(),
		id,
		req.Title,
		req.Author,
		req.Year,
	)

	// if errors.Is(err, book.ErrBookNotFound) {
	// 	c.JSON(404, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, data)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Book Deleted",
	})
}

func (h *Handler) Search(c *gin.Context) {
	query := c.Query("author")

	page, _ := strconv.Atoi(
		c.DefaultQuery("page", "1"),
	)

	limit, _ := strconv.Atoi(
		c.DefaultQuery("limit", "10"),
	)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	books, _, err := h.service.Search(
		c.Request.Context(),
		query,
		limit,
		page,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})

		return
	}

	if books == nil {
		books = []book.Book{}
	}

	c.JSON(200, books)

}

func (h *Handler) GenerateToken(c *gin.Context) {
	var req TokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"message": "invalid body",
		})
		return
	}

	if req.UserID == "" {
		req.UserID = "user-1"
	}

	token, err := h.jwt.GenerateToken(
		req.UserID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"message": "failed generate token",
		})

		return
	}

	c.JSON(200, gin.H{
		"message":      "success",
		"access_token": token,
		"token_type":   "Bearer",
		"user_id":      req.UserID,
	})

}
