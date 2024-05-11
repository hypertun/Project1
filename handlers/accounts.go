package handlers

import (
	"Project1/models"
	"Project1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Accounts struct {
	accountsService services.AccountsInterface
}

func NewAccountsHandler(accountsService services.AccountsInterface) Accounts {
	return Accounts{
		accountsService: accountsService,
	}
}

func (a Accounts) Post(c *gin.Context) {
	var postRequest models.AccountCreationRequest
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	account, err := postRequest.ConvertToAccount()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if err := a.accountsService.PostAccount(c.Request.Context(), account); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (a Accounts) Get(c *gin.Context) {
	accountIDString := c.Param("account_id")

	accountID, err := strconv.Atoi(accountIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	if accountID <= 0 {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Error:      models.ErrInvalidAccountID.Error(),
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	account, err := a.accountsService.GetAccount(c.Request.Context(), accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Error:      err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, account)
}
