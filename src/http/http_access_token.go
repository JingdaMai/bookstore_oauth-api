package http

import (
	"github.com/JingdaMai/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h * accessTokenHandler) GetById (c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNotImplemented, accessToken)
}
