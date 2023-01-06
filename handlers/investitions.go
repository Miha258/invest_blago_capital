package handlers


import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"invest_blango_criptal_backend/models"
)


func (h *Handler) createInvestition(c *gin.Context) {
	var input models.Investition

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid json parameters")
		return
	}
	
	id, err := h.services.CreateInvestition(input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"investition_id": id,
	})
}


