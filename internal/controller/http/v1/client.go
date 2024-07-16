package v1

import (
	"net/http"

	"github.com/andrew-nino/ATM/entity"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// Adding a new client. If successful, we get the client id
func (h *Handler) addAccount(c *gin.Context) {
	newClient := entity.Client{}

	if err := c.BindJSON(&newClient); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	clientID, err := h.services.AddAccount(newClient)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "client add failed")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: clientID})
}
