package resources

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
)

type Response struct {
	Message   string `json:"messsage"`
	Status    bool   `json:"status"`
	Data      any    `json:"data,omitempty"`
	Code      int    `json:"code,omitempty"`
	Duplicate bool   `json:"duplicate,omitempty"`
	Total     int    `json:"total,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

func Success(ctx *gin.Context, message string, data ...any) {
	response := Response{
		Message: message,
		Status:  true,
		Code:    http.StatusOK,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	ctx.JSON(http.StatusOK, response)
}
func Created(ctx *gin.Context, message string, data any) {
	ctx.JSON(http.StatusCreated, Response{
		Message: message,
		Status:  true,
		Code:    http.StatusCreated,
		Data:    data,
	})
}

func BadRequest(ctx *gin.Context, err any) {
	response := Response{
		Message: "validation failed",
		Status:  false,
		Code:    http.StatusBadRequest,
	}

	switch e := err.(type) {
	case string:
		response.Message = e
	case error:
		response.Message = e.Error()
	case map[string]string:
		response.Message = "validation failed"
		response.Data = e
	default:
		response.Message = fmt.Sprintf("%v", e)
	}

	ctx.JSON(http.StatusBadRequest, response)
}
func NotFound(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, Response{
		Message: err.Error(),
		Status:  false,
		Code:    http.StatusNotFound,
	})
}

func Conflict(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusConflict, Response{
		Message:   err.Error(),
		Status:    false,
		Code:      http.StatusConflict,
		Duplicate: true,
	})
}

func InternalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Message: err.Error(),
		Status:  false,
		Code:    http.StatusInternalServerError,
	})
}

func Paginated(ctx *gin.Context, message string, data any, total, offset, limit int) {
	ctx.JSON(http.StatusOK, Response{
		Message: message,
		Status:  true,
		Code:    http.StatusOK,
		Data:    data,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	})
}
