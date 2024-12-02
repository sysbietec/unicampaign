package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sysbietec/unicampaign/internal/usecases"
)

type OpportunitiesController struct {
	UseCase *usecases.OpportunitiesUseCase
}

func NewOpportunitiesController(uc *usecases.OpportunitiesUseCase) *OpportunitiesController {
	return &OpportunitiesController{UseCase: uc}
}

func (oc *OpportunitiesController) GetAvailableOpportunities(c *gin.Context) {
	opportunities, err := oc.UseCase.FetchAvailableOpportunities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": opportunities})
}
