package handlers

import (
	"os"

	"github.com/aaalik/api-alik/internal/api/constants"
	"github.com/aaalik/api-alik/pkg/ahttp"
	"github.com/aaalik/api-alik/pkg/alog"
	"github.com/kataras/iris/v12"
)

func HttpError(c iris.Context, httpError ahttp.ErrorResponse, err error) {
	c.StatusCode(httpError.Status)

	alog.Logger.Error(err)

	if os.Getenv(constants.AppEnv) == constants.EnvDevelopment {
		traces := alog.GetTracer(err)

		response := ahttp.ErrorDebugResponse{
			Status:  httpError.Status,
			Code:    httpError.Code,
			Message: err.Error(),
			Debug:   traces,
		}

		c.JSON(response)
	} else {
		response := ahttp.ErrorResponse{
			Status:  httpError.Status,
			Code:    httpError.Code,
			Message: err.Error(),
		}

		c.JSON(response)
	}

	c.StopExecution()
	return
}

func HttpSuccess(c iris.Context, data interface{}, message string) {
	response := ahttp.Response{}
	response.Status = ahttp.ResponseSuccess
	response.Message = message

	if data != nil {
		response.Data = data
	}

	c.JSON(response)
	c.StopExecution()
	return
}
