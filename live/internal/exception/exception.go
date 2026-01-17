package exception

import (
	"live/internal/response"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS         = 0
	PARAMETER_ERROR = 4001

	NOT_FOUND = 4041

	NO_ALLOW_METHDO = 4051

	INTERNA_ERROR     = 5000
	UNKNOWN_ERROR     = 5001
	INSERT_DATA_ERROR = 5002
	UPDATE_DATA_ERROR = 5003
	DELETE_DATA_ERROR = 5004
)

type HandlerFunc func(c *gin.Context) error

func ErrHandler(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		err = h(c)
		if err != nil {
			var apiException *response.HttpResp
			if h, ok := err.(*response.HttpResp); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				apiException = response.UnknownError(e.Error())
			} else {
				apiException = response.InternalError()
			}
			c.JSON(apiException.StatusCode, apiException)
			return
		}
	}
}

var (
	ErrInternal   = response.InternalError()
	ErrNoRows     = response.NotFound()
	ErrParameter  = response.ParameterError(nil)
	ErrInsertFail = response.InsertFail()
	ErrUpdateFail = response.UpdateFail()
	ErrDeleteFail = response.DeleteFail()
)
