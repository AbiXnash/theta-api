package auth

import "github.com/gin-gonic/gin"

func validationErrorMessage(message string, errors []ValidationError) gin.H {
	return gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	}
}

func businessErrorMessage(message string, errors map[string]string) gin.H {
	return gin.H{
		"status":  "error",
		"message": message,
		"errors":  errors,
	}
}

func successMessage(message string) gin.H {
	return gin.H{
		"status":  "success",
		"message": message,
	}
}
