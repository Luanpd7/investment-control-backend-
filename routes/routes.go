package routes

import (
	"controle-investimento-back-end/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *handlers.InvestmentRecordHandler) {

	router.POST("/saveInvestment", h.SaveInvestmentRecord)

	router.GET("/getAllInvestment", h.GetAllInvestmentRecords)

	router.GET("/dataDashboard", h.DataDashboard)

	router.GET("/assetGrowth", h.AssetGrowth)

	router.GET("/categoryGrowth", h.CategoryGrowth)

	router.GET("/availableYears", h.AvailableYears)

	router.GET("/lastInvestmentRecord", h.GetLastInvestmentRecord)
}
