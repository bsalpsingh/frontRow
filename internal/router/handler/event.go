package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EventRouter(v1 *gin.RouterGroup) {

	v1.GET("/events", getEvents)

}

func getEvents(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{

		"status":  "success",
		"content": []string{},
	})
}
