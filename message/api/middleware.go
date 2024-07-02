package api

import (
	"github.com/gin-gonic/gin"

	"github.com/loderunner/gin-500-middleware/message/service"
)

const messageServiceContextKey = "message-service"

func MessageServiceMiddleware(ctx *gin.Context) {
	ctx.Set(messageServiceContextKey, service.NewMessageService())
	ctx.Next()
}
