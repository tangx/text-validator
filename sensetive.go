package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
)

var fileter *sensitive.Filter

func InitValidator(dict string) {
	fileter = sensitive.New()
	fileter.LoadWordDict(dict)
}

func IsValidateText(msg string) bool {
	ok, _ := fileter.Validate(msg)
	return ok
}

func ValidateText(c *gin.Context) {

	query := c.Query("query")
	if len(query) == 0 {
		c.JSON(http.StatusOK, Response{Code: 0, Message: "null"})
		return
	}

	ok, word := fileter.Validate(query)
	if ok {
		c.JSON(http.StatusOK, Response{Code: 0, Message: "ok"})
		return
	}

	c.JSON(http.StatusBadRequest, Response{Code: -1, Message: word})

}
