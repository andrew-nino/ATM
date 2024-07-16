package v1

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/ATM/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deposit(c *gin.Context) {

	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}

	clientID, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	transaction := entity.Transaction{}

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body transaction")
		return
	}

	transaction.AccountId = clientID

	err = h.services.Deposit(transaction)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Operation deposit is fail")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) withdraw(c *gin.Context) {
	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}

	clientID, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	transaction := entity.Transaction{}

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body transaction")
		return
	}

	transaction.AccountId = clientID
	err = h.services.Withdraw(transaction)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Operation deposit is fail")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) getBalance(c *gin.Context) {
	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}

	clientID, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	balance := h.services.GetBalance(clientID)

	c.JSON(http.StatusOK, gin.H{
        "balance": balance,
    })

}
