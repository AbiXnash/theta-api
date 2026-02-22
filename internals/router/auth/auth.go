package auth

import (
	"net/http"

	"github.com/AbiXnash/theta-api/internals/components"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, _ *components.Components) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		errors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, validationErrorMessage("Validation failed", errors))
		return
	}

	c.JSON(http.StatusOK, successMessage("User logged in successfully"))
}

func UserRegister(c *gin.Context, _ *components.Components) {
	var req RegisterRequest

	// General validation
	if err := c.ShouldBind(&req); err != nil {
		errors := formatValidationErrors(err)
		c.JSON(http.StatusBadRequest, validationErrorMessage("Validation failed", errors))
		return
	}

	// Sastra email for Sastra students validation
	businessErrors := validateBusinessRules(req)
	if len(businessErrors) > 0 {
		c.JSON(http.StatusBadRequest, businessErrorMessage("Validation failed", businessErrors))
		return
	}

	c.JSON(http.StatusOK, successMessage("User registered successfully"))
}
