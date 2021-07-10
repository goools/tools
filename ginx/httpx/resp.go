package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/goools/tools/errorx"
	"net/http"
)

func SetRespErr(c *gin.Context, err error) {
	var newErr errorx.IError
	var ok bool
	if newErr, ok = err.(errorx.IError); !ok {
		newErr = errorx.NewError(errorx.ServerError, err)
	}
	c.JSON(http.StatusOK, MakeResultWithError(newErr))
}

func SetRespJSON(c *gin.Context, data interface{}, message string) {
	if message == "" {
		message = "success"
	}
	c.JSON(http.StatusOK, JSONResult{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

func SetRespJSONPaged(c *gin.Context, data interface{}, message string, total int64) {
	if message == "" {
		message = "success"
	}
	c.JSON(http.StatusOK, JSONResultPaged{
		Code:    0,
		Message: message,
		Data:    data,
		Total:   total,
	})
}
