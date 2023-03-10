package api

import (
	"database/sql"
	"net/http"

	db "github.com/CRAZYKAYZY/web/db/sqlc"
	"github.com/gin-gonic/gin"
)

// CREATE USER CONTROL
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

// GET SPECIFIC USER CONTROL
type GetUserRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, db.GetUserParams{ID: req.ID})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

// LIST USERS CONTROL
type ListAllUsers struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req ListAllUsers
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// UPDATE SPECIFIC USER CONTROL
// separation of concerns for the update function: UserURIReq for binding uri, and UpdateUserReq for binding json structs.
type UserURIReq struct {
	ID int32 `uri:"id" binding:"required"`
}

type UpdateUserReq struct {
	Username       string `json:"username" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var reqUri UserURIReq
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req UpdateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user := db.UpdateUserParams{
		ID:             reqUri.ID,
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
	}

	userUpdate, err := server.store.UpdateUser(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, userUpdate)
}

// DELETE SPECIFIC USER CONTROL
type DeleteUserRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req DeleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	err := server.store.DeleteUser(ctx, int32(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
