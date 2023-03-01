package api

import (
	"net/http"

	db "github.com/CRAZYKAYZY/web/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username       string `json:"username" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}
