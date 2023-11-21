package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/constant"
	"github.com/satriaprayoga/kofin/internal/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return buildResponse(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func buildResponse[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Key:     status,
		Message: message,
		Data:    data,
	}
}

func BuildResponseList(responseStatus constant.ResponseStatus, data interface{}, page int, total int64, lastPage int) dto.ApiResponseList {
	return buildResponseList(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data, page, total, lastPage)
}

func buildResponseList(status string, message string, data interface{}, page int, total int64, lastPage int) dto.ApiResponseList {
	return dto.ApiResponseList{
		Key:      status,
		Message:  message,
		Data:     data,
		Page:     page,
		Total:    total,
		LastPage: lastPage,
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	panicException(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func panicException(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")
		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusNotFound, buildResponse(key, msg, Null()))
			c.Abort()
		case constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, buildResponse(key, msg, Null()))
			c.Abort()
		case constant.UnknownError.GetResponseStatus():
			c.JSON(http.StatusBadGateway, buildResponse(key, msg, Null()))
			c.Abort()
		case constant.InvalidRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest, buildResponse(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, buildResponse(key, msg, Null()))
			c.Abort()
		}
	}
}
