package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/loderunner/gin-500-middleware/message/service"
)

type SendHandlerRequest struct {
	Message string `form:"message" binding:"required"`
	To      string `form:"to" binding:"required"`
}

type SendHandlerResponse struct{}

func SendHandler(ctx *gin.Context) {
	var req SendHandlerRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		// 400 error here: the user's request was not well-formed
		// No log as the system is operating as intended
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request."})
		return
	}

	messageServiceValue, hasMessageService := ctx.Get(messageServiceContextKey)
	if !hasMessageService {
		// Missing value on context: log error and abort handler
		// Will result in a 500 through the middleware
		fmt.Println("error: missing message service")
		ctx.Abort()
		return
	}

	messageService, ok := messageServiceValue.(*service.MessageService)
	if !ok {
		// Invalid value on context: log error and stop handler
		// Will result in a 500 through the middleware
		fmt.Println("error: invalid message service in context")
		ctx.Abort()
		return
	}

	err = messageService.Send(req.Message, req.To)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			// User is not found in database: respond with a 404 error
			// Do not log as the system is operating as intended
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				ErrorResponse{Message: fmt.Sprintf("User \"%s\" not found.", req.To)},
			)
			return
		}

		// Other unidentified error: log error and stop handler
		// Will result in a 500 error through the middleware
		fmt.Println("error:", err)
		ctx.Abort()
		return
	}

	ctx.Status(http.StatusAccepted)
}
