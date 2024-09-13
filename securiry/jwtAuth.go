package securiry

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// check for valid admin token
func JWTadminAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		if err := ValidateAdminRoleJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// check for valid customer token
func JWTAuthCustomer() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		if err := ValidateCustomerRoleJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only pay Customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

// check for valid Anonymous token
func JWTAuthAnonymous() gin.HandlerFunc {
	return func(context *gin.Context) {

		if err := ValidateJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}

		if err := ValidateAnonymousRoleJWT(context); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}
