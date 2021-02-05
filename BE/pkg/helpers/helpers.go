package helpers

import (
	"github.com/gin-gonic/gin"
)

// ErrorExit sets the context and returns true when an error is found
func ErrorExit(err error, ctx *gin.Context, msg string, status int) bool {
	if err != nil {
		ctx.JSON(status, gin.H{"detail": msg})
		return true
	}
	return false
}
