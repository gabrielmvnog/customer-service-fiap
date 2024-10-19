package controllers

import (
	"net/http"

	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"

	"github.com/gin-gonic/gin"
)

func Healthcheck(context *gin.Context) {
	var pingReponse dtos.PingReponse = dtos.PingReponse{Status: "OK"}

	context.IndentedJSON(http.StatusOK, pingReponse)
}
