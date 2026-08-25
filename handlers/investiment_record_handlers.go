package handlers

import (
	"controle-investimento-back-end/domain/entities"
	"controle-investimento-back-end/domain/usecase"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// InvestmentRecordHandler handles HTTP requests and delegates to usecases.
type InvestmentRecordHandler struct {
	uc *usecase.InvestmentRecordUseCase
}

func NewInvestmentRecordHandler(uc *usecase.InvestmentRecordUseCase) *InvestmentRecordHandler {
	return &InvestmentRecordHandler{uc: uc}
}

func (h *InvestmentRecordHandler) SaveInvestmentRecord(c *gin.Context) {
	var input entities.InvestmentRecord
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := h.uc.SaveInvestmentRecord(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process SaveInvestmentRecord"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Investment record saved successfully"})
}

func (h *InvestmentRecordHandler) GetAllInvestmentRecords(c *gin.Context) {
	filter := c.Query("filter")
	ascendingStr := c.Query("ascending")

	ascending := true

	if ascendingStr != "" {
		var err error

		ascending, err = strconv.ParseBool(ascendingStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ascending inválido",
			})
			return
		}
	}

	records, err := h.uc.GetAllInvestmentRecords(filter, ascending)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process GetAllInvestmentRecords"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": records})
}

func (h *InvestmentRecordHandler) DataDashboard(c *gin.Context) {
	filter := c.Query("filter")

	data, err := h.uc.DataDashboard(filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process DataDashboard"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *InvestmentRecordHandler) AssetGrowth(c *gin.Context) {
	filter := c.Query("filter")

	result, err := h.uc.AssetGrowth(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process AssetGrowth"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *InvestmentRecordHandler) CategoryGrowth(c *gin.Context) {
	filter := c.Query("filter")

	result, err := h.uc.CategoryGrowth(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process CategoryGrowth"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *InvestmentRecordHandler) AvailableYears(c *gin.Context) {
	result, err := h.uc.AvailableYears()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process AvailableYears"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *InvestmentRecordHandler) GetLastInvestmentRecord(c *gin.Context) {
	result, err := h.uc.GetLastInvestmentRecord()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process GetLastInvestmentRecord"})
		return
	}
	c.JSON(http.StatusOK, result)
}
