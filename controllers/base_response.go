package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	res := BaseResponse{}
	res.Meta.Status = http.StatusOK
	res.Meta.Message = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
	c.Abort()
}

func ErrorResponse(c *gin.Context, status int, err string, errs error) {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Messages = []string{errs.Error()}
	res.Data = err
	c.JSON(http.StatusOK, res)
	c.Abort()
}
