package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResp struct {
	StatusCode int         `json:"-"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func (e *HttpResp) Error() string {
	return e.Msg
}

func Success(data interface{}) *HttpResp {
	return &HttpResp{
		Code:       SUCCESS,
		StatusCode: http.StatusOK,
		Data:       data,
		Msg:        "success",
	}
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	c.JSON(handleErr.StatusCode, handleErr)
	return
}

func HandleNoAllowMethod(c *gin.Context) {
	handleErr := NoAllowMethod()
	c.JSON(handleErr.StatusCode, handleErr)
	return
}

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

// NotFound Not found page error response
func NotFound() *HttpResp {
	return newHttpException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound), nil)
}

// ParameterError 參數錯誤
func ParameterError(errMsg []string) *HttpResp {
	return newHttpException(http.StatusBadRequest, PARAMETER_ERROR, "Missing required parameter error or parameter setting error", nil)
}

// InternalError Service Internal Error response
func InternalError() *HttpResp {
	return newHttpException(http.StatusInternalServerError, INTERNA_ERROR, http.StatusText(http.StatusInternalServerError), nil)
}

// UnknownError Unknown Error response
func UnknownError(message string) *HttpResp {
	return newHttpException(http.StatusInternalServerError, UNKNOWN_ERROR, message, nil)
}

func InsertFail() *HttpResp {
	return newHttpException(http.StatusInternalServerError, INSERT_DATA_ERROR, "Data creation failed", nil)
}

func UpdateFail() *HttpResp {
	return newHttpException(http.StatusInternalServerError, UPDATE_DATA_ERROR, "Data update failed", nil)
}

func DeleteFail() *HttpResp {
	return newHttpException(http.StatusInternalServerError, DELETE_DATA_ERROR, "Data delete failed", nil)
}

// NoAllowMethod
// Not Allow Method
func NoAllowMethod() *HttpResp {
	return newHttpException(http.StatusMethodNotAllowed, NO_ALLOW_METHDO, http.StatusText(http.StatusMethodNotAllowed), nil)
}

func newHttpException(statusCode int, errorCode int, msg string, data []string) *HttpResp {
	return &HttpResp{
		Code:       errorCode,
		StatusCode: statusCode,
		Msg:        msg,
		Data:       data,
	}
}
