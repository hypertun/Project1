package handlers

import (
	"Project1/models"
	"Project1/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transactions struct {
	transactionsService services.TransactionsInterface
}

func NewTransactionsHandler(accountsService services.TransactionsInterface) Transactions {
	return Transactions{
		transactionsService: accountsService,
	}
}

func (t Transactions) Post(c *gin.Context) {
	var postRequest models.TransactionHandlerReq
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	transaction, err := postRequest.ConvertToTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if err := t.transactionsService.PostTransaction(c.Request.Context(), transaction); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
