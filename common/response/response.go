package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pan/model"
)

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, model.Resp{
		Msg:  "success",
		Data: data,
	})
}

func OnlyMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.Resp{
		Msg: msg,
	})
}

func Data(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, model.Resp{
		Msg:  msg,
		Data: data,
	})
}

func InternalServerError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, model.Resp{
		Msg: msg,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, model.Resp{
		Msg: msg,
	})
}
