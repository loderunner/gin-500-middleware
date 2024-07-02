package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/loderunner/gin-500-middleware/message/api"
)

func main() {
	r := gin.Default()

	// OOPS! Forgot to add the MessageServiceMiddleware, so no MessageService
	// will be on the context. This is a bug, and should generate a 500 error.
	// r.Use(MessageServiceMiddleware)
	r.Use(InternalServerErrorMiddleware)

	r.GET("/send", api.SendHandler)

	r.Run()
}

func InternalServerErrorMiddleware(ctx *gin.Context) {
	ctx.Next()
	if ctx.IsAborted() && ctx.Writer.Status() == 200 {
		// Aborted without a status code -> 500 error
		ctx.Status(http.StatusInternalServerError)
	}
}
