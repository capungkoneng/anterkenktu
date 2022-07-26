package api

import (
	"database/sql"
	// "fmt"

	"net/http"

	db "github.com/capungkoneng/anterkenktu/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type gambarMobil struct {
	Nama string        `json:"nama"`
	Url  []interface{} `json:"url"`
}

func tes() string {

	return "{teee}"
}

func resMobil(gambar []db.ListMobilNewRow) gambarMobil {
	var url []interface{}

	for _, row := range gambar {
		url = append(url, map[string]string{"name": row.Url})

	}

	return gambarMobil{
		Url: url,
		// Username:          user.Username,
		// FullName:          user.FullName,
		// Email:             user.Email,
		// PasswordChangedAt: user.PasswordChangedAt,
		// CreatedAt:         user.CreatedAt,
	}
}

type mobilResponse struct {
	Nama   string      `json:"nama"`
	Gambar gambarMobil `json:"gambar"`
}

// type getMobilRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// type ListMobil struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=1,max=1"`
// }

//Get akun list
func (server *Server) GetListMobil(ctx *gin.Context) {
	// var req getMobilRequest
	// if err := ctx.ShouldBindQuery(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, (err))
	// 	return
	// }
	// if err := ctx.ShouldBindUri(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, (err))
	// 	return
	// }
	gambarr, err := server.store.ListMobilNew(ctx)
	// mobilsemua, err := server.store.GetMobilJoinMany(ctx)
	// arg := db.GetMobilJoinManyParams{
	// 	Limit:  req.PageID,
	// 	Offset: (req.PageID - 1) * req.PageSize,
	// }

	// mobil, err := server.store.GetMobilJoinMany(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	rsp := mobilResponse{
		Gambar: resMobil(gambarr),
	}

	ctx.JSON(http.StatusOK, rsp)

	// ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"mobil": mobilsemua}, "imh": gin.H{"image": gambarr}})

}

type createMobilRequest struct {
	Nama       string         `json:"nama"`
	Deskripsi  sql.NullString `json:"deskripsi"`
	KategoriID int64          `json:"kategori_id"`
	UserID     string         `json:"user_id"`
	Gambar     []string       `json:"gambar"`
	Trf6jam    int64          `json:"trf_6jam"`
	Trf12jam   int64          `json:"trf_12jam"`
	Trf24jam   int64          `json:"trf_24jam"`
	Seat       sql.NullInt64  `json:"seat"`
	TopSpeed   sql.NullInt64  `json:"top_speed"`
	MaxPower   sql.NullInt64  `json:"max_power"`
	Pintu      sql.NullInt64  `json:"pintu"`
	Gigi       sql.NullString `json:"gigi"`
}

// Create User one
func (server *Server) CreateMobil(ctx *gin.Context) {
	var req createMobilRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, (err))
		return
	}

	arg := db.CreateMobilParams{
		Nama:       req.Nama,
		Deskripsi:  req.Deskripsi,
		KategoriID: req.KategoriID,
		UserID:     req.UserID,
		Gambar:     req.Gambar,
		Trf6jam:    req.Trf6jam,
		Trf12jam:   req.Trf12jam,
		Trf24jam:   req.Trf24jam,
		Seat:       req.Seat,
		TopSpeed:   req.TopSpeed,
		MaxPower:   req.MaxPower,
		Pintu:      req.Pintu,
		Gigi:       req.Gigi,
	}

	mobil, err := server.store.CreateMobil(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, (err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, (err))
		return
	}
	ctx.JSON(http.StatusOK, mobil)
}
