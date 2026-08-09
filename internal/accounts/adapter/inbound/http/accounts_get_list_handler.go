package http

import (
	responses "mini-project-a/internal/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AccountsHandler) GetAccountList(c *gin.Context) {
	accounts, err := h.accountsService.GetAccountList()

	if err != nil {
		responses.Error(
			c,
			http.StatusInternalServerError,
			"E9001",
			"Internal server error",
		)
		return
	}

	response := ToGetAccountListResponse(accounts)

	responses.Success(
		c,
		http.StatusOK,
		"S0000",
		"Accounts retrieved successfully",
		response,
	)

}
