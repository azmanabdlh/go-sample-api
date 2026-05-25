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
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})

		return
	}

	book, err := h.service.Create(
		c.Request.Context(),
		req.Title,
		req.Author,
	)

	if err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Book Created!!",
		Code:    201,
		Data:    book,
	})
}

func (h *Handler) FindByID(c *gin.Context) {
	id := c.Param("id")

	data, err := h.service.FindByID(
		c.Request.Context(),
		id,
	)

	if errors.Is(err, book.ErrBookNotFound) {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    404,
		})
		return
	}

	if err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Book ok",
		Code:    200,
		Data:    data,
	})
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	data, err := h.service.Update(
		c.Request.Context(),
		id,
		req.Title,
		req.Author,
	)

	if errors.Is(err, book.ErrBookNotFound) {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    404,
		})
		return
	}

	if err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Book Updated!",
		Code:    200,
		Data:    data,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(
		c.Request.Context(),
		id,
	)

	if err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Book Deleted",
		Code:    200,
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

	books, total, err := h.service.Search(
		c.Request.Context(),
		query,
		limit,
		page,
	)

	if err != nil {
		RespondJSON(c, Response{
			Message: err.Error(),
			Code:    500,
		})
		return
	}

	if books == nil {
		books = []book.Book{}
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Successfully",
		Code:    200,
		Data: gin.H{
			"books": books,
			"meta": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	})
}

func (h *Handler) GenerateToken(c *gin.Context) {
	var req TokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		RespondJSON(c, Response{
			Message: "invalid body",
			Code:    400,
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
		RespondJSON(c, Response{
			Message: "failed generate token",
			Code:    500,
		})

		return
	}

	RespondJSON(c, Response{
		Success: true,
		Message: "Token ok",
		Data: gin.H{
			"access_token": token,
			"token_type":   "Bearer",
			"user_id":      req.UserID,
		},
		Code: 200,
	})
}
