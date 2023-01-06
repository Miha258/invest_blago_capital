package handlers

import (
	"invest_blango_criptal_backend/service"
	"github.com/gin-gonic/gin"
)


type Handler struct {
	services *service.Service
}


func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{	
		api.GET("/account", h.getAccountData)
		
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.singUp)
			auth.POST("/sign-in", h.singIn)
		}

		api.Group("/user", h.userAuthorization)
		{
			account := api.Group("/account")
			{	
				account.PATCH("/set-password", h.singUp)
				account.PUT("/update-docs", h.updateAccountDocs)
				account.PATCH("/update-balance", h.updateAccountBalance)
				account.POST("/create-promocode", h.createPromocode)
				account.GET("/accept-promocode-usage", h.acceptPromocodeUsage)
			}
			investitions := api.Group("/investitions")
			{
				investitions.POST("/create", h.createInvestition)
			}
		}
	}
	return router
}