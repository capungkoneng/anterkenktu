package routes

import (
	"github.com/capungkoneng/anterkenktu/controllers"
	"github.com/gin-gonic/gin"
)

type KategoriRoutes struct {
	kategoricontroller controllers.KategoriController
}

func NewKategoriRoutes(kategoricontroller controllers.KategoriController) KategoriRoutes {
	return KategoriRoutes{kategoricontroller}
}

func (rc *KategoriRoutes) KategoriRoute(rg *gin.RouterGroup) {
	router := rg.Group("kategori")
	router.GET("/list/", rc.kategoricontroller.GetListKategori)
}
