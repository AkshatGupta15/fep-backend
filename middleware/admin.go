package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/constants"
)

func EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := GetRoleID(ctx)

		if role != constants.GOD {
			// if role != constants.OPC && role != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func EnsurePsuedoAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := GetRoleID(ctx)

		// if role != constants.OPC && role != constants.GOD && role != constants.APC && role != constants.CHAIR {
		if role != constants.GOD {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
