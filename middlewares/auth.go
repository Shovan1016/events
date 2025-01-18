package middlewares

import (
	"event-management/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	var userId int64
	var err error

	token := context.Request.Header.Get("Authorization")

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
		userId, err = utils.VarifyToken(token)
		fmt.Print(userId, "userId")
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
