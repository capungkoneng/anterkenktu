package api

import (
	"net/http"

	// db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

// type ListKategori struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
// }

// //Get akun list
// func (server *Server) GetListKategori(ctx *gin.Context) {
// 	var req ListKategori
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, (err))
// 		return
// 	}

// 	arg := db.ListKategoriParams{
// 		Limit:  req.PageSize,
// 		Offset: (req.PageID - 1) * req.PageSize,
// 	}
// 	akun, err := server.store.ListKategori(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, (err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, akun)

// }

func GetHandlerHallo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "Hellow word",
	})
}
