package api

import (
	"github.com/dilshodforever/restaurant-submodule/api/handler"
	_ "github.com/dilshodforever/restaurant-submodule/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8080
// @BasePath /
func NewGin(h *handler.Handler)*gin.Engine{
	r:=gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	u:=r.Group("/restoran")
	u.POST("/create", h.CreateRestoran)
	u.PUT("/update/:id", h.UpdateRestoran)
	u.DELETE("/delete/:id", h.DeleteRestoran)
	u.GET("/getall", h.GetAllRestoran)
	u.GET("/getbyid/:id", h.GetByIdRestoran)
	
	re:=r.Group("/reservation")
	re.POST("/create/:user_id/:restaurant_id", h.CreateReservation)
	re.PUT("/update/:id", h.UpdateReservation)
	re.DELETE("/delete/:id", h.DeleteReservation)
	re.GET("/getall", h.GetAllReservation)
	re.GET("/getbyid/:id", h.GetbyIdReservation)

	res:=r.Group("/reservationOrder")
	res.POST("/create", h.CreateReservationOrder)
	res.PUT("/update/:id", h.UpdateReservationOrder)
	res.DELETE("/delete/:id", h.DeleteReservationOrder)
	res.GET("/getall", h.GetAllReservationOrder)
	res.GET("/getbyid/:id", h.GetByIdReservationOrder)
	return r
}