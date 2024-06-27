package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, request int, message string, error string) {
	c.JSON(request, gin.H{
		"success": false,
		"message": message,
		"error":   error,
	})
}

func SuccessResponse(c *gin.Context, request int, message string, data interface{}) {
	c.JSON(request, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func AbortResponse(c *gin.Context, request int, message string) {
	c.AbortWithStatusJSON(request, gin.H{
		"success": false,
		"message": message,
		"error":   "Unauthorized",
	})
}

// ValidateID ensures the ID parameter is present and valid, returning it as a uint
func ValidateID(c *gin.Context, param string) (uint, bool) {
	id := c.Param(param)
	if id == "" {
		ErrorResponse(c, http.StatusBadRequest, "ID parameter is missing", "")
		return 0, false
	}

	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid ID parameter", err.Error())
		return 0, false
	}

	return uint(parsedID), true
}
