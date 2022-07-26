package controllers

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/capungkoneng/anterkenktu/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/pq"
)

type TokenM struct {
	config util.Config
}

type AuthController struct {
	db *db.Queries
}

func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(usr db.User) userResponse {
	return userResponse{
		Username:          usr.Username,
		FullName:          usr.FullName,
		Email:             usr.Email,
		PasswordChangedAt: usr.PasswordChangedAt,
		CreatedAt:         usr.CreatedAt,
	}
}

func (c *AuthController) SignUpUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, (err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := c.db.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, (err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, (err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccesToken string       `json:"acces_token"`
	User       userResponse `json:"user"`
}

func (c *AuthController) LoginUser(ctx *gin.Context) {
	var (
		req    loginUserRequest
		config TokenM
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userget, err := c.db.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, (err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, (err))
		return
	}

	err = util.CheckPassword(req.Password, userget.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, (err))
	}

	// Generate Token Jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userget.Username,
		"exp": config.config.AccessTokenDuration,
	})
	tokenString, err := token.SignedString([]byte(config.config.TokenSymmetricKey))
	if err != nil {
		return
	}
	rsp := loginUserResponse{
		AccesToken: tokenString,
		User:       newUserResponse(userget),
	}
	ctx.JSON(http.StatusOK, rsp)
}
