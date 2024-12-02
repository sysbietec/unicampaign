package router


import (
	"github.com/gin-gonic/gin"
	"github.com/sysbietec/unicampaign/internal/adapter/controllers"
)

func SetupRouter(opportunitiesController *controllers.OpportunitiesController) *gin.Engine{
	r := gin.Default()
	r.GET("/opportunities", opportunitiesController.GetAvailableOpportunities)
	return r
}