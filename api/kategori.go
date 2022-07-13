package api

import (
	"net/http"

	// db "github.com/capungkoneng/anterkenktu/db/sqlc"
	db "github.com/capungkoneng/anterkenktu/db/sqlc"
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
type KategoriController struct {
	db *db.Queries
}
type createKategRequest struct {
	NamaKategori string `json:"nama_kategori" binding:"required"`
	Deskripsi    string `json:"deskripsi" binding:"required,max=100"`
}

// func NewKategoriController(db *db.Queries) *KategoriController {
// 	return &KategoriController{db}
// }
func CreateKategori(ctx *gin.Context) {
	var a *KategoriController
	var req createKategRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := &db.CreateKategoriParams{
		NamaKategori: req.NamaKategori,
		Deskripsi:    req.Deskripsi,
	}

	kategori, err := a.db.CreateKategori(ctx, *arg)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"kategori": kategori}})
}
