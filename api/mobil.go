package api

import (
	"net/http"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListMobil struct {
	Nama     string `json:"nama"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1,max=10"`
}

//Get akun list
func (server *Server) GetListMobil(ctx *gin.Context) {
	var req ListMobil
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}

	arg := db.GetMobilJoinManyParams{
		Nama:   req.Nama,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	mobil, err := server.store.GetMobilJoinMany(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"mobil": mobil}})

}
