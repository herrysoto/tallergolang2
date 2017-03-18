package main

import (
	controller "taller/rest"
	//datos "./datos"
	"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
)
// Cors funcion para activar CORS
func Cors() gin.HandlerFunc {

	//http://stackoverflow.com/questions/31834408/golang-gin-and-go-socket-io-cors-issue
	return func(c *gin.Context) {
		fmt.Println("CORS middleware loaded...")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, WS, WSS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(Cors())
    v1 := r.Group("api/v1")
	{
		v1.GET("/firstcombo", controller.Getfirstcombo)
		v1.GET("/Getcodigovalidacion/:Codopera/:Codoperaservicio",controller.Getcodigovalidacion)
		v1.GET("/HorasHombres",controller.GetHorasHombres)
		v1.PUT("/PutHorasHombre", controller.PutHorasHombre)
		v1.GET("/Getnumcodigo",controller.Getnumcodigo)
		v1.POST("/PostHorasHombres", controller.PostHorasHombres)
		v1.GET("/firstcombo2", controller.Getfirstcombo2)
		v1.GET("/secondcomb/:CodServicio", controller.Getsecondcombo)
		v1.GET("/grid/:Operserv", controller.Getgrid)
		v1.GET("/BuscarOperServiciosParam/:param", controller.BuscarOperServiciosParam)
		v1.GET("/LISTAROPERSERVCONTENIDOS/:param/:param1", controller.LISTAR_OPERSERVCONTENIDOS)
		v1.POST("/post", controller.Postgrid)
		v1.PUT("/PutOperaServi", controller.PutOperaServi)
		v1.GET("/Getnumcodigoop", controller.Getnumcodigoop)
		// v1.PUT("/put/:putOpeSer/:putOpeSer1",controller.PutOperaServi)
	}

	r.Run(":8081")
}