package api

import (
	"fmt"
	"net/http"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ListKategori struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

//Get akun list
func (server *Server) GetListKategori(ctx *gin.Context) {
	fmt.Println("p")
	var req ListKategori
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}
	fmt.Println("L")

	arg := db.ListKategoriParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	fmt.Println("q")

	kategori, err := server.store.ListKategori(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, kategori)

}

type createKetegRequest struct {
	NamaKategori string `json:"nama_kategori"`
	Deskripsi    string `json:"deskripsi"`
}

func (server *Server) CreateKategori(ctx *gin.Context) {
	var req createKetegRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}

	arg := db.CreateKategoriParams{
		NamaKategori: req.NamaKategori,
		Deskripsi:    req.Deskripsi,
	}

	kategori, err := server.store.CreateKategori(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, (err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"kategori": kategori}})
}
