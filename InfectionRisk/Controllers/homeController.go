package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "github/marogosteen/InfectionRisk/InfectionRisk/Models"
)

type HomeController struct {
	Port string
}

func (c *HomeController) RunServer() error {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	router.Static("/scripts", "InfectionRisk/Views/Scripts")
	router.LoadHTMLGlob("InfectionRisk/Views/*.html")

	router.GET("/", homeHandler)
	router.GET("/result/", resultHandler)

	router.Run(":" + c.Port)

	return nil
}

func homeHandler(c *gin.Context) {
	m := models.NewRiskFromViewModel()
	c.HTML(200, "index.html", m)
}

func resultHandler(c *gin.Context) {
	m, err := models.ConvertResultViewModel(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Risk計算を失敗しました")
		return
	}

	c.HTML(200, "result.html", m)
}
