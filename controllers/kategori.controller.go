package controllers

import (
	"net/http"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type KategoriController struct {
	db *db.Queries
}

func NewKategoriController(db *db.Queries) *KategoriController {
	return &KategoriController{db}
}

type ListKategoriParam struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

//Get akun list
func (c *KategoriController) GetListKategori(ctx *gin.Context) {

	var req ListKategoriParam
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	arg := db.ListKategoriParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	kategori, err := c.db.ListKategori(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"kategori": kategori}})

}
