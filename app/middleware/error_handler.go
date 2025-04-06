package middleware

import (
	"errors"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/zayn1510/goarchi/app/resources"
)

var (
	ErrNotFound      = errors.New("data tidak ditemukan")
	ErrDuplicate     = errors.New("data duplikat")
	ErrUnauthorized  = errors.New("tidak memiliki izin")
	ErrInvalidInput  = errors.New("input tidak valid")
	ErrInternalError = errors.New("terjadi kesalahan pada server")
)

func GetHttpStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrDuplicate):
		return http.StatusConflict
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// ambil error terakhir
		err := c.Errors.Last()
		if err != nil {
			if strings.Contains(err.Error(), "tidak ditemukan") {
				c.JSON(http.StatusNotFound, resources.Response{
					Message:   err.Error(),
					Status:    false,
					Code:      http.StatusNotFound,
					Duplicate: true,
				})
				return
			}
		}
	}
}
